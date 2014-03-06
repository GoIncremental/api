package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	m := Server()

	Convey("The site should render hello world", t, func() {
		recorder := httptest.NewRecorder()
		m.ServeHTTP(recorder, r)

		So(recorder.Code, ShouldEqual, http.StatusOK)
		So(recorder.Body.String(), ShouldEqual, "Hello world!")
	})
}
