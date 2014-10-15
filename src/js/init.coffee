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
  'hotkey'
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
])
