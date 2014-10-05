###
route.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
CONST =
  msg:         0
  login:       1
  logout:      2
  count:       3
  userAll:     4
  dict:        5
  dictVer:     6
  getDict:     7
  updateDict:  8
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
  $routeProvider.when('/',
      templateUrl: 'base/home.html'
      controller: 'HomeCtrl'
      name: '首页'
    ).when('/users',
      templateUrl: 'base/users.html'
      controller: 'UsersCtrl'
      name: '用户管理'
    ).when('/dict',
      templateUrl: 'base/dict.html'
      controller: 'DictCtrl'
      name: '字典管理'
    ).when('/dict/:type',
      templateUrl: 'base/dict.html'
      controller: 'DictCtrl'
      name: '字典明细'
    ).otherwise(
      redirectTo: '/'
    )
])
