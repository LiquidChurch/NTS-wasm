package rockapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"	
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type GuestDetail struct {
	Id					int				`json:"Id"`
	Guid				string			`json:"Guid"`
	RegistrationId		int				`json:"RegistrationId"`
	PersonAliasId		int				`json:"PersonAliasId"`
	Attributes       	GuestAtt		`json:"AttributeValues"`
	Person				PersonDetail	`json:"Person"`
}

type GuestId struct {
	GuestId			int
	PersonId 		int
	PersonAliasId	int
}

type GuestAtt struct {
	Age						AttributeValueDetail 	`json:"Age"`
	GroupHomeFlag			AttributeValueDetail 	`json:"GroupHomeFlag"`
	GroupHome				AttributeValueDetail 	`json:"GroupHome"`
	GroupHomeStaffFlag		AttributeValueDetail	`json:"GroupHomeStaffFlag"`
	GroupHomeStaffName		AttributeValueDetail	`json:"GroupHomeStaffName"`
	GroupHomeStaffPhone		AttributeValueDetail	`json:"GroupHomeStaffPhone"`
	EmergencyContactName	AttributeValueDetail	`json:"EmergencyContactName"`
	EmergencyContactPhone	AttributeValueDetail	`json:"EmergencyContactPhone"`
	EmergencyContactEmail	AttributeValueDetail	`json:"EmergencyContactEmail"`
	EmergencyContactRelat	AttributeValueDetail	`json:"EmergencyContactRelationship"`
	PreviousAttendance		AttributeValueDetail	`json:"PreviousAttendance"`
	SupervisionLevel		AttributeValueDetail	`json:"SupervisionLevel"`
	BathroomAssistance		AttributeValueDetail	`json:"BathroomAssistance"`
	SpecialNeedDescription	AttributeValueDetail	`json:"SpecialNeedDescription"`
	FoodOption				AttributeValueDetail	`json:"FoodOption"`
	MedicationFlag			AttributeValueDetail	`json:"MedicationFlag"`
	MedicationNameTime		AttributeValueDetail	`json:"MedicationNameTime"`
	MedicationAdminister	AttributeValueDetail	`json:"MedicationAdminister"`
	MedicationAdminPhone	AttributeValueDetail	`json:"MedicationAdministerPhone"`
	ConfirmedMedication		AttributeValueDetail	`json:"ConfirmMedicationAdminister"`
	CommunicationLevel		AttributeValueDetail	`json:"CommunicationLevel"`
	LimoFlag				AttributeValueDetail	`json:"LimoFlag"`
	AdditionalInfo			AttributeValueDetail	`json:"AdditionalInfo"`
	MediaReleaseFlag		AttributeValueDetail	`json:"MediaReleaseFlag"`
	MedicalPager			AttributeValueDetail	`json:"MedicalPager"`
	AppKey					AttributeValueDetail	`json:"AppKey"`
}

func LoadGuestByAppKey(
	ctx app.Context,
	appKey string,
) *GuestDetail {
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

	var respJson = new([]GuestDetail)

	err = json.Unmarshal(body, &respJson)
	if err != nil {
		return nil
	}

	guest := (*respJson)[0]
	guestAlias := LoadPersonAliasById(ctx, guest.PersonAliasId)
	guest.Person = guestAlias.Person

	return &guest
}