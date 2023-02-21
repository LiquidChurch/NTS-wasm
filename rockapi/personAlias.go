package rockapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PersonAliasDetail struct {
	Id				int				`json:"Id"`
	PersonId		int				`json:"PersonId"`
	Person 			PersonDetail 	`json:"Person"`
}

func LoadPersonAliasById(
	ctx app.Context,
	aliasId int,
) *PersonAliasDetail {
	url := rockURL + "PersonAlias/" + strconv.Itoa(aliasId) 

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

	var respJson = new(PersonAliasDetail)

	err = json.Unmarshal(body, &respJson)
	if err != nil {
		return nil
	}

	return respJson
}