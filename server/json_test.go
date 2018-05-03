package server

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	jsonStr := `{"name":"abc","age":6}`
	va, err := json.Marshal(jsonStr)
	if err != nil {
		t.Error(err)
	}
	//t.Log(va)
	fmt.Println(string(va))
	t.Error("")
}
