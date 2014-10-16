###
volunteer.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
EmployeeCtrl= ($scope, $log, $modalInstance, d, pen)->
  $scope.d = d
  $scope.pen = pen
  $scope.ok = ->
    $log.debug 'ok'
    $log.debug $scope.d
    $scope.isSend = true
    $scope.pen.send('class', CLASS.雇员编辑, $scope.d)
  $scope.del = ->
    $scope.isSend = true
    $scope.pen.send('class', CLASS.雇员删除, $scope.d.Id)
  $scope.cancel = ->
    $modalInstance.dismiss('cancel')
  $scope.pen.registerEvent('class', CLASS.雇员编辑, (data)->
    if 'ok' == data
      $modalInstance.close($scope.d)
  )
  $scope.pen.registerEvent('class', CLASS.雇员删除, (data)->
    if 'ok' == data
      $modalInstance.close($scope.d)
  )
EmployeeCtrl.$inject = [
  '$scope'
  '$log'
  '$modalInstance'
  'd'
  'pen'
]
