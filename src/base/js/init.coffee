###
init.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###

BASE =
  msg:         0
  login:       1
  logout:      2
  count:       3
  userAll:     4
  updateUser:  5
  dict:        6
  dictVer:     7
  getDict:     8
  updateDict:  9
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
        controller: 'UsersCtrl'
        name: '用户管理'
        url: '/users'
        menu: true
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
