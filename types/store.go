package types

import (
	"api-gateway/common"
	"strings"
)

type RuleStore struct {
	Store map[string]MethodWrapper
}

func CreateRuleStore() *RuleStore {
	store := make(map[string]MethodWrapper)
	return &RuleStore{Store: store}
}

func (rs *RuleStore) Match(key string) *MatchedMethod {
	if v, ok := rs.Store[key]; ok {
		return &MatchedMethod{MethodWrapper: v}
	}
	precisionSet := new(PrecisionSet)
	paths := strings.Split(key, "/")

	for keyInDef, methodWrapper := range rs.Store {
		partsInDef := strings.Split(keyInDef, "/")
		if len(paths) == len(partsInDef) {
			value := ""
			precision := 0
			for i := 0; i < len(paths); i++ {
				if strings.HasPrefix(partsInDef[i], "{") {
					value = value + "," + partsInDef[i] + "=" + paths[i]
					precision = precision + 1
				} else if partsInDef[i] == paths[i] {
					precision = precision + 2
				} else {
					goto LABEL //exit sub-loop
				}
			}
			method := MatchedMethod{precision: precision, mergeValue: value, MethodWrapper: methodWrapper}
			log.Debug(method)
			*precisionSet = append(*precisionSet, &method)
		}
	LABEL:
	}

	return precisionSet.Max()
}

type MatchedMethod struct {
	precision int
	MethodWrapper
	mergeValue string //json format
}
