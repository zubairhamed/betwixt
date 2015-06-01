package pages

type HomePage struct {
}

func (p *HomePage) GetContent() []byte {
	return []byte(p.content())
}

func (p *HomePage) content() string {
	return `
        <html>
            <head>
                <title>Registered Clients</title>
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css">
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap-theme.min.css">
                <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
            </head>
            <body role="document">
                <!-- Fixed navbar -->
                <nav class="navbar navbar-inverse navbar-fixed-top">
                  <div class="container">
                    <div class="navbar-header">
                      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                      </button>
                      <a class="navbar-brand" href="#">Title</a>
                    </div>
                    <div id="navbar" class="navbar-collapse collapse">
                      <ul class="nav navbar-nav">
                        <li class="active"><a href="#">Home</a></li>
                        <li><a href="#item1">Item 1</a></li>
                        <li><a href="#contact">Contact</a></li>
                        <li class="dropdown">
                          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">Dropdown <span class="caret"></span></a>
                          <ul class="dropdown-menu" role="menu">
                            <li><a href="#">Action</a></li>
                            <li><a href="#">Another action</a></li>
                            <li><a href="#">Something else here</a></li>
                            <li class="divider"></li>
                            <li class="dropdown-header">Nav header</li>
                            <li><a href="#">Separated link</a></li>
                            <li><a href="#">One more separated link</a></li>
                          </ul>
                        </li>
                      </ul>
                    </div><!--/.nav-collapse -->
                  </div>
                </nav>

                <div class="container theme-showcase" role="main">
                  <br /><br />

                  <div class="page-header">
                    <h1>Registered Clients</h1>
                  </div>
                  <div class="row">

                    <div class="col-md-12">
                      <table class="table">
                        <thead>
                          <tr>
                            <th>Endpoint</th>
                            <th>Registration ID</th>
                            <th>Registration Date</th>
                            <th>Last Update</th>
                            <th>Actions</th>
                          </tr>
                        </thead>
                        <tbody>


                          {{ range . }}
                          <tr>
                            <td>{{.Endpoint}}</td>
                            <td>{{.RegistrationID}}</td>
                            <td>{{.RegistrationDate}}</td>
                            <td>{{.LastUpdate}}</td>
                            <td>
                              <h4>
                                <span class="label label-info">View</span>
                                <span class="label label-danger">Remove</span>
                              </h4>
                            </td>
                          </tr>
                          {{ end }}

                        </tbody>
                      </table>
                    </div>

                  </div>

                </div> <!-- /container -->

                <!-- Bootstrap core JavaScript
                ================================================== -->
                <!-- Placed at the end of the document so the pages load faster -->
                <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
            </body>
        </html>
    `
}
