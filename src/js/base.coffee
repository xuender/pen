###
route.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
penApp = angular.module('pen', [
  'ngRoute'
  'ui.bootstrap'
  'ngSocket'
  'LocalStorageModule'
  'ngTable'
  #'hotkey'
  #'angularFileUpload'
  #'textAngular'
])
penApp.config(['$routeProvider', ($routeProvider)->
  $routeProvider.
    when('/',
      templateUrl: 'base/home.html'
      controller: 'HomeCtrl'
    ).when('/users',
      templateUrl: 'base/users.html'
      controller: 'UserCtrl'
    ).otherwise({
      redirectTo: '/'
    })
])
