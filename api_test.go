package api

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI(t *testing.T) {
	Convey("If I create an api with endpoint /myapp/api", t, func() {
		api := New("/myapp/api")

		Convey("Then I should ignore requests to /myapp", func() {
			r, _ := http.NewRequest("GET", "/myapp", nil)
			recorder := httptest.NewRecorder()
			api.ServeHTTP(recorder, r)
			fmt.Printf("%v\n", recorder)
			fmt.Println("bob")
			So(recorder.Code, ShouldEqual, http.StatusCreated)

		})

	})
}
