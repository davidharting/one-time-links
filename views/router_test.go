package views

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_getRoute(t *testing.T) {
	homeIndexRequest := createRequest(t, http.MethodGet, "/")
	createMessageRequest := createRequest(t, http.MethodPost, "/")
	showMessageRequestMissingEndingSlash := createRequest(t, http.MethodGet, "/message")
	showMessageRequest := createRequest(t, http.MethodGet, "/message/")
	showMessageUnsupportedMethod := createRequest(t, http.MethodPut, "/message/")
	notFoundRequest := createRequest(t, http.MethodGet, "/acf3sefe")

	tests := []struct {
		name string
		r    *http.Request
		want route
	}{
		{name: "GET /", want: routeHomeIndex, r: homeIndexRequest},
		{name: "POST /", want: routeCreateMessage, r: createMessageRequest},
		{name: "GET /message", want: routeShowMessage, r: showMessageRequestMissingEndingSlash},
		{name: "GET /message/", want: routeShowMessage, r: showMessageRequest},
		{name: "PUT /message/", want: routeBadRequest, r: showMessageUnsupportedMethod},
		{name: "GET /acf3sefe/", want: routeNotFound, r: notFoundRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRoute(tt.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func createRequest(t *testing.T, method string, path string) *http.Request {
	r, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.FailNow()
	}
	return r
}

func Test_getPath(t *testing.T) {
	tests := []struct {
		name string
		r    *http.Request
		want string
	}{
		{name: "/", r: createRequest(t, http.MethodGet, "/"), want: "/"},
		{name: "/message", r: createRequest(t, http.MethodGet, "/message"), want: "/message/"},
		{name: "/message/", r: createRequest(t, http.MethodGet, "/message/"), want: "/message/"},
		{name: "/message/?id=1", r: createRequest(t, http.MethodGet, "/message/?id=1"), want: "/message/"},
		{name: "/message?id=1", r: createRequest(t, http.MethodGet, "/message?id=1"), want: "/message/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPath(tt.r); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
