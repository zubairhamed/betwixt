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
                    angular.module('betwixt-app', ['ui.bootstrap']).config(function($locationProvider) {
                        $locationProvider.html5Mode(true);
                    })
                    angular.module('betwixt-app').controller('BetwixtController', function ($scope, $http, $location) {
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

                            <div class="panel-body">
                                <div class="panel-heading" align="left" ng-repeat="objInstance in value.Instances">
                                    <h4><button type="button" class="btn btn-xs btn-danger">delete</button> Instance #{{ objInstance }} - /{{ key }}/{{ objInstance }}</h4>
                                    <h5>{{ value.Definition.Description }}</h5>
                                </div>
                                <table class="table table-condensed">
                                    <thead>
                                        <th style="width: 20px;">Path</th>
                                        <th style="width: 100px;">Operations</th>
                                        <th width="400">Name</th>
                                        <th>Description</th>
                                    </thead>
                                    <tbody>
                                        <tr ng-repeat="resource in value.Definition.Resources">
                                            <td>/{{ key }}/{{ objInstance }}/{{ resource.Id }}</td>
                                            <td>
                                                &nbsp;
                                                <!-- TODO -->
                                                {{ IsExecutable }}
                                                <button type="button" class="btn btn-xs btn-success">exec</button>
                                                {{ end }}

                                                {{ IsReadable }}
                                                <button type="button" class="btn btn-xs btn-primary">observe</button>
                                                <button type="button" class="btn btn-xs btn-primary">stop</button>
                                                |
                                                <button type="button" class="btn btn-xs btn-primary">read</button>
                                                {{ end }}

                                                {{ IsWritable }}
                                                <button type="button" class="btn btn-xs btn-warning">write</button>
                                                {{ end }}
                                            </td>
                                            <td>{{ resource.Name }}</td>
                                            <td>{{ resource.Description }}</td>
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

