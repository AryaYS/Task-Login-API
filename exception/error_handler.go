package exception

import (
	"apilogin/model/response"
	"encoding/json"
	"net/http"
)

func ErrorHandler(wr http.ResponseWriter, req *http.Request, err interface{}) {
	if unauthorizedfunc(wr, req, err) {
		return
	}

	if notFoundError(wr, req, err) {
		return
	}
	if BadRequestError(wr, req, err) {
		return
	}
	internalServerError(wr, req, err)
}

func unauthorizedfunc(wr http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, er := err.(Unauthorized)
	if er {
		wr.Header().Set("Content-Type", "application/json")

		webResponse := response.WebResponse{
			Code:   http.StatusConflict,
			Status: "Unauthorized",
			Data:   exception.Error,
		}
		wr.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(wr)
		err = encoder.Encode(webResponse)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

func BadRequestError(wr http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, er := err.(BadRequest)
	if er {
		wr.Header().Set("Content-Type", "application/json")

		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Something Wrong with Request",
			Data:   exception.Error,
		}
		wr.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(wr)
		err = encoder.Encode(webResponse)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

func notFoundError(wr http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, er := err.(NotFoundError)
	if er {
		wr.Header().Set("Content-Type", "application/json")

		webResponse := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found Any User",
			Data:   exception.Error,
		}
		wr.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(wr)
		err = encoder.Encode(webResponse)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

func internalServerError(wr http.ResponseWriter, req *http.Request, err interface{}) {
	wr.Header().Set("Content-Type", "application/json")

	webResponse := response.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal server Error",
		Data:   err,
	}
	wr.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(wr)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}
