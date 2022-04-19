package main

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/volumestream"
	"github.com/michimani/gotwi/tweet/volumestream/types"
)

func samplingTweets(c *gotwi.Client, count int) {
	p := &types.SampleStreamInput{}
	s, err := volumestream.SampleStream(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	cnt := 0
	for s.Receive() {
		cnt++
		t, err := s.Read()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(gotwi.StringValue(t.Data.ID), gotwi.StringValue(t.Data.Text))
		}

		if cnt > count {
			s.Stop()
			break
		}
	}
}
