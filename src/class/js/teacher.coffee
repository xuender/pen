###
volunteer.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
TeacherCtrl= ($scope, $log, $modalInstance, d, pen)->
  $scope.d = d
  $scope.pen = pen
  $scope.ok = ->
    $log.debug 'ok'
    $scope.isSend = true
    $scope.pen.send('class', CLASS.教师编辑, $scope.d)
  $scope.del = ->
    $scope.isSend = true
    $scope.pen.send('class', CLASS.教师删除, $scope.d.Id)
  $scope.cancel = ->
    $modalInstance.dismiss('cancel')
  $scope.close = (data)->
    if 'ok' == data
      $modalInstance.close($scope.d)
  $scope.pen.registerEvent('class', CLASS.教师编辑, $scope.close)
  $scope.pen.registerEvent('class', CLASS.教师删除, $scope.close)
  $scope.openDate = ($event)->
    $event.preventDefault()
    $event.stopPropagation()
    $scope.opened = true
TeacherCtrl.$inject = [
  '$scope'
  '$log'
  '$modalInstance'
  'd'
  'pen'
]
