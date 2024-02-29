package controller

import (
	"apilogin/exception"
	"apilogin/model/response"
	"apilogin/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type ControllerImpl struct {
	service service.ServiceInterface
}

// LogOut implements UserController.
func (*ControllerImpl) LogOut(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	cookie := &http.Cookie{
		Name:    "token",
		MaxAge:  -1,
		Expires: time.Now().Add(-1 * time.Hour),
	}
	http.SetCookie(wr, cookie)
	web := response.WebResponse{
		Code:   200,
		Status: "Logged Out",
		Data:   "Logged Out",
	}
	wr.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(wr)
	err := encode.Encode(web)
	if err != nil {
		panic(err)
	}
}

// GetAllWorkerBasedOnRole implements UserController.
func (c *ControllerImpl) GetAllWorkerBasedOnRole(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	cookie, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			panic(exception.NotFoundErrorF("Not Found any Cookie"))
		}
		panic(err)
	}

	tokenStr := cookie.Value
	tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_key"), nil
	})

	if err != nil {
		panic(err)
	}
	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		sid := int(claims["role_id"].(float64))
		if sid != 1 {
			panic(exception.NewUnauthorizedf("Only Admin can access"))
		}
	}

	get := params.ByName("id")
	id, err := strconv.Atoi(get)
	if err != nil {
		panic(err)
	}
	resp := c.service.AllWorkerByJobService(req.Context(), id)
	web := response.WebResponse{
		Code:   200,
		Status: "Authorized",
		Data:   resp,
	}

	wr.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(wr)
	err = encode.Encode(web)
	if err != nil {
		panic(err)
	}

}

// LoginController implements UserController.
func (c *ControllerImpl) LoginController(wr http.ResponseWriter, req *http.Request, param httprouter.Params) {
	decode := json.NewDecoder(req.Body)
	userLogIn := response.User_req{}
	err := decode.Decode(&userLogIn)
	if err != nil {
		panic(err)
	}
	ctrl, token := c.service.LoginService(req.Context(), userLogIn)
	web := response.WebResponse{
		Code:   200,
		Status: "Authorized",
		Data:   ctrl,
	}
	expirate := time.Now().Add(time.Minute * 5)

	http.SetCookie(wr, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirate,
	})

	wr.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(wr)
	err = encode.Encode(web)
	if err != nil {
		panic(err)
	}
}

// RegisterController implements UserController.
func (c *ControllerImpl) RegisterController(wr http.ResponseWriter, req *http.Request, param httprouter.Params) {
	decode := json.NewDecoder(req.Body)
	userRegis := response.Create_req{}

	err := decode.Decode(&userRegis)
	if err != nil {
		panic(exception.BadRequestF("It can be wrong data type"))
	}
	c.service.RegisterService(req.Context(), userRegis)
	web := response.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   "Account Created",
	}
	wr.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(wr)
	err = encode.Encode(web)
	if err != nil {
		panic(err)
	}
}

func NewController(s service.ServiceInterface) UserController {
	return &ControllerImpl{
		service: s,
	}
}
