###
dict.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
DictCtrl = ($scope, $log, $route, $routeParams, ngTableParams, $filter)->
  $scope.type = 'type'
  if 'type' of $routeParams
    $scope.type = $routeParams.type
    $scope.addHistory($route.current, $scope.dict_type[$scope.type])
  else
    $scope.addHistory($route.current)

  $scope.dict = []
  $scope.readDict = (data)->
    # 读取字典
    $log.debug('收到字典信息: %s', data)
    if data and 'null' != data
      $scope.dict = JSON.parse(data)
    else
      $scope.dict = []
    if $scope.tableParams
      $scope.tableParams.reload()
      return
    $scope.tableParams = new ngTableParams(
      page: 1
      count: 10
    ,
      total: $scope.dict.length
      getData: ($defer, params)->
        # 过滤
        nData = if params.filter() then $filter('filter')($scope.dict, params.filter()) else $scope.dict
        # 排序
        nData = if params.sorting() then $filter('orderBy')(nData, params.orderBy()) else nData
        # 设置过滤后条数
        params.total(nData.length)
        # 分页
        $defer.resolve(nData.slice((params.page() - 1) * params.count(), params.page() * params.count()))
    )
  update = false
  $scope.$watch('dict', (n, o)->
    update = true
  , true)
  $scope.save = (d)->
    # 保存
    d.$edit = false
    if update
      update = false
      $scope.send('base', CONST.updateDict, d)
      $scope.send('base', CONST.getDict, $scope.type)
  $scope.add = ->
    # 增加
    $log.debug 'add...'
    $scope.dict.push(
      Type: $scope.type
      $edit: true
    )
    $scope.tableParams.reload()
    $log.debug $scope.dict
  $scope.registerEvent('base', CONST.getDict, $scope.readDict)
  $scope.ready(->
    $scope.send('base', CONST.getDict, $scope.type)
  )

DictCtrl.$inject = [
  '$scope'
  '$log'
  '$route'
  '$routeParams'
  'ngTableParams'
  '$filter'
]


