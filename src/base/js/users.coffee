###
users.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
UsersCtrl = ($scope, $log, $route, $modal, ngTableParams, $filter)->
  $scope.addHistory($route.current)
  $scope.users = []
  $scope.userAll = (data)->
    # 获取所有用户
    #$log.debug('收到用户信息: %s', data)
    if data and 'null' != data
      $scope.users = JSON.parse(data)
    else
      $scope.users = []
    if $scope.tableParams
      $scope.tableParams.reload()
      return
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
  $scope.edit = (user)->
    # 用户编辑
    i = $modal.open(
      templateUrl: 'base/user.html'
      controller: UserCtrl
      backdrop: 'static'
      keyboard: true
      size: 'lg'
      resolve:
        user: ->
          angular.copy(user)
        gender: ->
          $scope.dict_gender
        pen: ->
          $scope
    )
    i.result.then((user)->
      $log.info '修改'
      $scope.send('base', BASE.用户列表)
    ,->
      $log.info '取消'
    )
  $scope.registerEvent('base', BASE.用户列表, $scope.userAll)
  $scope.ready(->
    $scope.send('base', BASE.用户列表)
  )

UsersCtrl.$inject = [
  '$scope'
  '$log'
  '$route'
  '$modal'
  'ngTableParams'
  '$filter'
]

