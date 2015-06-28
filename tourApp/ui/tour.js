function Tour($scope, $http) {
    $http.get('/api/tour/2012').
        success(function(data) {
            $scope.tour = data;
        });
    $http.get('/api/gambler/2012').
        success(function(data) {
            $scope.gamblers = data;
        });
    $http.get('/api/news/2012/news').
        success(function(data) {
            $scope.news = data;
        });

}