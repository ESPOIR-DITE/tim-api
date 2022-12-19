package api

import (
	"gopkg.in/resty.v1"
	"tim-api/config"
)

// const BASE_URL string = "http://localhost:9000/ostm/"
// const BASE_URL string = "http://172.17.0.2:9000/ostm/"
const AdminId = "R-f5bf0580-a243-47dc-8c3b-e3b07530cf83"
const SuperAdminId = "R-32f18fc5-f4fb-45e6-a97d-a609b93c6cee"
const AgentId = "R-b39356d4-5d0a-47b2-931f-2a3b7b51f217"
const TestBaseURL = "http://localhost:8081/"

const BASE_URL string = "http://159.69.222.82:9000/ostm/"

func Rest() *resty.Request {
	return resty.R().SetAuthToken("").
		SetHeader("Access-Control-Allow-Origin", "http://localhost:8082").
		SetHeader("Accept", "application/json").
		SetHeader("email", "email").
		SetHeader("site", "site").
		SetHeader("Access-Control-Allow-Origin", "*").
		SetHeader("Content-Type", "application/json")
}

var JSON = config.ConfigWithCustomTimeFormat
