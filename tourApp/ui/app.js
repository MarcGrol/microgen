'use strict';

angular.module('tourApp', [
    'ui.router',
    'ngResource',
 //   'ngSanitize',
//    'ui.bootstrap',
 //   'ui.select'
])
.config(function ($stateProvider) {
    $stateProvider
        .state('error', {
            url: '/error',
            templateUrl: 'error.html'
        });
})
;
