package types

import "github.com/michimani/gotwi/resources"

type ListRulesOutput struct {
	Data   []resources.FilterdStreamRule `json:"data"`
	Meta   resources.FilterdStreamRulesGetMeta
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListRulesOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
