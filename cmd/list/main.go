package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stefanoschrs/ory-keto-testing/utils"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := utils.GetClient()

	var res *resty.Response
	var err error

	fmt.Println(utils.GetHeaderStr("List ORY Access Control Policies"))
	res, err = client.R().
		Get("/engines/acp/ory/exact/policies")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode() != http.StatusOK {
		log.Fatal(res.Status(), string(res.Body()))
	}
	fmt.Println(utils.GetIndentedJson(res.Body()))

	fmt.Println(utils.GetHeaderStr("List ORY Access Control Policy Roles"))
	res, err = client.R().
		Get("/engines/acp/ory/exact/roles")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode() != http.StatusOK {
		log.Fatal(res.Status(), string(res.Body()))
	}
	fmt.Println(utils.GetIndentedJson(res.Body()))
}
