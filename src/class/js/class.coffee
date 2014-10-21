###
volunteer.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
ClassCtrl= ($scope, $log, $modalInstance, d, pen)->
  $scope.d = d
  $scope.pen = pen
  $scope.ok = ->
    $log.debug 'ok'
    $scope.isSend = true
    $scope.pen.send('class', CLASS.BJBJ, $scope.d)
  $scope.cancel = ->
    $modalInstance.dismiss('cancel')
  $scope.del = ->
    $scope.isSend = true
    $scope.pen.send('class', CLASS.BJSC, $scope.d.Id)
  $scope.close = (data)->
    if 'ok' == data
      $modalInstance.close($scope.d)
  $scope.pen.registerEvent('class', CLASS.BJBJ, $scope.close)
  $scope.pen.registerEvent('class', CLASS.BJSC, $scope.close)
  $scope.openDate = ($event, end=false)->
    $event.preventDefault()
    $event.stopPropagation()
    if end
      $scope.endOpened = true
    else
      $scope.opened = true
ClassCtrl.$inject = [
  '$scope'
  '$log'
  '$modalInstance'
  'd'
  'pen'
]
