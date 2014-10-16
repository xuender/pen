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
        controller: 'ObjectsCtrl'
        name: '工作人员'
        url: '/employees'
        menu: true
        object: {
          code: 'class'
          templateUrl: 'class/employee.html'
          controller: 'EmployeeCtrl'
          getId: CLASS.雇员查询
        }
      }
      {
        templateUrl: 'class/teachers.html'
        controller: 'ObjectsCtrl'
        name: '教师'
        url: '/teachers'
        menu: true
        object: {
          code: 'class'
          templateUrl: 'class/teacher.html'
          controller: 'TeacherCtrl'
          getId: CLASS.教师查询
        }
      }
      {
        templateUrl: 'class/teacher.html'
        controller: 'TeacherCtrl'
        name: '班级'
        url: '/teacher2'
        menu: true
      }
      {
        templateUrl: 'class/students.html'
        controller: 'ObjectsCtrl'
        name: '学员'
        url: '/students'
        menu: true
        object: {
          code: 'class'
          templateUrl: 'class/student.html'
          controller: 'StudentCtrl'
          getId: CLASS.学员查询
        }
      }
    ]
  }
)

