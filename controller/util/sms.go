package util

import (
	"encoding/json"
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS() {
	from := "+15185033052"
	to := "+27617825205"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: "AC655a14622c3e3d5809967e4c6e8883d0",
		Password: "fbe2f523e5aff20a4e0b1ec5a5b319c5",
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello there voila la vie est mikuna ni mpongo")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
