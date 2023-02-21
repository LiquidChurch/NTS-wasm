package rockapi

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/maxence-charriere/go-app/v9/pkg/app"

)

type GroupMemberDetail struct {
	Id					int				 `json:"Id"`
	GroupId				int				 `json:"GroupId"`
	PersonId			int				 `json:"PersonId"`
	GroupRoleId			int				 `json:"GroupRoleId"`
	Group				GroupDetail		 `json:"Group"`
	Person				PersonDetail 	 `json:"Person"`
	Attendance			AttendanceDetail `json:"Attendance"`
}

func LoadGroupMembersByPersonId(
	ctx app.Context,
	personId int,
	personAliasId int,
) []*GroupMemberDetail {
	url := rockURL + "GroupMembers?$filter=PersonId eq " + strconv.Itoa(personId) + " and (GroupRoleId eq " + strconv.Itoa(MemberRoleId) + " or GroupRoleId eq " + strconv.Itoa(LeaderRoleId) + ")"

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

	var respJson = new([]*GroupMemberDetail)

	err = json.Unmarshal(body, &respJson)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	for i, value := range *respJson {
		log.Println("load group detail", value.GroupId)
		groupDetail := LoadGroupById(ctx, value.GroupId, personAliasId)
		(*respJson)[i].Group = *groupDetail
	}
	log.Println("groupMember", respJson)
	return *respJson
}


func LoadGroupMembersByGroupId(
	ctx app.Context,
	groupId int,
) []*GroupMemberDetail {
	url := "https://rock.liquid.church/api/GroupMembers?$filter=GroupId eq " + strconv.Itoa(groupId)

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

	var respJson = new([]*GroupMemberDetail)

	err = json.Unmarshal(body, &respJson)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return *respJson
}
