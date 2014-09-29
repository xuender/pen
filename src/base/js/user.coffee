###
user.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
UserCtrl = ($scope, $log, ngTableParams, $filter)->
  $log.info '用户管理'
  $scope.users = []
  $scope.userAll = (data)->
    # 获取所有用户
    #$log.debug('收到用户信息: %s', data)
    $scope.users = JSON.parse(data)
    $scope.tableParams = new ngTableParams(
      page: 1
      count: 10
    ,
      total: $scope.users.length
      getData: ($defer, params)->
        # 过滤
        nData = if params.filter() then $filter('filter')($scope.users, params.filter()) else $scope.users
        # 排序
        nData = if params.sorting() then $filter('orderBy')(nData, params.orderBy()) else nData
        # 设置过滤后条数
        params.total(nData.length)
        # 分页
        $defer.resolve(nData.slice((params.page() - 1) * params.count(), params.page() * params.count()))
    )
  $scope.registerEvent('base', CONST.userAll, $scope.userAll)
  $scope.send('base', CONST.userAll)

UserCtrl.$inject = [
  '$scope'
  '$log'
  'ngTableParams'
  '$filter'
]

