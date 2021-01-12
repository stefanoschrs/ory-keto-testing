package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/stefanoschrs/ory-keto-testing/config"
	"github.com/stefanoschrs/ory-keto-testing/utils"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := utils.GetClient()

	var res *resty.Response
	var err error

	type Policy struct {
		Id        string
		Resources []string
		Actions   []string
		Subjects  []string
	}
	policyMatrix := []Policy{
		{
			config.ProductAProject1PolicyManage,
			[]string{config.ProductAProject1Resource},
			[]string{config.ProductAActionRead, config.ProductAActionEdit, config.ProductAActionDelete},
			[]string{config.ProductASubjectAdmin, config.ProductAProject1SubjectManage},
		},
		{
			config.ProductAProject1PolicyEdit,
			[]string{config.ProductAProject1Resource},
			[]string{config.ProductAActionRead, config.ProductAActionEdit},
			[]string{config.ProductAProject1SubjectEdit},
		},
		{
			config.ProductAProject1PolicyRead,
			[]string{config.ProductAProject1Resource},
			[]string{config.ProductAActionRead},
			[]string{config.ProductAProject1SubjectRead},
		},

		{
			config.ProductAProject2PolicyManage,
			[]string{config.ProductAProject2Resource},
			[]string{config.ProductAActionRead, config.ProductAActionEdit, config.ProductAActionDelete},
			[]string{config.ProductASubjectAdmin, config.ProductAProject2SubjectManage},
		},
		{
			config.ProductAProject2PolicyEdit,
			[]string{config.ProductAProject2Resource},
			[]string{config.ProductAActionRead, config.ProductAActionEdit},
			[]string{config.ProductAProject2SubjectEdit},
		},
		{
			config.ProductAProject2PolicyRead,
			[]string{config.ProductAProject2Resource},
			[]string{config.ProductAActionRead},
			[]string{config.ProductAProject2SubjectRead},
		},
	}

	roles := []string{
		config.ProductASubjectAdmin,

		config.ProductAProject1SubjectManage,
		config.ProductAProject1SubjectEdit,
		config.ProductAProject1SubjectRead,

		config.ProductAProject2SubjectManage,
		config.ProductAProject2SubjectEdit,
		config.ProductAProject2SubjectRead,
	}

	usersMatrix := []struct{
		Id string
		Role string
	}{
		{config.UserAdmin, config.ProductASubjectAdmin},

		{config.User1, config.ProductAProject1SubjectManage},
		{config.User2, config.ProductAProject1SubjectEdit},
		{config.User3, config.ProductAProject1SubjectRead},
		{config.User5, config.ProductAProject1SubjectRead},

		{config.User1, config.ProductAProject2SubjectManage},
		{config.User4, config.ProductAProject2SubjectEdit},
		{config.User5, config.ProductAProject2SubjectRead},
	}

	// Add policies
	for _, policy := range policyMatrix {
		fmt.Println(utils.GetHeaderStr("Upsert an ORY Access Control Policy - " + policy.Id))
		res, err = client.R().
			SetBody(map[string]interface{}{
				"id":          policy.Id,
				"description": "a description string",
				"resources":   policy.Resources,
				"actions":     policy.Actions,
				"subjects":    policy.Subjects,
				"effect":      config.EffectAllow,
			}).
			Put("/engines/acp/ory/exact/policies")
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode() != http.StatusOK {
			log.Fatal(res.Status(), string(res.Body()))
		}
		fmt.Println(utils.GetIndentedJson(res.Body()))
	}

	// Add roles
	for _, subjectId := range roles {
		fmt.Println(utils.GetHeaderStr("Upsert an ORY Access Control Policy Role - " + subjectId))
		res, err = client.R().
			SetBody(map[string]interface{}{
				"id":          subjectId,
				"description": "a description string",
				"members":     []string{},
			}).
			Put("/engines/acp/ory/exact/roles")
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode() != http.StatusOK {
			log.Fatal(res.Status(), string(res.Body()))
		}
		fmt.Println(utils.GetIndentedJson(res.Body()))
	}

	// Add users to roles
	for _, user := range usersMatrix {
		fmt.Println(utils.GetHeaderStr("Add a Member to an ORY Access Control Policy Role - " + user.Role))
		res, err = client.R().
			SetBody(map[string]interface{}{
				"members": []string{user.Id},
			}).
			Put("/engines/acp/ory/exact/roles/" + user.Role + "/members")
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode() != http.StatusOK {
			log.Fatal(res.Status(), string(res.Body()))
		}
		fmt.Println(utils.GetIndentedJson(res.Body()))
	}
}
