'use strict';

angular.module('tourApp')
    .config(function ($stateProvider, $urlRouterProvider) {
        $urlRouterProvider.otherwise('/tour/2015/dashboard');

        console.log("routing");

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
