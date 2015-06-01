package server
import (
    "net/http"
    "log"
    . "github.com/zubairhamed/go-commons/network"
)

func NewDefaultHttpServer() (*HttpServer) {
    return &HttpServer{}
}

type HttpServer struct {
    routes     []*Route
}

func (h *HttpServer) Start() {
    h.serveServer()
}

func (h *HttpServer) serveServer() {
    log.Println("Started HTTP Server @ Port 8080")

    wh := &WrappedHandler{
        routes: h.routes,
    }
    http.ListenAndServe(":8080", wh)
}

func (h *HttpServer) NewRoute(path string, method string, fn RouteHandler) *Route {
    route := CreateNewRoute(path, method, fn)
    h.routes = append(h.routes, route)

    return route
}

type WrappedHandler struct {
    routes     []*Route
}

func (wh *WrappedHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    route, attrs, err := MatchingRoute(r.URL.Path, r.Method, nil, wh.routes)

    log.Println(route, attrs, err)
}

