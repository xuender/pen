###
user.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
UserCtrl = ($scope, $log, $modalInstance, user, gender, pen)->
  $scope.user = user
  $scope.gender = gender
  $scope.pen = pen
  $scope.ok = ->
    $log.debug 'ok'
    $scope.pen.send('base', CONST.updateUser, $scope.user)
    #$modalInstance.close($scope.user)
  $scope.cancel = ->
    $modalInstance.dismiss('cancel')
  $scope.pen.registerEvent('base', CONST.updateUser, (data)->
    if 'ok' == data
      $modalInstance.close($scope.user)
      #$scope.current.$edit = false
      #$scope.send('base', CONST.getDict, $scope.type)
  )

UserCtrl.$inject = [
  '$scope'
  '$log'
  '$modalInstance'
  'user'
  'gender'
  'pen'
]

