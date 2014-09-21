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
  ws = ngSocket("ws://#{location.origin.split('//')[1]}/ws")
  ws.onMessage((data)->
    #dmsg = JSON.parse(data.data)
    console.info(data)
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
    ws.send(
      Event: 'base.login'
      Data: JSON.stringify($scope.user)
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
