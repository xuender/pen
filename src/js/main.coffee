###
main.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
angular.module('pen', [
  'ui.bootstrap'
  'ngSocket'
  'LocalStorageModule'
  #'hotkey'
  #'angularFileUpload'
  #'textAngular'
])

PenCtrl = ($scope, $modal, ngSocket, lss)->
  ### 主控制器 ###
  $scope.token = ''
  commands = {}
  $scope.eventLogin = (data)->
    # 登陆事件
    if data == 'ERROR_PASSWORD'
      $scope.showLogin()
    if data == 'ERROR_NICK'
      $scope.showLogin()

  commands['base.login'] = $scope.eventLogin
  ws = ngSocket("ws://#{location.origin.split('//')[1]}/ws")
  ws.onMessage((data)->
    dmsg = JSON.parse(JSON.parse(data.data))
    $scope.tract = dmsg.tract
    commands[dmsg.event](dmsg.data)
  )

  $scope.send = ->
    ### 发送消息 ###
    console.info('test')
    ws.send(
      Event: 'test'
      Data: 'xxdfdfdfa'
    )
  $scope.isLogin = false
  $scope.wsLogin = ->
    #登录
    console.info('login', $scope.user)
    v = $scope.user
    $scope.token = md5($scope.tract + $scope.user.token)
    v.token = $scope.token
    ws.send(
      event: 'base.login'
      data: JSON.stringify(v)
      token: $scope.token
    )
  $scope.init = ->
    ### 初始化 ###
    user = lss.get('user')
    if user == null
      $scope.user =
        nick: '来宾'
        token: ''
      $scope.showLogin(true)
    else
      $scope.user = user
      $scope.wsLogin()
      #ws.send(
      #  Command: 'init'
      #)
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
          angular.copy($scope.user)
        init: ->
          init
    )
    i.result.then((user)->
      $scope.user = angular.copy(user)
      $scope.user.token = md5((new Date()).format('yyyy-MM-dd') + md5(md5($scope.user.token)))
      lss.set('user', $scope.user)
      $scope.wsLogin()
      #if init
      #  ws.send(
      #    Command: 'init'
      #  )
    ,->
      console.info '取消'
    )
  $scope.edit = ->
    ### 编辑用户 ###
    $scope.showLogin()

  $scope.logout = ->
    ### 登出 ###
    $scope.user =
      email: ''
      nick: ''
    lss.remove('user')
    ws.send(
      Command: 'logout'
    )
    $scope.messages = []
    $scope.showLogin(true)
  $scope.init()
PenCtrl.$inject = [
  '$scope'
  '$modal'
  'ngSocket'
  'localStorageService'
]
Date.prototype.format = (format)->
  o =
    "M+": this.getMonth()+1 #month 
    "d+": this.getDate() #day 
    "h+": this.getHours() #hour 
    "m+": this.getMinutes() #minute 
    "s+": this.getSeconds() #second 
    "q+": Math.floor((this.getMonth()+3)/3) #quarter 
    "S": this.getMilliseconds() #millisecond 

  if(/(y+)/.test(format))
    format = format.replace(RegExp.$1, (this.getFullYear()+"").substr(4 - RegExp.$1.length))
  for k of o
    if new RegExp("(#{k})").test(format)
      v = o[k]
      if RegExp.$1.length!=1
        v = ("00"+ o[k]).substr((""+ o[k]).length)
      format = format.replace(RegExp.$1, v)
  format
