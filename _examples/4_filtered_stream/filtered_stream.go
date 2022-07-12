package main

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/filteredstream"
	"github.com/michimani/gotwi/tweet/filteredstream/types"
)

// createSearchStreamRules lists search stream rules.
func listSearchStreamRules() {
	c, err := newGotwiClientWithTimeout(30)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.ListRulesInput{}
	res, err := filteredstream.ListRules(context.Background(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", gotwi.StringValue(r.ID), gotwi.StringValue(r.Value), gotwi.StringValue(r.Tag))
	}
}

func deleteSearchStreamRules(ruleID string) {
	c, err := newGotwiClientWithTimeout(30)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.DeleteRulesInput{
		Delete: &types.DeletingRules{
			IDs: []string{
				ruleID,
			},
		},
	}

	res, err := filteredstream.DeleteRules(context.TODO(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", gotwi.StringValue(r.ID), gotwi.StringValue(r.Value), gotwi.StringValue(r.Tag))
	}
}

// createSearchStreamRules creates a search stream rule.
func createSearchStreamRules(keyword string) {
	c, err := newGotwiClientWithTimeout(30)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.CreateRulesInput{
		Add: []types.AddingRule{
			{Value: gotwi.String(keyword), Tag: gotwi.String(keyword)},
		},
	}

	res, err := filteredstream.CreateRules(context.TODO(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", gotwi.StringValue(r.ID), gotwi.StringValue(r.Value), gotwi.StringValue(r.Tag))
	}
}

// execSearchStream call GET /2/tweets/search/stream API
// and outputs up to 10 results.
func execSearchStream() {
	c, err := newGotwiClientWithTimeout(120)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.SearchStreamInput{}
	s, err := filteredstream.SearchStream(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	cnt := 0
	for s.Receive() {
		t, err := s.Read()
		if err != nil {
			fmt.Println(err)
		} else {
			if t != nil {
				cnt++
				fmt.Println(gotwi.StringValue(t.Data.ID), gotwi.StringValue(t.Data.Text))
			}
		}

		if cnt > 10 {
			s.Stop()
			break
		}
	}
}
