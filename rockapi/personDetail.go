package rockapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PersonDetail struct {
	Id				int			`json:"Id"`
	FirstName		string		`json:"FirstName"`
	NickName		string		`json:"NickName"`
	MiddleName		string		`json:"MiddleName"`
	LastName		string		`json:"LastName"`
	FullName		string		`json:"FullName"`
	Email			string		`json:"Email"`
	PrimaryAliasId	int			`json:"PrimaryAliasId"`
}

func LoadPersonById(
	ctx app.Context,
	personId int,
) *PersonDetail {

	var cached *PersonDetail
	ctx.ObserveState("p"+strconv.Itoa(personId)).Value(&cached)

	if cached != nil {
		return cached
	}

	url := rockURL + "People/" + strconv.Itoa(personId) 
	
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil
	}
	req.Header.Set("Authorization-Token", authToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil
	} else if resp.StatusCode != 200 {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var respJson = new(PersonDetail)

	err = json.Unmarshal(body, &respJson)
	if err != nil {
		return nil
	}

	ctx.SetState("p"+strconv.Itoa(personId), &respJson, app.Persist)
	
	return respJson
}