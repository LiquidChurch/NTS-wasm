package rockapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type RegistrantDetail struct {
	Id					int				`json:"Id"`
	Guid				string			`json:"Guid"`
	RegistrationId		int				`json:"RegistrationId"`
	PersonAliasId		int				`json:"PersonAliasId"`
	Attributes       	RegistrantAtt	`json:"AttributeValues"`
	Person				PersonDetail	`json:"Person"`
}

type RegistrantAtt struct {
	BGCRequired 	AttributeValueDetail `json:"BGCRequired"`
	PriorServed 	AttributeValueDetail `json:"PriorServed"`
	SNExperience 	AttributeValueDetail `json:"SNExperience"`
	BGCException 	AttributeValueDetail `json:"BGCException"`
}

func LoadRegistrantByAppKey(
	ctx app.Context,
	appKey string,
) *RegistrantDetail {
	url := rockURL + "RegistrationRegistrants/GetByAttributeValue?attributeKey=AppKey&value=" + appKey + "&loadAttributes=simple"
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

	var respJson = new([]RegistrantDetail)

	err = json.Unmarshal(body, &respJson)
	if err != nil {
		return nil
	}

	registrant := (*respJson)[0]
	personAlias := LoadPersonAliasById(ctx, registrant.PersonAliasId)
	registrant.Person = personAlias.Person

	return &registrant
}