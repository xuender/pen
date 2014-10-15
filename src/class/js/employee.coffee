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
    $scope.pen.send('class', CLASS.编辑雇员, $scope.d)
  $scope.del = ->
    $scope.pen.send('class', CLASS.删除雇员, $scope.d.Id)
  $scope.cancel = ->
    $modalInstance.dismiss('cancel')
  $scope.pen.registerEvent('class', CLASS.编辑雇员, (data)->
    if 'ok' == data
      $modalInstance.close($scope.d)
  )
  $scope.pen.registerEvent('class', CLASS.删除雇员, (data)->
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
