package helpers

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// Render data to client
func Render(res http.ResponseWriter, data interface{}) {
	
	responseHeaders(res)

	jsonString, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(res, string( jsonString ) )
}

func JsonDecode(res http.ResponseWriter, req *http.Request) map[string] interface{} {

	decode := json.NewDecoder(req.Body)
	body := make( map[string] interface{} )

	if err := decode.Decode(&body); err != nil{
		panic(err)
	}

	defer req.Body.Close()

	return body
}

// Set all response header
func responseHeaders(res http.ResponseWriter) {

	res.Header().Set("Content-Type", "application/json")

}