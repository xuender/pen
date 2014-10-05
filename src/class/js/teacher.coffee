###
home.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
TeacherCtrl = ($scope, $log, $route)->
  $scope.addHistory($route.current)

TeacherCtrl.$inject = [
  '$scope'
  '$log'
  '$route'
]
