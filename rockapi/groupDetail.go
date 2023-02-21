package rockapi

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type GroupDetail struct {
	Id				int 				`json:"Id"`
	ParentGroupId	int					`json:"ParentGroupId"`
	GroupTypeId		int					`json:"GroupTypeId"`
    Name			string				`json:"Name"`
	Description		string				`json:"Description"`
	Attributes      GroupAtt			`json:"AttributeValues"`
	Occurrence		OccurrenceDetail	`json:"Occurrence"`
	Attendance		AttendanceDetail	`json:"Attendance"`
}

type GroupAtt struct {
	WhattoExpect	AttributeValueDetail	`json:"WhattoExpect"`
	Map				AttributeValueDetail	`json:"Map"`
	CheckinDate		AttributeValueDetail	`json:"CheckinDate"`
}

func LoadGroupById(
	ctx app.Context,
	groupId int,
	personAliasId int,
) *GroupDetail {
	log.Println("LoadGroupById", groupId)
	url := rockURL + "Groups/" + strconv.Itoa(groupId) + "?loadAttributes=simple"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	req.Header.Set("Authorization-Token", authToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		log.Println(resp.StatusCode)
		return nil
	} else if resp.StatusCode != 200 {
		log.Println(resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	var respJson = new(GroupDetail)

	err = json.Unmarshal(body, &respJson)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	groupDetail := (*respJson)

	occurrence := LoadOccurrenceByGroupId(ctx, respJson.Id)

	if occurrence != nil {
		groupDetail.Occurrence = *occurrence
		if personAliasId != 0 {
			attendance := LoadAttendanceByPersonAliasId(ctx, personAliasId, occurrence.Id)
			log.Println(attendance)
			if attendance != nil {
				groupDetail.Attendance = *attendance
			}
		}
	}
	return &groupDetail
}