###
main.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
PenCtrl = ($scope, $log, $modal, ngSocket, lss, $q)->
  ### 主控制器 ###
  $scope.token = ''
  $scope.showLeft = true
  commands = {}
  $scope.eventLogin = (data)->
    # 登陆事件
    $scope.isLogin = false
    if data == 'OK'
      $scope.isLogin = true
      # 获取版本
      $scope.send('base', CONST.dictVer, JSON.stringify($scope.dictVer))
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
  $scope.registerEvent('base', CONST.login, $scope.eventLogin)
  $scope.registerEvent('base', CONST.count, $scope.eventCount)
  $scope.registerEvent('base', CONST.dictVer, (data)->
    $scope.dictVer = JSON.parse(data)
    for k of $scope.dictVer
      lss.bind($scope, 'dict_' + k, {})
  )
  $scope.registerEvent('base', CONST.dict, (data)->
    $log.debug("dict....."+data)
    d = JSON.parse(data)
    $scope['dict_' + d.type] = d.data
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
  ws = ngSocket("ws://#{location.origin.split('//')[1]}/ws")
  ws.onMessage((data)->
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
    $log.info('login', $scope.user)
    v = $scope.user
    $log.info($scope.tract)
    $scope.token = md5($scope.tract + $scope.user.token)
    v.token = $scope.token
    $scope.send('base', CONST.login, JSON.stringify(v))
  $scope.init = ->
    ### 初始化 ###
    lss.bind($scope, 'dictVer', {})
    for k of $scope.dictVer
      lss.bind($scope, 'dict_' + k, {})
    user = lss.get('user')
    if user == null
      $scope.user =
        nick: '来宾'
        token: ''
    else
      $scope.user = user
    ws.send(
      code: 'base'
      event: 100
    )
  $scope.showLogin = (init = false)->
    ### 显示登录窗口 ###
    i = $modal.open(
      templateUrl: 'login.html'
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
      lss.set('user', $scope.user)
      $scope.wsLogin()
    ,->
      $log.info '取消'
    )
  $scope.edit = ->
    ### 编辑用户 ###
    $scope.showLogin()

  $scope.logout = ->
    ### 登出 ###
    $scope.user =
      nick: '来宾'
      token: ''
    lss.remove('user')
    $scope.send('base', CONST.logout)
    $scope.showLogin(true)
  $scope.send = (code, event, data=null)->
    # 发送数据
    $log.debug('code:%s, event:%d', code, event)
    if data
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
  $scope.addHistory = (route)->
    # 增加历史操作
    $log.info route.$$route.name
    $scope.history.unshift(
      name: route.$$route.name
      path: route.$$route.originalPath
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
]
