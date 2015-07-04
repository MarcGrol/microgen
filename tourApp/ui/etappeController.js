'use strict';

angular.module('tourApp')
    .controller('EtappeController', ['$scope', '$stateParams', '$state', '$resource',
        function ($scope, $stateParams, $state, $resource) {

        console.log("etappe");
        $scope.year = $stateParams.year;

 }]);
