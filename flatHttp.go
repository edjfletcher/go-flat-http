package go_flat_http

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type FlatHttp struct {
	*mux.Router
}

func NewFlatHttp() FlatHttp {
	return FlatHttp{mux.NewRouter().StrictSlash(true)}
}

func (f *FlatHttp) AddFlatHandler(h HandlerInterface, name string) {
	f.HandleFunc(name, handlerFactory(h)).Methods(http.MethodPost, http.MethodOptions)
}

func handlerFactory(applicationHandler HandlerInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == http.MethodOptions {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			return
		}

		reqBody, err := ioutil.ReadAll(r.Body)

		var response Response
		if err != nil {
			response = (Response{}).WithResponseCode(http.StatusInternalServerError)
		} else {
			response = applicationHandler(r.Context(), reqBody, Response{})
		}

		data := getCorrectData(response)

		responseFormatAsJson(w, response.ResponseCode(), data)
	}
}

func getCorrectData(response Response) interface{} {
	if response.HasErr() {
		return ErrorStruct{
			Message: response.Err().Error(),
		}
	} else {
		return response.Data()
	}
}

type ErrorStruct struct {
	Message string `json:"message"`
}
