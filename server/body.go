package server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/api-gateway/types"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

func mergeBody(req *http.Request, sm *types.MatchedMethod) (string, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	//body could be consumed again
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	bodyStr := string(body)
	pathStr := toJSONStr(sm.InputTypeDescriptor, sm.PathValues)
	queryStr := toJSONStr(sm.InputTypeDescriptor, req.URL.Query())

	if bodyStr == "" {
		bodyStr = "{}"
	}

	jsonStr := strings.TrimSuffix(bodyStr, "}") + pathStr + queryStr + "}"
	replacer := strings.NewReplacer("{,", "{")
	return replacer.Replace(jsonStr), nil
}

func toJSONStr(inputType *descriptor.DescriptorProto, values url.Values) (str string) {
	for k, v := range values {
		field := fieldType(k, inputType)
		if field != nil {
			switch *field.Type.Enum() {
			case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
				descriptor.FieldDescriptorProto_TYPE_FLOAT,
				descriptor.FieldDescriptorProto_TYPE_INT64,
				descriptor.FieldDescriptorProto_TYPE_UINT64,
				descriptor.FieldDescriptorProto_TYPE_INT32,
				descriptor.FieldDescriptorProto_TYPE_FIXED64,
				descriptor.FieldDescriptorProto_TYPE_FIXED32,
				descriptor.FieldDescriptorProto_TYPE_BOOL,
				descriptor.FieldDescriptorProto_TYPE_UINT32,
				descriptor.FieldDescriptorProto_TYPE_SFIXED32,
				descriptor.FieldDescriptorProto_TYPE_SFIXED64,
				descriptor.FieldDescriptorProto_TYPE_SINT32,
				descriptor.FieldDescriptorProto_TYPE_SINT64:
				str = str + ",\"" + k + "\":" + v[0] + ""
			default:
				goto DEFAULT
			}
		}
	DEFAULT:
		str = str + ",\"" + k + "\":\"" + v[0] + "\""
	}

	return str
}

func fieldType(key string, inputType *descriptor.DescriptorProto) *descriptor.FieldDescriptorProto {
	for _, v := range inputType.Field {
		if *v.Name == key {
			return v
		}
	}
	return nil
}
