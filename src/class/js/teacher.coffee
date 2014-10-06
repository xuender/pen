###
home.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
TeacherCtrl = ($scope, $log, $route, $routeParams, ngTableParams, $filter)->
  $scope.addHistory($route.current)
  $scope.teachers = []

TeacherCtrl.$inject = [
  '$scope'
  '$log'
  '$route'
  '$routeParams'
  'ngTableParams'
  '$filter'
]
