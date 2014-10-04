###
route.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
CONST =
  login:       0
  logout:      1
  count:       2
  userAll:     3
  dict:        4
  dictVer:     5
  getDict:     6
  updateDict:  7
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
      controller: 'UserCtrl'
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
