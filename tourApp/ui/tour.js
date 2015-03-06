function Tour($scope, $http) {
    $http.get('/api/tour/2012').
        success(function(data) {
            $scope.tour = data;
        });
}