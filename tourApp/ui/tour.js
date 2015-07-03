'use strict';

function Tour($scope, $http) {
    $http.get('/api/tour/2015').
        success(function(data) {
            $scope.tour = data;
        });
    $http.get('/api/gambler').
        success(function(data) {
            $scope.gamblers = data;
        });
    $http.get('/api/news/2015/news  ').
        success(function(data) {
            $scope.news = data.newsItems;
        });

}
