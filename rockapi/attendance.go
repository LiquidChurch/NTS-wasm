package rockapi

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type AttendanceDetail struct {
	Id				int		`json:"Id"`
	OccurrenceId	int		`json:"OccurrenceId"`
	PersonAliasId	int		`json:"PersonAliasId"`
	DidAttend		bool	`json:"DidAttend"`
}

func LoadAttendanceByPersonAliasId(
	ctx app.Context,
	personAliasId int,
	occurrenceId int,
) *AttendanceDetail {
	url := rockURL + "Attendances?$filter=OccurrenceId eq " + strconv.Itoa(occurrenceId) + " and PersonAliasId eq " + strconv.Itoa(personAliasId)

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

	var respJson = new([]AttendanceDetail)
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

func LoadAttendancesByOccurrenceId(
	ctx app.Context,
	occurrenceId int,
) []*AttendanceDetail {
	url := "https://rock.liquid.church/api/Attendances?$filter=OccurrenceId eq " + strconv.Itoa(occurrenceId)

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

	var respJson = new([]*AttendanceDetail)
	err = json.Unmarshal(body, &respJson)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return *respJson
}