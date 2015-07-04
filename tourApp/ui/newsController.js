'use strict';

angular.module('tourApp')
    .controller('NewsController', ['$scope', '$stateParams', '$state', '$resource',
        function ($scope, $stateParams, $state, $resource) {

	console.log("news");
        $scope.year = $stateParams.year;

 }]);
