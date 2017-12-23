package www

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

func WrapModel(model interface{}) map[string]interface{} {

	wrapper := map[string]interface{}{
		"content": model}

	return wrapper
}

func renderJson(w http.ResponseWriter, status int, data interface{}) {

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)

}
func renderXml(w http.ResponseWriter, status int, data interface{}) {

	w.WriteHeader(status)

	xml.NewEncoder(w).Encode(data)

}

func RenderStatus(w http.ResponseWriter, req *http.Request, status int) {

	Render(w, req, status, http.StatusText(status))

}

func RenderOK(w http.ResponseWriter, req *http.Request, data interface{}) {
	Render(w, req, http.StatusOK, data)
}

func Render(w http.ResponseWriter, req *http.Request, status int, data interface{}) {

	accept := req.Header.Get("Accept")
	//logger.Debug.Printf("Request Accept header : %v", accept)

	if strings.HasPrefix(accept, "application/json") {
		renderJson(w, status, data)
	} else if strings.HasPrefix(accept, "text/xml") {
		renderXml(w, status, data)
	} else {
		panic(fmt.Sprintf("Not supported media type for Header Accept [%v]", accept))

	}
}
