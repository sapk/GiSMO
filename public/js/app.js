// For any third party dependencies, like jQuery, place them in the lib folder.

// Configure loading modules from the lib directory,
// except for 'app' ones, which are in a sibling
// directory.
requirejs.config({
    baseUrl: 'assets',
    paths: {
        app: 'js',
        lib: 'lib',
        jquery: 'lib/jquery/jquery-2.1.4.min',
        underscore: 'lib/underscore/underscore-min',
        backbone: 'lib/backbone/backbone-min'
    },
    shim: {
        jquery: {
            exports: '$'
        },
        underscore: {
            exports: '_'
        },
        backbone: {
            deps: [
                'underscore',
                'jquery'
            ],
            exports: 'Backbone'
        }
    }
});


// Start loading the main app file. Put all of
// your application logic in there.
requirejs(['backbone'], function (Backbone) {
    requirejs(['app/main']);
});