package str_test

import (
	"context"
	"testing"

	"github.com/TheGrizzlyDev/yuba/internal/provider/str"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPadLeft(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"pad-left pads N times char": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue("a"),
					types.Int32Value(6),
					types.StringValue("i"),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("iiiiia")),
			},
		},
		"keeps string intact if desired-len >= len(string)": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue("aaaaaaaaa"),
					types.Int32Value(6),
					types.StringValue("i"),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("aaaaaaaaa")),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := function.RunResponse{
				Result: function.NewResultData(types.StringUnknown()),
			}

			str.PadLeft{}.Run(context.Background(), testCase.request, &got)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"pad-right pads N times char": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue("a"),
					types.Int32Value(6),
					types.StringValue("i"),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("aiiiii")),
			},
		},
		"keeps string intact if desired-len >= len(string)": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue("aaaaaaaaa"),
					types.Int32Value(6),
					types.StringValue("i"),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("aaaaaaaaa")),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := function.RunResponse{
				Result: function.NewResultData(types.StringUnknown()),
			}

			str.PadRight{}.Run(context.Background(), testCase.request, &got)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
