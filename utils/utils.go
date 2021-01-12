package utils

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/go-resty/resty/v2"
)

func GetClient() *resty.Client {
	client := resty.New()
	client.SetHostURL("http://localhost:4466")

	return client
}

func GetHeaderStr(s string) string {
	line := "//" + strings.Repeat("*", len(s)+2) + "//"

	return line + "\n" +
		"// " + s + " //\n" +
		line
}

func GetIndentedJson(b []byte) string {
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, b, "", "\t")

	return prettyJSON.String()
}
