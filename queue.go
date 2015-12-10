package bambooapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type QueueResponse struct {
	Message        string `json:"message"`
	StatusCode     int    `json:"status-code"`
	Plankey        string `json:"planKey"`
	Buildnumber    int    `json:"buildNumber"`
	Buildresultkey string `json:"buildResultKey"`
	Triggerreason  string `json:"triggerReason"`
	Link           struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"link"`
}

func Queue(server string, user string, pass string, buildkey string) (response QueueResponse) {
	url := "https://" + user + ":" + pass + "@" + server
	url += "/builds/rest/api/latest/queue/" + buildkey
	url += ".json?executeAllStages=true&os_authType=basic"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(""))
	if err != nil {
		log.Print(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
	}
	json.Unmarshal(body, &response)
	return response
}
