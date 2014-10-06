###
init.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
angular.module('pen', [
  'ngRoute'
  'ui.bootstrap'
  'ngSocket'
  'LocalStorageModule'
  'ngTable'
  #'hotkey'
  #'angularFileUpload'
  #'textAngular'
]).config(['$routeProvider', ($routeProvider)->
  $routeProvider.otherwise(
    redirectTo: '/'
  )
  for ms in MENUS
    console.info ms.name
    for m in ms.routes
      console.info m.name
      $routeProvider.when(m.url, m)

  #$routeProvider.when('/',
  #    templateUrl: 'base/home.html'
  #    controller: 'HomeCtrl'
  #    name: '首页'
  #  ).when('/users',
  #    templateUrl: 'base/users.html'
  #    controller: 'UsersCtrl'
  #    name: '用户管理'
  #  ).when('/dict',
  #    templateUrl: 'base/dict.html'
  #    controller: 'DictCtrl'
  #    name: '字典管理'
  #  ).when('/dict/:type',
  #    templateUrl: 'base/dict.html'
  #    controller: 'DictCtrl'
  #    name: '字典明细'
  #  ).otherwise(
  #    redirectTo: '/'
  #  )
])
