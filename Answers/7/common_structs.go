package dao

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Alias    string `json:"alias"`
}

type Session struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Expiry   string `json:"expiry"`
}

func AuthenticateUser(cred UserCredentials) Session {
	session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("user")
	query := clctn.Find(bson.M{"username": cred.Username, "password": cred.Password})
	err := query.One(&response)

	uuidStr := uuid.Must(uuid.NewV4()).String()
	timein := time.Now()

	timeout := timein.Local().Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(5) +
		time.Second*time.Duration(0))

	fmt.Println(timein, timeout)
	sessionStruct := Session{cred.Username, uuidStr, timeout.String()}
	if err != nil {
		return Session{}
	}

	sessionClctn := session.DB("simplesurveys").C("session")
	sessionClctn.Insert(sessionStruct)
	return sessionStruct
}

func GetSessionDetails(token string) UserCredentials {
	session := MgoSession.Clone()
	defer session.Close()

	var response Session
	sessionClctn := session.DB("simplesurveys").C("session")
	query := sessionClctn.Find(bson.M{"token": token})
	err := query.One(&response)
	if err != nil {
		return UserCredentials{}
	}

	var cred UserCredentials
	clctn := session.DB("simplesurveys").C("user")
	query = clctn.Find(bson.M{"username": response.Username})
	err = query.One(&cred)
	return cred
}

func SignupUser(cred UserCredentials) string {
	session := MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("simplesurveys").C("user")

	var signup UserCredentials
	query := clctn.Find(bson.M{"username": cred.Username})
	err := query.One(&signup)
	if signup.Username != "" {
		return "User Already Present"
	}

	err = clctn.Insert(&cred)
	if err != nil {
		panic(err)
	}
	return "User Registered"

}
