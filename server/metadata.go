package server

// import (
// 	"context"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"

// 	"google.golang.org/grpc/metadata"
// 	"nevis.io/fx/common/constants"
// 	"nevis.io/fx/common/log"
// 	"nevis.io/fx/core/types"
// )

// func attachMD(ctx context.Context, req *http.Request, options map[string]interface{}) context.Context {
// 	m := make(map[string]string)
// 	//attach options
// 	if v, err := json.Marshal(options); err == nil {
// 		m[constants.RPC_METHOD_OPTIONS] = string(v)
// 	} else {
// 		log.Error(err)
// 	}

// 	//attach http request
// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	req.ParseForm()
// 	newRequest := types.HttpRequest{
// 		Header: req.Header,
// 		Method: req.Method,
// 		Host:   req.Host,
// 		URL:    req.URL,
// 		Form:   req.Form,
// 		Body:   string(body),
// 	}

// 	if v, err := json.Marshal(newRequest); err == nil {
// 		m[constants.HTTP_REQUEST] = string(v)
// 	} else {
// 		log.Errorf("attachMD() %s", err)
// 	}
// 	md := metadata.New(m)
// 	md[constants.RPC_FROM] = []string{"ui", "api-gateway"}

// 	return metadata.NewOutgoingContext(ctx, md)
// }
