###
init.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
MENUS.push(
  {
    name: '学习班'
    order: 10
    routes: [
      {
        templateUrl: 'class/teacher.html'
        controller: 'TeacherCtrl'
        name: '老师信息'
        url: '/teacher'
        menu: true
      }
      {
        templateUrl: 'class/teacher.html'
        controller: 'TeacherCtrl'
        name: '义工信息'
        url: '/teacher1'
        menu: true
      }
      {
        templateUrl: 'class/teacher.html'
        controller: 'TeacherCtrl'
        name: '班级信息'
        url: '/teacher2'
        menu: true
      }
      {
        templateUrl: 'class/teacher.html'
        controller: 'TeacherCtrl'
        name: '学员信息'
        url: '/teacher3'
        menu: true
      }
    ]
  }
)

