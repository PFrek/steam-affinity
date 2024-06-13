package urls

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLastChar(t *testing.T) {
	type testCase struct {
		input    URLChain
		expected string
	}

	tests := map[string]testCase{
		"empty string":      {input: URLChain{URL: ""}, expected: ""},
		"length 1 string":   {input: URLChain{URL: "f"}, expected: "f"},
		"number at the end": {input: URLChain{URL: "test123"}, expected: "3"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := tc.input.lastChar()

			diff := cmp.Diff(result, tc.expected)

			if diff != "" {
				t.Fatal(diff)
			}

		})
	}
}

func TestChain(t *testing.T) {
	type testCase struct {
		base     URLChain
		input    string
		expected string
	}

	tests := map[string]testCase{
		"base with / empty arg":  {base: URLChain{URL: "/base/"}, input: "", expected: "/base/"},
		"base no / empty arg":    {base: URLChain{URL: "/base"}, input: "", expected: "/base/"},
		"base with / arg no /":   {base: URLChain{URL: "/base/"}, input: "path", expected: "/base/path"},
		"base no / arg no /":     {base: URLChain{URL: "/base"}, input: "path", expected: "/base/path"},
		"base no / arg with /":   {base: URLChain{URL: "/base"}, input: "/path", expected: "/base/path"},
		"base with / arg with /": {base: URLChain{URL: "/base/"}, input: "/path", expected: "/base/path"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := tc.base.chain(tc.input)

			diff := cmp.Diff(result.URL, tc.expected)

			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMultipleChains(t *testing.T) {
	type testCase struct {
		base     URLChain
		inputs   []string
		expected string
	}

	base := URLChain{URL: "/base"}

	result := base.chain("api").chain("v1").chain("users")

	diff := cmp.Diff(result.URL, "/base/api/v1/users")

	if diff != "" {
		t.Fatal(diff)
	}

}
