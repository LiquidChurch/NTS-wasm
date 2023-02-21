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

func AddAttendanceByGroupId(
	ctx app.Context,
	personId int,
	groupId int,
) *AttendanceDetail {
	occurDate := time.Now().Format("2006-01-02")

	url := rockURL + "Attendances/AddAttendance?groupId=" + strconv.Itoa(groupId) + "&locationId=14&scheduleId=" + scheduleId  + "&occurrenceDate=" + occurDate + "&personId=" + strconv.Itoa(personId)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("PUT", url, nil)

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

	var respJson = new(AttendanceDetail)
	err = json.Unmarshal(body, &respJson)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	
	return respJson
}