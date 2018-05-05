package dao

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Question struct {
	QuestionString string   `json:"questionString" bson:"questionString"`
	Options        []string `json:"options" bson:"options"`
}

type Answer struct {
	Question Question `json:"question" bson:"question"`
	Answer   string   `json:"answer" bson:"answer"`
}

type Survey struct {
	SurveyName  string     `json:"surveyName" bson:"surveyName"`
	Heading     string     `json:"heading" bson:"heading"`
	Description string     `json:"description" bson:"description"`
	Questions   []Question `json:"questions" bson:"questions"`
	Status      bool       `json:"status" bson:"status"`
	Expiry      string     `json:"expiry" bson:"expiry"`
}

type SurveyResponse struct {
	UserName string   `json:"userName" bson:"userName"`
	Survey   Survey   `json:"survey" bson:"survey"`
	Answers  []Answer `json:"answers" bson:"answers"`
}

func GetActiveSurveys() interface{} {
	session := MgoSession.Clone()
	defer session.Close()

	var response []interface{}
	clctn := session.DB("simplesurveys").C("survey")
	query := clctn.Find(bson.M{"status": true})
	err := query.All(&response)

	if err != nil {
		return nil
	} else {
		return response
	}
}

func GetSurveysForUser(userName string) interface{} {
	session := MgoSession.Clone()
	defer session.Close()

	var response []interface{}
	clctn := session.DB("simplesurveys").C("survey_response")
	query := clctn.Find(bson.M{"userName": userName})
	err := query.All(&response)

	/*
		var notExpiredResponse []interface{}

		for i := 0; i < len(response); i++ {
			layout := "2006-01-02T15:04:05.000Z"
			str := response[i]["expiry"]
			t, err := time.Parse(layout, str)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(t)
			notExpiredResponse = append(notExpiredResponse)
		}
	*/

	if err != nil {
		return nil
	} else {
		return response
	}
}

func GetSurveyByName(surveyName string) interface{} {
	fmt.Println("GetSurveyByName:" + surveyName)
	session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("survey")
	query := clctn.Find(bson.M{"surveyname": surveyName})
	err := query.One(&response)

	if err != nil {
		return nil
	} else {
		return response
	}
}

func InsertUserResponse(userResponse SurveyResponse) {
	session := MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("simplesurveys").C("survey_response")
	clctn.Insert(userResponse)
}
