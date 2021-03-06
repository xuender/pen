###
main.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
PenCtrl = ($scope, $log, $modal, ngSocket, lss, $q, $location)->
  ### 主控制器 ###
  $scope.menus = MENUS
  $scope.token = ''
  $scope.showLeft = true
  commands = {}
  $scope.alerts = []
  $scope.alert = (msg, type='success')->
    # 提示
    $scope.alerts.push(
      type: type
      msg: msg
    )
  $scope.error = (msg)->
    # 错误
    $scope.alert(msg, 'danger')
  $scope.closeAlert = (index)->
    # 关闭提示框
    $scope.alerts.splice(index, 1)
  $scope.eventLogin = (data)->
    # 登陆事件
    $scope.isLogin = false
    if data == 'OK'
      $scope.isLogin = true
      # 获取字典版本
      $scope.send('base', BASE.ZDBB, JSON.stringify($scope.dictVer))
      for r in readies
        console.info 'run ready'
        r()
    if data == 'ERROR_PASSWORD'
      $scope.showLogin()
    if data == 'ERROR_NICK'
      $scope.showLogin()
    if data == 'login'
      if $scope.user.token
        $scope.wsLogin()
      else
        $scope.showLogin(true)
  # 用户数量
  $scope.userCount = 0
  $scope.eventCount = (data)->
    # 在线用户数修改事件
    $scope.userCount = data
    $log.info 'userCount: ', data
  $scope.registerEvent = (code, event, cb)->
    # 注册事件
    commands["#{code}.#{event}"] = cb
  $scope.registerEvent('base', BASE.MSG, (data)->
    d = JSON.parse(data)
    $log.debug d
    if Object.prototype.toString.call(d) == "[object String]"
      alert(d)
    else
      alert(d.msg)
  )
  $scope.registerEvent('base', BASE.DL, $scope.eventLogin)
  $scope.registerEvent('base', BASE.RS, $scope.eventCount)
  $scope.registerEvent('base', BASE.ZDHQ, (data)->
    $log.debug("dict....."+data)
    d = JSON.parse(data)
    $log.debug(d)
    $scope.dictVer[d.type] = d.ver
    lss.bind($scope, 'dict_' + d.type, {})
    $scope['dict_' + d.type] = d.dict
  )
  $scope.getDict = (type)->
    # 获取字典
    def = $q.defer()
    ret = []
    for k, v of $scope['dict_' + type]
      ret.push(
        id: k
        title: v
      )
    def.resolve(ret)
    def
  # 消息处理
  $scope.like = false
  if typeof(WEB) == "undefined"
    ws = ngSocket("ws://#{location.origin.split('//')[1]}/ws")
  else
    ws = ngSocket("ws://#{WEB}/ws")
  ws.onMessage((data)->
    $scope.like = !$scope.like
    dmsg = JSON.parse(data.data)
    $scope.tract = dmsg.tract
    $log.debug("ws onMessage:#{dmsg.code}.#{dmsg.event} data:#{dmsg.data}")
    k = "#{dmsg.code}.#{dmsg.event}"
    $log.info "k:#{k}"
    if k of commands
      commands[k](dmsg.data)
  )
  # 登录状态
  $scope.isLogin = false
  $scope.wsLogin = ->
    #登录
    $log.debug('login', $scope.user)
    $log.debug($scope.tract)
    $scope.token = md5($scope.tract + $scope.user.token)
    $scope.send('base', BASE.DL, JSON.stringify(
      'nick': $scope.user.nick
      'token': $scope.token
    ))
  $scope.init = ->
    ### 初始化 ###
    lss.bind($scope, 'dictVer', {})
    for k of $scope.dictVer
      lss.bind($scope, 'dict_' + k, {})

    lss.bind($scope, 'user',
      nick: '来宾'
      token: ''
    )
    ws.send(
      code: 'base'
      event: 100
    )
  $scope.showLogin = (init = false)->
    ### 显示登录窗口 ###
    i = $modal.open(
      templateUrl: 'base/login.html'
      controller: LoginCtrl
      backdrop: 'static'
      keyboard: false
      size: 'sm'
      resolve:
        user: ->
          {
            nick: $scope.user.nick
            password: ''
          }
        init: ->
          init
    )
    i.result.then((user)->
      $scope.user.nick= user.nick
      $scope.user.token = md5((new Date()).format('yyyy-MM-dd') + md5(md5(user.nick + user.password)))
      $scope.wsLogin()
    ,->
      $log.info '取消'
    )
  $scope.edit = ->
    ### 编辑用户 ###
    $scope.showLogin()

  $scope.logout = ->
    ### 登出 ###
    $scope.user.token = ''
    $scope.send('base', BASE.DC)
    $scope.showLogin(true)
  $scope.send = (code, event, data=null)->
    # 发送数据
    $log.debug('code:%s, event:%d', code, event)
    if data
      if (typeof data != 'string') || data.constructor != String
        data = JSON.stringify(data)
      $log.debug 'send data:%s', data
      ws.send(
        code: code
        event: event
        data: data
        token: $scope.token
      )
    else
      ws.send(
        code: code
        event: event
        token: $scope.token
      )
  readies = []
  $scope.ready = (f)->
    # 准备好执行的命令
    console.info 'ready'
    if $scope.isLogin
      f()
    else
      readies.push(f)
  lss.bind($scope, 'history', [])
  $scope.addHistory = (route, subName=null)->
    # 增加历史操作
    name = route.$$route.name
    if subName
      name = "#{name} [ #{subName} ]"
    $scope.history.unshift(
      name: name
      path: $location.$$path
      time: (new Date()).format('yyyy-MM-dd hh:mm:ss')
    )
    if $scope.history.length > 10
      $scope.history.pop()
  $scope.init()

PenCtrl.$inject = [
  '$scope'
  '$log'
  '$modal'
  'ngSocket'
  'localStorageService'
  '$q'
  '$location'
]
