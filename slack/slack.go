package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type request struct {
	Text string `json:"text"`
}

func NoticeForgotPass(text string) {
	log.Print(text)
	err := postSlack(request{Text: "Forgot Password : " + text})
	if err != "" {
		log.Println(err)
	}
}

// TODO herokuでのenvの書き方？を調べる
const (
	domain = "https://hooks.slack.com/services/"
	adminWebHock = "T013G7QJRJ5/B0188DNEB8W/DSSIgi86iOq1VkSoT89Nd9G6"
)


func postSlack(requestBody interface{}) string {
	jsonModel, err := json.Marshal(requestBody)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	req, err := http.NewRequest("POST", domain + adminWebHock, bytes.NewBuffer(jsonModel))
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	return ""
}
