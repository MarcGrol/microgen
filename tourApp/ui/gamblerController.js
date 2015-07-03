'use strict';

angular.module('tour')
    .controller('GamblerController', ['$scope', '$stateParams', '$state', '$resource',
        function ($scope, $stateParams, $state, $resource) {

        $scope.param.year = $stateParams.uid;

 }]);