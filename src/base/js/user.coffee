###
user.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
UserCtrl = ($scope, $log, $modalInstance, user, gender)->
  $scope.user = user
  $scope.gender = gender
  $scope.ok = ->
    $modalInstance.close($scope.user)
  $scope.cancel = ->
    $modalInstance.dismiss('cancel')

UserCtrl.$inject = [
  '$scope'
  '$log'
  '$modalInstance'
  'user'
  'gender'
]

