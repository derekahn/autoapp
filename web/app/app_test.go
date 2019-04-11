package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("it should return Server", func(t *testing.T) {
		s := Server{
			Router: http.NewServeMux(),
			Welcome: &Welcome{
				Name: "Test",
				Time: time.Now().Format(time.Stamp),
			},
		}

		got := reflect.TypeOf(&s)
		expect := reflect.TypeOf(&Server{})
		equals(t, expect, got)
	})
}

func TestAllRoutes(t *testing.T) {
	s := Server{
		Router: http.NewServeMux(),
		Welcome: &Welcome{
			Name: "Test",
			Time: time.Now().Format(time.Stamp),
		},
	}
	s.Routes("../")

	name := "King Kamehameha"
	index := fmt.Sprintf("%s?name=%s", indexRoute, name)
	routes := [3]string{index, staticRoute, faviconRoute}

	for _, r := range routes {
		t.Run(fmt.Sprintf("it should return a '200' on 'GET' %s", r), func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, r, nil)
			ok(t, err)

			res := httptest.NewRecorder()
			s.Router.ServeHTTP(res, req)

			if str := os.Getenv("FAILED"); str == "true" {
				equals(t, http.StatusBadGateway, res.Code)
			} else {
				equals(t, http.StatusOK, res.Code)
			}
		})
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
