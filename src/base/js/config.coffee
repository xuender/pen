###
dict.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
ConfigCtrl = ($scope, $log, $route)->
  $scope.addHistory($route.current)

  $scope.config = {}
  $scope.readConfig = (data)->
    # 配置查询
    $log.debug('收到配置信息: %s', data)
    if data and 'null' != data
      $scope.config = JSON.parse(data)
    else
      $scope.config = {}
  $scope.dbInit = ->
    # 数据库初始化
    $scope.isSend = true
    $scope.send('base', BASE.数据库初始化)
  $scope.save = ->
    # 保存
    $scope.isSend = true
    $scope.send('base', BASE.配置编辑, $scope.config)
  $scope.registerEvent('base', BASE.配置编辑, (data)->
    if 'ok' == data
      $scope.isSend = false
      $scope.alert('保存成功')
      $scope.send('base', BASE.配置查询)
  )
  $scope.registerEvent('base', BASE.数据库初始化, (data)->
    if 'ok' == data
      $scope.isSend = false
      $scope.alert('数据库初始化完毕')
  )
  $scope.registerEvent('base', BASE.配置查询, $scope.readConfig)
  $scope.ready(->
    $scope.send('base', BASE.配置查询)
  )

ConfigCtrl.$inject = [
  '$scope'
  '$log'
  '$route'
]


