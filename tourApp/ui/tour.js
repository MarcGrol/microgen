function Tour($scope, $http) {
    $http.get('/api/tour/2015').
        success(function(data) {
            $scope.tour = data;
        });
}