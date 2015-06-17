package pages

type ClientDetailPage struct {
}

func (p *ClientDetailPage) GetContent() []byte {
	return []byte(p.content())
}

func (p *ClientDetailPage) content() string {
	return `
        <html ng-app="betwixt-app">
            <head>
                <title>Betwixt</title>
                <base href="/">
                <script src="http://ajax.googleapis.com/ajax/libs/angularjs/1.3.13/angular.js"></script>
                <script src="http://angular-ui.github.io/bootstrap/ui-bootstrap-tpls-0.13.0.js"></script>
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css">
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap-theme.min.css">
                <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
                <script>
                    var OPERATION_NONE = 0;
	                var OPERATION_R    = 1;
	                var OPERATION_W    = 2;
	                var OPERATION_RW   = 3;
	                var OPERATION_E    = 4;
	                var OPERATION_RE   = 5;
	                var OPERATION_WE   = 6;
	                var OPERATION_RWE  = 7;

                    angular.module('betwixt-app', ['ui.bootstrap']).config(function($locationProvider) {
                        $locationProvider.html5Mode(true);
                    })

                    angular.module('betwixt-app').controller('BetwixtController', function ($scope, $http, $location) {
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
                            $scope.ClientID = data.Endpoint
                            $scope.Objects = data.Objects
                        });
                    });
                </script>
            </head>
            <body role="document" ng-controller="BetwixtController">
                <!-- Fixed navbar -->
                <nav class="navbar navbar-inverse navbar-fixed-top">
                    <div class="container">
                        <div class="navbar-header">
                          <a class="navbar-brand" href="#">Betwixt LWM2M Server</a>
                        </div>
                        <div id="navbar" class="navbar-collapse collapse">
                          <ul class="nav navbar-nav">
                            <li class="active"><a href="#">Home</a></li>
                          </ul>
                        </div>
                    </div>
                </nav>
                <div class="container theme-showcase" role="main">
                    <br /><br /><br />

                    <div class="row" style="text-align: center">
                        <div class="page-header">
                            <h3>Client: {{ ClientID }}</h3>
                        </div>

                        <!-- Content Start -->
                        <!-- Each Object -->
                        <div class="panel panel-primary" width="700" ng-repeat="(key, value) in Objects">
                            <div class="panel-heading">
                                <h3 class="panel-title" align="left">
                                    <button type="button" class="btn btn-xs btn-info">+ new instance</button> {{ value.Definition.Name }}
                                </h3>
                            </div>

                            <div class="panel-body" ng-repeat="objInstance in value.Instances">
                                <div class="panel-heading" align="left">
                                    <h4><button type="button" class="btn btn-xs btn-danger">delete</button> Instance #{{ objInstance }} - /{{ key }}/{{ objInstance }}</h4>
                                    <h5>{{ value.Definition.Description }}</h5>
                                </div>
                                <table class="table table-condensed">
                                    <thead>
                                        <th style="width: 20px;">Path</th>
                                        <th style="width: 100px;">Operations</th>
                                        <th width="400">Name</th>
                                        <th>Description</th>
                                        <th>&nbsp</th>
                                    </thead>
                                    <tbody>
                                        <tr ng-repeat="resource in value.Definition.Resources">
                                            <td>/{{ key }}/{{ objInstance }}/{{ resource.Id }}</td>
                                            <td>
                                                &nbsp;
                                                <button type="button" class="btn btn-xs btn-success" ng-show="{{ IsExecutable(resource) }}" ng-click="opExecute(ClientID, key, objInstance, resource.Id)">exec</button>

                                                <button type="button" class="btn btn-xs btn-primary" ng-show="{{ IsReadable(resource) }}" ng-click="opObserve(ClientID, key, objInstance, resource.Id)">observe</button>
                                                <button type="button" class="btn btn-xs btn-primary" ng-show="{{ IsReadable(resource) }}" ng-click="opCancelObserve(ClientID, key, objInstance, resource.Id)">stop</button>
                                                <button type="button" class="btn btn-xs btn-primary" ng-show="{{ IsReadable(resource) }}" ng-click="opRead(ClientID, key, objInstance, resource.Id)">read</button>

                                                <button type="button" class="btn btn-xs btn-warning" ng-show="{{ IsWritable(resource) }}" ng-click="opWrite(ClientID, key, objInstance, resource.Id)">write</button>

                                                <button type="button" class="btn btn-xs btn-default" ng-show="{{ IsNone(resource) }}">none</button>
                                            </td>
                                            <td>{{ resource.Name }}</td>
                                            <td>{{ resource.Description }}</td>
                                            <td>{{ resourcevalue['/' + key + '/' + objInstance + '/' + resource.Id] }}</td>
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                        <!-- Content End -->
                    </div>
                </div>

                <!-- Bootstrap core JavaScript
                ================================================== -->
                <!-- Placed at the end of the document so the pages load faster -->
                <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
            </body>
        </html>
    `
}

