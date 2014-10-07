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
        templateUrl: 'class/employees.html'
        controller: 'EmployeesCtrl'
        name: '工作人员'
        url: '/employees'
        menu: true
      }
      {
        templateUrl: 'class/teachers.html'
        controller: 'TeachersCtrl'
        name: '教师'
        url: '/teachers'
        menu: true
      }
      {
        templateUrl: 'class/teacher.html'
        controller: 'TeacherCtrl'
        name: '班级'
        url: '/teacher2'
        menu: true
      }
      {
        templateUrl: 'class/teacher.html'
        controller: 'TeacherCtrl'
        name: '学员'
        url: '/teacher3'
        menu: true
      }
    ]
  }
)

