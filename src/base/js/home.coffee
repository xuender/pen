###
home.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
HomeCtrl = ($scope, $log, $route)->
  $scope.addHistory($route.current)

HomeCtrl.$inject = [
  '$scope'
  '$log'
  '$route'
]
