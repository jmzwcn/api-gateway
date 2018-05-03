package server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/api-gateway/types"
)

func mergeBody(req *http.Request, sm *types.MatchedMethod) (string, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	//body could be consumed again
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	bodyStr := string(body)
	pathStr := toJSONStr(sm.Method.InputType, sm.PathValues)
	queryStr := toJSONStr(sm.Method.InputType, req.URL.Query())

	if bodyStr == "" {
		bodyStr = "{}"
	}

	jsonStr := strings.TrimSuffix(bodyStr, "}") + pathStr + queryStr + "}"
	replacer := strings.NewReplacer("{,", "{")
	return replacer.Replace(jsonStr), nil
}

func toJSONStr(inputType *string, values url.Values) (str string) {
	for k, v := range values {
		//field := reflect.ValueOf(msg).Elem().FieldByName(strings.Title(k))
		//if field.IsValid() {
		// switch field.Type().Name() {
		// case "int", "int8", "int16", "int32", "int64",
		// 	"uint", "uint8", "uint16", "uint32", "uint64",
		// 	"float32", "float64", "bool":
		// 	str = str + ",\"" + k + "\":" + v[0] + ""
		// default:
		// 	str = str + ",\"" + k + "\":\"" + v[0] + "\""
		// }
		//}
		str = str + ",\"" + k + "\":\"" + v[0] + "\""
	}

	return str
}
