'use strict';

angular.module('tour')
    .controller('EtappeController', ['$scope', '$stateParams', '$state', '$resource',
        function ($scope, $stateParams, $state, $resource) {

        $scope.param.year = $stateParams.year;

 }]);