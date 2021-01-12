package main

import (
	"encoding/json"
	"log"

	"github.com/stefanoschrs/ory-keto-testing/utils"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := utils.GetClient()

	var res *resty.Response
	var err error

	// Clean policies
	res, err = client.R().
		Get("/engines/acp/ory/exact/policies")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode() >= 400 {
		log.Fatal(res.Status(), string(res.Body()))
	}

	var policies []struct{
		Id string `json:"id"`
	}
	err = json.Unmarshal(res.Body(), &policies)
	if err != nil {
		log.Fatal(err)
	}

	for _, policy := range policies {
		res, err = client.R().
			Delete("/engines/acp/ory/exact/policies/" + policy.Id)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode() >= 400 {
			log.Fatal(res.Status(), string(res.Body()))
		}
	}

	// Clean roles
	res, err = client.R().
		Get("/engines/acp/ory/exact/roles")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode() >= 400 {
		log.Fatal(res.Status(), string(res.Body()))
	}

	var roles []struct{
		Id string `json:"id"`
	}
	err = json.Unmarshal(res.Body(), &roles)
	if err != nil {
		log.Fatal(err)
	}

	for _, role := range roles {
		res, err = client.R().
			Delete("/engines/acp/ory/exact/roles/" + role.Id)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode() >= 400 {
			log.Fatal(res.Status(), string(res.Body()))
		}
	}
}
