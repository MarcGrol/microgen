'use strict';

angular.module('tour')
    .config(function ($stateProvider, $urlRouterProvider) {
        $urlRouterProvider.otherwise('/tour/2015/dashboard');

        $stateProvider
            .state('dashboard', {
                url: '/tour/:year/dashboard',
                templateUrl: 'dashboard.html',
                controller: 'DashboardController'
            })
            .state('gambler', {
                url: '/tour/:year/gambler',
                templateUrl: 'gambler.html',
                controller: 'GamblerController'
            })
            .state('etappe', {
                url: '/tour/:year/etappe/:etappeId',
                templateUrl: 'etappe.html',
                controller: 'EtappeController'
            })
            .state('news', {
                url: '/tour/:year/news',
                templateUrl: 'news.html',
                controller: 'NewsController'
            });

    });
