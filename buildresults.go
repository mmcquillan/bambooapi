package bambooapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type BuildResultResponse struct {
	Results struct {
		Result []struct {
			Link struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"link"`
			Plan struct {
				Shortname string `json:"shortName"`
				Shortkey  string `json:"shortKey"`
				Type      string `json:"type"`
				Enabled   bool   `json:"enabled"`
				Key       string `json:"key"`
				Name      string `json:"name"`
				Plankey   struct {
					Key string `json:"key"`
				} `json:"planKey"`
			} `json:"plan"`
			Buildresultkey string `json:"buildResultKey"`
			Lifecyclestate string `json:"lifeCycleState"`
			ID             int    `json:"id"`
			Key            string `json:"key"`
			Planresultkey  struct {
				Key       string `json:"key"`
				Entitykey struct {
					Key string `json:"key"`
				} `json:"entityKey"`
				Resultnumber int `json:"resultNumber"`
			} `json:"planResultKey"`
			State       string `json:"state"`
			Buildstate  string `json:"buildState"`
			Number      int    `json:"number"`
			Buildnumber int    `json:"buildNumber"`
		} `json:"result"`
	} `json:"results"`
}

func BuildResults(server string, user string, pass string) (response BuildResultResponse) {
	url := "https://" + user + ":" + pass + "@" + server
	url += "/builds/rest/api/latest/result.json?os_authType=basic"
	req, err := http.NewRequest("GET", url, bytes.NewBufferString(""))
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
