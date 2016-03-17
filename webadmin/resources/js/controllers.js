var app = angular.module("betwixt-app", [])

app.run(["$http","$rootScope", "$interval", function($http, $rootScope, $interval) {

    $interval(function(){
        $http.get("/api/server/stats").success(function(data){
            $rootScope.stats = data;
        })
    }, 2000)
}]);

app.controller("indexController", function($scope, $http, $interval){
    $interval(function(){
        $http.get("/api/clients").success(function(data){
            $scope.clients = data;
        })
    }, 3000)
});

app.controller("clientController", function($scope, $http){

});

app.controller("settingsController", function($scope, $http){

});

app.controller("logsController", function($scope, $http){

});
