###
volunteer.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
StudentCtrl= ($scope, $log, $modalInstance, d, pen)->
  $scope.d = d
  $scope.pen = pen
  $scope.ok = ->
    $log.debug 'ok'
    $scope.isSend = true
    $scope.pen.send('class', CLASS.XYBJ, $scope.d)
  $scope.cancel = ->
    $modalInstance.dismiss('cancel')
  $scope.del = ->
    $scope.isSend = true
    $scope.pen.send('class', CLASS.XYSC, $scope.d.Id)
  $scope.close = (data)->
    if 'ok' == data
      $modalInstance.close($scope.d)
  $scope.pen.registerEvent('class', CLASS.XYBJ, $scope.close)
  $scope.pen.registerEvent('class', CLASS.XYSC, $scope.close)
  $scope.openDate = ($event)->
    $event.preventDefault()
    $event.stopPropagation()
    $scope.opened = true
StudentCtrl.$inject = [
  '$scope'
  '$log'
  '$modalInstance'
  'd'
  'pen'
]
