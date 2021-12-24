package gotwi_test

import (
	"testing"
	"time"

	"github.com/michimani/gotwi"
	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	cases := []struct {
		name string
		s    string
	}{
		{
			name: "normal",
			s:    "test string",
		},
		{
			name: "normal: empty string",
			s:    "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			sp := gotwi.String(c.s)
			assert.Equal(tt, c.s, *sp)
		})
	}
}

func Test_StringValue(t *testing.T) {
	cases := []struct {
		name string
		s    string
	}{
		{
			name: "normal",
			s:    "test string",
		},
		{
			name: "normal: empty string",
			s:    "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			sv := gotwi.StringValue(&c.s)
			assert.Equal(tt, c.s, sv)
		})
	}

	// nil case
	t.Run("nil case", func(tt *testing.T) {
		sv := gotwi.StringValue(nil)
		assert.Empty(tt, sv)
	})
}

func Test_Bool(t *testing.T) {
	cases := []struct {
		name string
		b    bool
	}{
		{
			name: "normal: true",
			b:    true,
		},
		{
			name: "normal: false",
			b:    false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			bp := gotwi.Bool(c.b)
			assert.Equal(tt, c.b, *bp)
		})
	}
}

func Test_BoolValue(t *testing.T) {
	cases := []struct {
		name string
		b    bool
	}{
		{
			name: "normal: true",
			b:    true,
		},
		{
			name: "normal: false",
			b:    false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			bv := gotwi.BoolValue(&c.b)
			assert.Equal(tt, c.b, bv)
		})
	}

	// nil case
	t.Run("nil case", func(tt *testing.T) {
		bv := gotwi.BoolValue(nil)
		assert.False(tt, bv)
	})
}

func Test_Int(t *testing.T) {
	cases := []struct {
		name string
		i    int
	}{
		{
			name: "normal: zero",
			i:    0,
		},
		{
			name: "normal: positive",
			i:    1,
		},
		{
			name: "normal: negative",
			i:    -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ip := gotwi.Int(c.i)
			assert.Equal(tt, c.i, *ip)
		})
	}
}

func Test_IntValue(t *testing.T) {
	cases := []struct {
		name string
		i    int
	}{
		{
			name: "normal: zero",
			i:    0,
		},
		{
			name: "normal: positive",
			i:    1,
		},
		{
			name: "normal: negative",
			i:    -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			iv := gotwi.IntValue(&c.i)
			assert.Equal(tt, c.i, iv)
		})
	}

	// nil case
	t.Run("nil case", func(tt *testing.T) {
		iv := gotwi.IntValue(nil)
		assert.Equal(tt, iv, 0)
	})
}

func Test_Float64(t *testing.T) {
	cases := []struct {
		name string
		i    float64
	}{
		{
			name: "normal: zero",
			i:    0,
		},
		{
			name: "normal: positive",
			i:    1,
		},
		{
			name: "normal: negative",
			i:    -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ip := gotwi.Float64(c.i)
			assert.Equal(tt, c.i, *ip)
		})
	}
}

func Test_Float64Value(t *testing.T) {
	cases := []struct {
		name string
		i    float64
	}{
		{
			name: "normal: zero",
			i:    0,
		},
		{
			name: "normal: positive",
			i:    1,
		},
		{
			name: "normal: negative",
			i:    -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			iv := gotwi.Float64Value(&c.i)
			assert.Equal(tt, c.i, iv)
		})
	}

	// nil case
	t.Run("nil case", func(tt *testing.T) {
		iv := gotwi.Float64Value(nil)
		assert.Equal(tt, iv, float64(0))
	})
}

func Test_Time(t *testing.T) {
	cases := []struct {
		name string
		t    time.Time
	}{
		{
			name: "normal",
			t:    time.Now(),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tp := gotwi.Time(c.t)
			assert.Equal(tt, c.t, *tp)
		})
	}
}

func Test_TimeValue(t *testing.T) {
	cases := []struct {
		name string
		t    time.Time
	}{
		{
			name: "normal",
			t:    time.Now(),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tv := gotwi.TimeValue(&c.t)
			assert.Equal(tt, c.t, tv)
		})
	}

	// nil case
	t.Run("nil case", func(tt *testing.T) {
		tv := gotwi.TimeValue(nil)
		assert.Equal(tt, tv, time.Time{})
	})
}
