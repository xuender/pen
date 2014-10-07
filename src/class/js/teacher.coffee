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
    $scope.pen.send('class', CLASS.编辑教师, $scope.d)
  $scope.cancel = ->
    $modalInstance.dismiss('cancel')
  $scope.pen.registerEvent('class', CLASS.编辑教师, (data)->
    if 'ok' == data
      $modalInstance.close($scope.d)
  )
  $scope.open = ($event)->
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
