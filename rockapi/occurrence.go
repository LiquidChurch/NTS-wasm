package rockapi

import (
	"log"
	"time"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type OccurrenceDetail struct {
	Id				int		`json:"Id"`
	GroupId 		int 	`json:"GroupId"`
	LocationId		int		`json:"LocaitonId"`
	ScheduleId		int		`json:"ScheduleId"`
	OccurrenceDate	string	`json:"OccurrenceDate"`
}

func LoadOccurrenceByGroupId(
	ctx app.Context,
	groupId int,
) *OccurrenceDetail {
	occurDate := time.Now().Format("2006-01-02")

	url := rockURL + "AttendanceOccurrences?$filter=GroupId eq " +  strconv.Itoa(groupId) + " and ScheduleId eq " + scheduleId  + " and OccurrenceDate eq datetime'" + occurDate + "'"

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

	var respJson = new([]OccurrenceDetail)
	err = json.Unmarshal(body, &respJson)
	
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	if len(*respJson) == 0 {
		return nil
	} else {
		return &(*respJson)[0]
	}
}

