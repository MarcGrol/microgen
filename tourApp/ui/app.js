'use strict';

angular.module('webApp', [
    'ui.router',
    'ngResource',
    'ngSanitize',
    'interceptors',
    'ui.bootstrap',
    'ui.select'])
.config(function ($stateProvider) {
    $stateProvider
        .state('error', {
            url: '/error',
            templateUrl: 'error.html'
        });
})
.run(function ($rootScope, $state, ERROR_EVENTS) {
    $rootScope.$on(ERROR_EVENTS.serverError, function () {
        $state.go('error');
    });
});
