###
user.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
UserCtrl = ($scope, $log)->
  $log.info '用户管理'
  $scope.users = []
  $scope.userAll = (data)->
    # 获取所有用户
    $log.debug('收到用户信息: %s', data)
    $scope.users = JSON.parse(data)
  $scope.registerEvent('base', CONST.userAll, $scope.userAll)
  $scope.send('base', CONST.userAll)

UserCtrl.$inject = [
  '$scope'
  '$log'
]

