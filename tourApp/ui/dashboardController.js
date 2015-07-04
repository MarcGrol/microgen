'use strict';

angular.module('tourApp')
    .controller('DashboardController', ['$scope', '$stateParams', '$state', '$resource',
        function ($scope, $stateParams, $state, $resource) {

        $scope.tour = {}
        $scope.gamblers = {}
        $scope.news = {}

        $scope.year = $stateParams.year;
          
        console.log("dashboard");

	   $resource('/api/tour/:year', {year:$scope.year})
	   		.get(function(data) {
		        console.log("got tour-data:");
            	$scope.tour = data;
        });

	   $resource('/api/gambler')
	   		.query(function(data) {
		        console.log("got tour-gamblers:");
	            $scope.gamblers = data;
        });

	   $resource('/api/news/:year/news',{year:$scope.year})
	   		.get(function(data) {
		        console.log("got tour-news:");
		        $scope.news = data;
        });

 }]);
