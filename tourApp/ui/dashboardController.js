'use strict';

angular.module('tourApp')
    .controller('DashboardController', ['$scope', '$stateParams', '$state', '$resource',
        function ($scope, $stateParams, $state, $resource) {

        $scope.year = $stateParams.year;
          
        console.log("dashboard");

 }]);
