package deploy_test

import (
	"github.com/aws-cloudformation/rain/internal/cmd/deploy"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestListToMap(t *testing.T) {
	testCases := []struct {
		input    []string
		expected map[string]string
	}{
		{[]string{"Foo=bar"}, map[string]string{"Foo": "bar"}},
		{[]string{"Foo=bar", "Baz=quux"}, map[string]string{"Foo": "bar", "Baz": "quux"}},
		{[]string{"Foo=bar", "baz"}, map[string]string{"Foo": "bar,baz"}},
		{[]string{"Foo=bar", "Baz=quux", "mooz"}, map[string]string{"Foo": "bar", "Baz": "quux,mooz"}},
		{[]string{"Foo=bar", "Baz=quux", "mooz", "Xyzzy=garply"}, map[string]string{"Foo": "bar", "Baz": "quux,mooz", "Xyzzy": "garply"}},
		{[]string{"Foo=bar", "Baz=quux", "Mooz=xyzzy", "garply"}, map[string]string{"Foo": "bar", "Baz": "quux", "Mooz": "xyzzy,garply"}},
	}

	for _, testCase := range testCases {
		actual := deploy.ListToMap("test", testCase.input)

		if d := cmp.Diff(testCase.expected, actual); d != "" {
			t.Errorf(d)
		}
	}
}
