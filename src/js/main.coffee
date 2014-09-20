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
PenCtrl.$inject = [
  '$scope'
  '$modal'
  'ngSocket'
  'localStorageService'
]
