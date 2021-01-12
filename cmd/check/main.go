package main

import (
	"fmt"
	"log"

	"github.com/stefanoschrs/ory-keto-testing/config"
	"github.com/stefanoschrs/ory-keto-testing/utils"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := utils.GetClient()

	var res *resty.Response
	var err error

	type TestCase struct{
		Resource, Action, User string
	}

	var testMatrix []TestCase
	for _, r := range []string{config.ProductAProject1Resource, config.ProductAProject2Resource} {
		for _, a := range []string{config.ProductAActionRead, config.ProductAActionEdit, config.ProductAActionDelete} {
			for _, u := range []string{config.UserAdmin, config.User1, config.User2, config.User3, config.User4, config.User5} {
				testMatrix = append(testMatrix, TestCase{r, a, u})
			}
		}
	}

	for _, testCase := range testMatrix {
		fmt.Println(utils.GetHeaderStr(fmt.Sprintf("Check If a Request is Allowed - %v", testCase)))
		res, err = client.R().
			SetBody(map[string]interface{}{
				"resource": testCase.Resource,
				"action":   testCase.Action,
				"subject":  testCase.User,
			}).
			Post("/engines/acp/ory/exact/allowed")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Status())
	}
}
