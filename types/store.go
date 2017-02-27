package types

import (
	"api-gateway/common"
	"strings"
)

type RuleStore map[string]MethodWrapper

type MatchedMethod struct {
	MethodWrapper
	Precision  int
	MergeValue string //json format
}

func (rs RuleStore) Match(key string) *MatchedMethod {
	if v, ok := rs[key]; ok {
		return &MatchedMethod{MethodWrapper: v}
	}
	ps := new(PrecisionSet)
	paths := strings.Split(key, "/")

	for keyInDef, methodWrapper := range rs {
		partsInDef := strings.Split(keyInDef, "/")
		if len(paths) == len(partsInDef) {
			value := ""
			precision := 0
			for i := 0; i < len(paths); i++ {
				if strings.HasPrefix(partsInDef[i], "{") {
					property := strings.TrimSuffix(strings.TrimPrefix(partsInDef[i], "{"), "}")
					value = value + ",\"" + property + "\":\"" + paths[i] + "\""
					precision = +1
				} else if partsInDef[i] == paths[i] {
					precision = +2
				} else {
					goto NEXT_LOOP
				}
			}
			method := MatchedMethod{Precision: precision, MergeValue: value, MethodWrapper: methodWrapper}
			log.Debug(method)
			*ps = append(*ps, &method)
		}
	NEXT_LOOP:
	}
	return ps.Max()
}
