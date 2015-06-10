package pages

type ClientDetailPage struct {
}

func (p *ClientDetailPage) GetContent() []byte {
	return []byte(p.content())
}

func (p *ClientDetailPage) content() string {
	return `
        <html>
            <head>
                <title>Betwixt</title>
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css">
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap-theme.min.css">
                <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
            </head>
            <body role="document">
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
                            <h3>Client: {{ .ClientId }}</h3>
                        </div>

                        <!-- Content Start -->
                        <!-- Each Object -->
                        {{ range $key, $value := .Objects }}
                        <div class="panel panel-primary" width="700">
                            <div class="panel-heading">
                                <h3 class="panel-title" align="left">
                                    <button type="button" class="btn btn-xs btn-info">+ new instance</button> {{ $value.GetDefinition.GetName }}
                                </h3>
                            </div>

                            <div class="panel-body">
                            {{ range $objInstance := $value.GetInstances }}
                                <div class="panel-heading" align="left">
                                    <h4><button type="button" class="btn btn-xs btn-danger">delete</button> Instance #{{ $objInstance }} - /{{ $key }}/{{ $objInstance }}</h4>
                                    <h5>{{ $value.GetDefinition.GetDescription }}</h5>
                                </div>
                                <table class="table table-condensed">
                                    <thead>
                                        <th style="width: 20px;">Path</th>
                                        <th style="width: 100px;">Operations</th>
                                        <th width="400">Name</th>
                                        <th>Description</th>
                                    </thead>
                                    <tbody>
                                        {{ range $resource := $value.GetDefinition.GetResources }}
                                        <tr>
                                            <td>/{{ $key }}/{{ $objInstance }}/{{ $resource.GetId }}</td>
                                            <td>
                                                &nbsp;
                                                {{ if .IsExecutable }}
                                                <button type="button" class="btn btn-xs btn-success">exec</button>
                                                {{ end }}

                                                {{ if .IsReadable }}
                                                <button type="button" class="btn btn-xs btn-primary">observe</button>
                                                <button type="button" class="btn btn-xs btn-primary">stop</button>
                                                |
                                                <button type="button" class="btn btn-xs btn-primary">read</button>
                                                {{ end }}

                                                {{ if .IsWritable }}
                                                <button type="button" class="btn btn-xs btn-warning">write</button>
                                                {{ end }}
                                            </td>
                                            <td>{{ .GetName }}</td>
                                            <td>{{ .GetDescription }}</td>
                                        </tr>
                                        {{ end }}
                                    </tbody>
                                </table>
                            {{ end }}
                            </div>
                        </div>
                        {{ end }}
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

// TODO: func (p *ClientDetailPage) handleRequest()
