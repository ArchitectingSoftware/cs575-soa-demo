var gulp = require('gulp');
var ts = require('gulp-typescript');
var nodemon = require('gulp-nodemon');
var sourcemaps = require('gulp-sourcemaps');
var merge = require('merge2');

//  var tsProject = ts.createProject({
//      declaration: true
//  });

var tsProject = ts.createProject('tsconfig.json', { noImplicitAny: true });

gulp.task('ts', function() {
    var tsResult = tsProject.src()
        .pipe(sourcemaps.init())
        .pipe(tsProject());

    return merge([ // Merge the two output streams, so this task is finished when the IO of both operations is done.
        tsResult.dts.pipe(gulp.dest('release/definitions')),
        tsResult.js.pipe(sourcemaps.write("./", { sourceRoot: __dirname }))
            .pipe(gulp.dest('dest'))
    ]);
});

gulp.task('views', function() {
    console.log("jade fn")
    return gulp.src('views/*.jade')
        .pipe(gulp.dest('dest/views'))
});

gulp.task('public-files', function() {
    console.log("public fn")
    return gulp.src('public/**/*')
        .pipe(gulp.dest('dest/public'))
});

gulp.task('watch', ['ts','views','public-files'], function() {
    console.log("watch fn")
    gulp.watch('**/*.ts', ['ts']);
    gulp.watch('**/*.jade', ['views']);
    gulp.watch('public/**/*', ['public-files']);
});

gulp.task('build', ['ts','views','public-files']);

gulp.task('nodemon', ['watch'], function() {
    nodemon({script: 'dest/app.js'});
});
