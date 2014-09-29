###
main.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
CONST =
  login: 0
  logout: 1
  count: 2
  userAll: 3
PenCtrl = ($scope, $log, $modal, ngSocket, lss)->
  ### 主控制器 ###
  $scope.token = ''
  $scope.showLeft = true
  commands = {}
  $scope.eventLogin = (data)->
    # 登陆事件
    $scope.isLogin = false
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
    $scope.isLogin = true
  $scope.init = ->
    ### 初始化 ###
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
      $scope.user.token = md5((new Date()).format('yyyy-MM-dd') + md5(md5(user.password)))
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
  $scope.init()

PenCtrl.$inject = [
  '$scope'
  '$log'
  '$modal'
  'ngSocket'
  'localStorageService'
]
