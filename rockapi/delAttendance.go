package rockapi

import (
	"log"
	"bytes"
	"net/http"
	"strconv"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func DelAttendanceById(
	ctx app.Context,
	attendanceId int,
) *AttendanceDetail {
	url := rockURL + "Attendances/" + strconv.Itoa(attendanceId)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	
	var jsonData = []byte(`{"DidAttend": false}`)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	req.Header.Set("Authorization-Token", authToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		log.Println(resp.StatusCode)
		return nil
	} else if resp.StatusCode != 204 {
		log.Println(resp.StatusCode)
		return nil
	}

	return nil
}