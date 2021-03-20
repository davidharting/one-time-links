package views

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("view router\t")

	route := getRoute(r)
	log.Println(fmt.Sprintf("Determined route to be %v", route))

	switch route {
	case routeHomeIndex:
		homeIndex(w, r, make(map[string]string))
	case routeCreateMessage:
		createMessage(w, r)
	case routeShowMessage:
		showMessage(w, r)
	case routeNotFound:
		render(w, "not_found", make(map[string]string))
	case routeBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func normalizePath(path string) string {
	if strings.HasSuffix(path, "/") {
		return path
	}
	return fmt.Sprintf("%v/", path)
}

func getPath(r *http.Request) string {
	return normalizePath(r.URL.EscapedPath())
}

type route = string

const (
	routeHomeIndex     route = "homeIndex"
	routeCreateMessage route = "createMessage"
	routeShowMessage   route = "showMessage"
	routeNotFound      route = "notFound"
	routeBadRequest    route = "badRequest"
)

func getRoute(r *http.Request) route {
	path := getPath(r)
	method := r.Method

	if path == "/message/" && method == http.MethodGet {
		return routeShowMessage
	}

	if path == "/" && method == http.MethodGet {
		return routeHomeIndex
	}

	if path == "/" && method == http.MethodPost {
		return routeCreateMessage
	}

	if method == http.MethodGet {
		return routeNotFound
	}
	return routeBadRequest
}
