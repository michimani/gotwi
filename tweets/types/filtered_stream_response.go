package types

import "github.com/michimani/gotwi/resources"

type FilteredStreamRulesGetResponse struct {
	Data   resources.FilterdStreamRule `json:"data"`
	Meta   resources.FilterdStreamRulesGetMeta
	Errors []resources.PartialError `json:"errors"`
}

func (r *FilteredStreamRulesGetResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
