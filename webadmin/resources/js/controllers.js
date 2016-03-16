var app = angular.module("betwixt-app", [])

app.run(["$http","$rootScope", function($http, $rootScope) {
    $http.get("/api/server/stats").success(function(data){
        $rootScope.stats = data;
    })
}]);

app.controller("indexController", function($scope, $http){
    $http.get("/api/clients").success(function(data){
        $scope.clients = data;
    })
});

app.controller("clientController", function($scope){

});

