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
          getId: CLASS.GYCX
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
          getId: CLASS.JSCX
        }
      }
      {
        templateUrl: 'class/classes.html'
        controller: 'ObjectsCtrl'
        name: '班级'
        url: '/class'
        menu: true
        object: {
          code: 'class'
          templateUrl: 'class/class.html'
          controller: 'ClassCtrl'
          getId: CLASS.BJCX
        }
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
          getId: CLASS.XYCX
        }
      }
    ]
  }
)

