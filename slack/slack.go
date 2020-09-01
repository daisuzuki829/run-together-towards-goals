package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type request struct {
	Text string `json:"text"`
}

// NoticeForgotPass ...
func NoticeForgotPass(text string) {
	err := postSlack(request{Text: "Forgot Password : " + text})
	if err != "" {
		log.Println(err)
	}
}

func postSlack(requestBody interface{}) string {
	jsonModel, err := json.Marshal(requestBody)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	var adminWebHock string
	if os.Getenv("SLACK_ADMIN_WEBHOCK") != "" {
		adminWebHock = os.Getenv("SLACK_ADMIN_WEBHOCK")
	} else {
		return "Error : Unset slack Admin WebHock"
	}
	req, err := http.NewRequest("POST", adminWebHock, bytes.NewBuffer(jsonModel))
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
