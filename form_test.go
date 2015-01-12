package api

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPINew(t *testing.T) {

	Convey("API exports New() function", t, func() {

		api, err := New()

		Convey("If okay, error is nil", func() {

			So(err, ShouldBeNil)
			So(api, ShouldNotBeNil)

		})

	})

}

func TestFormAPI(t *testing.T) {
	Convey("Form Creation", t, func() {

		api, _ := New()

		Convey("When I recieve a Post to /v1/form", func() {

			req, _ := http.NewRequest("POST", "/v1/form", nil)
			recorder := httptest.NewRecorder()

			api.ServeHTTP(recorder, req)
			Convey("I recieve a 201 Created Status Code", func() {

				So(recorder.Code, ShouldEqual, http.StatusCreated)

			})

			Convey("I recieve a JSON response", func() {

				So(recorder.HeaderMap["Content-Type"], ShouldNotBeNil)
				So(recorder.HeaderMap.Get("Content-Type"), ShouldEqual, "application/json")

				decoder := json.NewDecoder(recorder.Body)
				var result PostResult
				decoder.Decode(&result)

				Convey("JSON Response contains _ID", func() {

					So(result.ID, ShouldNotEqual, "")

				})
				Convey("ID is a valid Hex ID", func() {

					So(result.ID.Valid(), ShouldBeTrue)

				})

			})
		})

		Convey("When I recieve a GET to /v1/", func() {

			req, _ := http.NewRequest("GET", "/v1/form", nil)
			recorder := httptest.NewRecorder()

			api.ServeHTTP(recorder, req)

			Convey("I recieved a not found code (404)", func() {

				So(recorder.Code, ShouldEqual, http.StatusNotFound)

			})

		})
	})
}
