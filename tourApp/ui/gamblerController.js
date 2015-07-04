'use strict';

angular.module('tourApp')
    .controller('GamblerController', ['$scope', '$stateParams', '$state', '$resource',
        function ($scope, $stateParams, $state, $resource) {

        console.log("gambler");
        $scope.year = $stateParams.uid;

 }]);
