var gulp = require("gulp")
var sass = require("gulp-sass")
var gulpUtil = require('gulp-util');
var uglify = require('gulp-uglify')
var inlinejs = require("gulp-inline-js")

gulp.task("sass", function(){
    gulp.src("scss/**/*.scss")
        .pipe(sass())
        .pipe(gulp.dest("../css"));
})

gulp.task("js", function(){
    gulp.src("js/**/*.js")
        .pipe(inlinejs())
        .pipe(uglify().on("error",gulpUtil.log))
        .pipe(gulp.dest("../js/"))
})

gulp.task("watchjs", function(){
    gulp.watch("js/**/*.js",function(){
        gulp.run("js");
    })
})

gulp.task("watchcss", function(){
    gulp.watch("scss/**/*.scss", function(){
        gulp.run("sass");
    })
})

gulp.task("default",function(){
    gulp.run("watchjs","watchcss","sass","js");
})