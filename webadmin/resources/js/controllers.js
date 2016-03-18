angular.module('betwixt-app', []).config(function($locationProvider) {
    $locationProvider.html5Mode(true);
})

var app = angular.module("betwixt-app");

app.run(["$http","$rootScope", "$interval", function($http, $rootScope, $interval) {

    $rootScope.stats = {
        "MemUsage": "0",
        "ClientsCount": "0",
        "Requests": "0",
        "Errors": "0"
    }

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

app.controller("clientController", function($scope, $http, $location){
    $scope.resourcevalue = {}

    $scope.IsExecutable = function (o) {
        op = o.Operations
        return (op == OPERATION_E || op == OPERATION_RE || op == OPERATION_RWE || op == OPERATION_WE)
    }

    $scope.IsReadable = function (o) {
        op = o.Operations
        return (op == OPERATION_RE || op == OPERATION_R || op == OPERATION_RWE || op == OPERATION_RW)
    }

    $scope.IsWritable = function (o) {
        op = o.Operations
        return (op == OPERATION_RW || op == OPERATION_RWE || op == OPERATION_WE || op == OPERATION_W)
    }

    $scope.IsNone = function (o) {
        op = o.Operations
        return (op == OPERATION_NONE)
    }

    $scope.opExecute = function (client, object, instance, resource) {
        alert("Execute");
        // POST     /api/clients/{client}/{object}/{instance}/{resource}
    }

    $scope.opExecuteWithOptions = function (client, object, instance, resource) {
        alert("Execute with Options");
        // POST     /api/clients/{client}/{object}/{instance}/{resource}
    }

    $scope.opRead = function (client, object, instance, resource) {
        key = "/" + object + "/" + instance + "/" + resource;
        $http.get("/api/clients/" + client + key).success(function(data) {
            out = ""
            for (i=0; i < data.Content.length; i++) {
                v = data.Content[i]
                out += v.Value + " "
            }
            $scope.resourcevalue[key] = out;
        });
    }

    $scope.opObserve = function (client, object, instance, resource) {
        alert("Observe");
        // POST     /api/clients/{client}/{object}/{instance}/{resource}/observe
    }

    $scope.opCancelObserve = function (client, object, instance, resource) {
        alert("Cancel");
        // DELETE   /api/clients/{client}/{object}/{instance}/{resource}/observe
    }

    $scope.opWrite = function (client, object, instance, resource) {
        alert("Write");
        // PUT      /api/clients/{client}/{object}/{instance}/{resource}
    }

    $scope.opDelete = function (client, object, instance) {
        // DELETE   /api/clients/{client}/{object}/{instance}
    }

    $scope.opCreate = function() {
        // POST     /api/clients/{client}/{object}/{instance}
    }

    clientId = $location.path().split("/")[2];
    $http.get("/api/clients/" + clientId).success(function(data) {
        $scope.client = data
    });
});

app.controller("settingsController", function($scope, $http){

});

app.controller("logsController", function($scope, $http){

});
