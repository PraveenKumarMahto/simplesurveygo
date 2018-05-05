package servicehandlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simplesurveygo/dao"
)

type SignupHandler struct {
}

func (p SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p SignupHandler) Get(r *http.Request) SrvcRes {
	return Response200OK("All Good")
}

func (p SignupHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p SignupHandler) Post(r *http.Request) SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var cred dao.UserCredentials
	err = json.Unmarshal(body, &cred)
	fmt.Println(cred)
	res := dao.SignupUser(cred)
	if res == "User Already Present" {
		return SimpleBadRequest(res)
	} else {
		return Response200OK(res)
	}
}
