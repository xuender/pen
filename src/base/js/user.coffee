###
user.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
UserCtrl = ($scope, $log, $modalInstance, d, pen)->
  $scope.d = d
  $scope.gender = gender
  $scope.pen = pen
  $scope.ok = ->
    $log.debug 'ok'
    $scope.pen.send('base', BASE.修改用户, $scope.d)
    #$modalInstance.close($scope.user)
  $scope.cancel = ->
    $modalInstance.dismiss('cancel')
  $scope.pen.registerEvent('base', BASE.修改用户, (data)->
    if 'ok' == data
      $modalInstance.close($scope.d)
      #$scope.current.$edit = false
      #$scope.send('base', BASE.getDict, $scope.type)
  )

UserCtrl.$inject = [
  '$scope'
  '$log'
  '$modalInstance'
  'user'
  'pen'
]

