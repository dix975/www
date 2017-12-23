package www

import (
	"net/http"
	"fmt"
	"runtime/debug"
	"gopkg.in/mgo.v2"
	"dix975.com/logger"
	"time"
)

func Init() {
}

type Handle struct {
	ControllerFunc func(http.ResponseWriter, *http.Request)
}

func (handle Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {

			if r == mgo.ErrNotFound {
				http.Error(w, http.StatusText(404), 404)

			} else {
				message := fmt.Sprintf("%v\nRecovered panic : %v", http.StatusText(500), r)
				logger.Error.Printf("%s: %s", message, debug.Stack())
				http.Error(w, message, 500)
			}
		}
	}()

	start := time.Now()
	handle.ControllerFunc(w, r)

	fmt.Println("Request time : ", time.Since(start))
}
