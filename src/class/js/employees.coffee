###
volunteer.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
EmployeesCtrl = ($scope, $log, $route, $modal, ngTableParams, $filter)->
  $scope.addHistory($route.current)
  $scope.ds = []
  $scope.readDs = (data)->
    # 读取datas
    $log.debug('收到义工信息: %s', data)
    if data and 'null' != data
      $scope.ds = JSON.parse(data)
    else
      $scope.ds = []
    if $scope.tableParams
      $scope.tableParams.reload()
      return
    $scope.tableParams = new ngTableParams(
      page: 1
      count: 10
    ,
      total: $scope.ds.length
      getData: ($defer, params)->
        # 过滤
        nData = if params.filter() then $filter('filter')($scope.ds, params.filter()) else $scope.ds
        # 排序
        nData = if params.sorting() then $filter('orderBy')(nData, params.orderBy()) else nData
        # 设置过滤后条数
        params.total(nData.length)
        # 分页
        $defer.resolve(nData.slice((params.page() - 1) * params.count(), params.page() * params.count()))
    )
  $scope.add = ()->
    # 用户编辑
    i = $modal.open(
      templateUrl: 'class/employee.html'
      controller: EmployeeCtrl
      backdrop: 'static'
      keyboard: true
      size: 'lg'
      resolve:
        d: ->
          {}
        pen: ->
          $scope
    )
    i.result.then((d)->
      $log.info '修改'
      $scope.send('class', CLASS.雇员)
    ,->
      $log.info '取消'
    )
  $scope.edit = (d)->
    # 用户编辑
    i = $modal.open(
      templateUrl: 'class/employee.html'
      controller: EmployeeCtrl
      backdrop: 'static'
      keyboard: true
      size: 'lg'
      resolve:
        d: ->
          angular.copy(d)
        pen: ->
          $scope
    )
    i.result.then((d)->
      $log.info '修改'
      $scope.send('class', CLASS.雇员)
    ,->
      $log.info '取消'
    )
  #$scope.registerEvent('class', CLASS.volunteer, (data)->
  #  if 'ok' == data
  #    $scope.current.$edit = false
  #    $scope.send('base', BASE.getDict, $scope.type)
  #)
  $scope.registerEvent('class', CLASS.雇员, $scope.readDs)
  $scope.ready(->
    $scope.send('class', CLASS.雇员)
  )

EmployeesCtrl.$inject = [
  '$scope'
  '$log'
  '$route'
  '$modal'
  'ngTableParams'
  '$filter'
]
