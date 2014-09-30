###
route.coffee
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
  $routeProvider.when('/',
      templateUrl: 'base/home.html'
      controller: 'HomeCtrl'
      name: '首页'
    ).when('/users',
      templateUrl: 'base/users.html'
      controller: 'UserCtrl'
      name: '用户管理'
    ).otherwise(
      redirectTo: '/'
    )
])
