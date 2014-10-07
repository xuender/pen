module.exports = (grunt)->
  grunt.loadNpmTasks('grunt-contrib-clean')
  grunt.loadNpmTasks('grunt-contrib-uglify')
  grunt.loadNpmTasks('grunt-contrib-watch')
  grunt.loadNpmTasks('grunt-contrib-copy')
  grunt.loadNpmTasks('grunt-karma')
  grunt.loadNpmTasks('grunt-contrib-coffee')
  grunt.loadNpmTasks('grunt-contrib-cssmin')
  grunt.loadNpmTasks('grunt-coffeelint')
  grunt.loadNpmTasks('grunt-bumpx')

  grunt.initConfig(
    pkg:
      grunt.file.readJSON('package.json')
    clean:
      dist: [
        'dist'
        'src_go/public'
      ]
    bump:
      options:
        part: 'patch'
      files: [ 'package.json', 'bower.json']
    copy:
      root:
        files: [
          cwd: 'src'
          src: [
            '**/*.html'
          ]
          dest: 'src_go/public'
          filter: 'isFile'
          expand: true
        ]
      bootstrap:
        files: [
          cwd: 'bower_components/bootstrap/dist'
          src: [
            'css/*.min.css'
            'css/*.map'
            'fonts/*'
            'js/*.min.js'
            'js/*.map'
          ]
          dest: 'src_go/public'
          expand: true
        ]
      angular:
        files: [
          cwd: 'bower_components/angular/'
          src: [
            'angular.js'
            'angular.min.js'
            'angular.min.js.map'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      angularI18n:
        files: [
          cwd: 'bower_components/angular-i18n/'
          src: [
            'angular-locale_zh-cn.js'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      angular_route:
        files: [
          cwd: 'bower_components/angular-route/'
          src: [
            'angular-route.js'
            'angular-route.min.js'
            'angular-route.min.js.map'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      ng_table_js:
        files: [
          cwd: 'bower_components/ng-table/'
          src: [
            'ng-table.js'
            'ng-table.min.js'
            'ng-table.map'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      ng_table_css:
        files: [
          cwd: 'bower_components/ng-table/'
          src: [
            'ng-table.min.css'
          ]
          dest: 'src_go/public/css'
          expand: true
          filter: 'isFile'
        ]
      storage:
        files: [
          cwd: 'bower_components/angular-local-storage/'
          src: [
            'angular-local-storage.min.js'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      md5:
        files: [
          cwd: 'bower_components/blueimp-md5/js'
          src: [
            'md5.min.js'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      jquery:
        files: [
          cwd: 'bower_components/jquery/dist'
          src: [
            'jquery.min.js'
            'jquery.min.map'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      upload:
        files: [
          cwd: 'bower_components/ng-file-upload'
          src: [
            'angular-file-upload-html5-shim.min.js'
            'angular-file-upload.min.js'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      hotkey:
        files: [
          cwd: 'bower_components/ng-hotkey'
          src: [
            'hotkey.min.js'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      socket:
        files: [
          cwd: 'bower_components/ngSocket/dist'
          src: [
            'ngSocket.js'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      ui:
        files: [
          cwd: 'bower_components/angular-bootstrap'
          src: [
            'ui-bootstrap-tpls.min.js'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      fontCss:
        files: [
          cwd: 'bower_components/font-awesome/css'
          src: [
            'font-awesome.min.css'
          ]
          dest: 'src_go/public/css'
          expand: true
          filter: 'isFile'
        ]
      font:
        files: [
          cwd: 'bower_components/font-awesome/fonts'
          src: [
            '*'
          ]
          dest: 'src_go/public/fonts'
          expand: true
          filter: 'isFile'
        ]
      text:
        files: [
          cwd: 'bower_components/textAngular/dist'
          src: [
            'textAngular-sanitize.min.js'
            'textAngular.min.js'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      resize:
        files: [
          cwd: 'bower_components/jquery-resize'
          src: [
            'jquery.ba-resize.min.js'
          ]
          dest: 'src_go/public/js'
          expand: true
          filter: 'isFile'
        ]
      css:
        files: [
          cwd: 'src/css'
          src: [
            'main.css'
          ]
          dest: 'src_go/public/css'
          expand: true
          filter: 'isFile'
        ]
    coffee:
      options:
        bare: true
      main:
        files:
          'src_go/public/js/main.min.js': [
            'src/js/init.coffee'
            'src/js/utils.coffee'
            'src/js/pen.coffee'
            'src/base/js/init.coffee'
            'src/base/js/login.coffee'
            'src/base/js/home.coffee'
            'src/base/js/user.coffee'
            'src/base/js/dict.coffee'
            'src/base/js/objects.coffee'
            'src/class/js/init.coffee'
            'src/class/js/employee.coffee'
            'src/class/js/teacher.coffee'
          ]
    uglify:
      main:
        files:
          'dist/go/public/js/main.min.js': [
            'src_go/public/js/main.min.js'
          ]
    cssmin:
      toolbox:
        expand: true
        cwd: 'src_go/public/css/'
        src: ['*.css', '!*.min.css'],
        dest: 'dist/css/'
        #ext: '.min.css'
    watch:
      css:
        files: [
          'src/css/*.css'
        ]
        tasks: ['copy:css']
      html:
        files: [
          'src/**/*.html'
        ]
        tasks: ['copy:root']
      coffee:
        files: [
          'src/**/*.coffee'
        ]
        tasks: ['coffee']
    karma:
      options:
        configFile: 'karma.conf.js'
      dev:
        colors: true,
      travis:
        singleRun: true
        autoWatch: false
  )
  grunt.registerTask('test', ['karma'])
  grunt.registerTask('dev', [
    'clean'
    'copy',
    'coffee'
  ])
  grunt.registerTask('dist', [
    'dev'
    'uglify'
  ])
  grunt.registerTask('deploy', [
    'bump'
    'dist'
  ])
  grunt.registerTask('travis', 'travis test', ['karma:travis'])
  grunt.registerTask('default', ['dist'])
