###
init.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
MENUS = [
  {
    name: '系统管理'
    order: 3000
    routes: [
      {
        templateUrl: 'base/home.html'
        controller: 'HomeCtrl'
        name: '首页'
        url: '/'
        menu: false
      }
      {
        templateUrl: 'base/users.html'
        controller: 'ObjectsCtrl'
        name: '用户管理'
        url: '/users'
        menu: true
        object: {
          code: 'base'
          templateUrl: 'base/user.html'
          controller: 'EmployeeCtrl'
          getId: BASE.用户列表
        }
      }
      {
        templateUrl: 'base/dict.html'
        controller: 'DictCtrl'
        name: '字典管理'
        url: '/dict'
        menu: true
      }
      {
        templateUrl: 'base/dict.html'
        controller: 'DictCtrl'
        name: '字典明细'
        url: '/dict/:type'
        menu: false
      }
    ]
  }
]
