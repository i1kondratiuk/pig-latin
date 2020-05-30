package piglatin

import (
	"reflect"
	"testing"
)

var testData = []struct {
	in  string
	out string
}{
	{
		in:  " ",
		out: " ",
	},
	{
		in:  "pig",
		out: "igpay",
	},
	{
		in:  "latin banana",
		out: "atinlay ananabay",
	},
	{
		in:  "smile",
		out: "ilesmay",
	},
	{
		in:  "string stupid glove",
		out: "ingstray upidstay oveglay",
	},
	{
		in:  "eat",
		out: "eatyay",
	},
	{
		in:  "omelet are egg explain",
		out: "omeletyay areyay eggyay explainyay",
	},
}

func TestTranslate(t *testing.T) {
	for _, test := range testData {
		AssertEqual(t, test.out, Translate(test.in))
	}
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}

	t.Errorf("Expected %v (type %v), received %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}
