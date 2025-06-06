package str

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = PadLeft{}

var (
	paddingParams = []function.Parameter{
		function.StringParameter{
			Name:           "string",
			AllowNullValue: false,
		},
		function.Int32Parameter{
			Name:           "length",
			AllowNullValue: false,
		},
		function.StringParameter{
			Name:           "padding",
			AllowNullValue: false,
			// TODO add a validator that length is 1
		},
	}
)

type PadLeft struct{}

func NewPadLeft() function.Function {
	return PadLeft{}
}

func (p PadLeft) Definition(ctx context.Context, req function.DefinitionRequest, res *function.DefinitionResponse) {
	res.Definition = function.Definition{
		Summary:    "Pad to the left a string to a given size using a given character",
		Parameters: paddingParams,
	}
}

func (p PadLeft) Metadata(ctx context.Context, req function.MetadataRequest, res *function.MetadataResponse) {
	res.Name = "pad_left"
}

func (p PadLeft) Run(ctx context.Context, req function.RunRequest, res *function.RunResponse) {
	padInDirection(true, ctx, req, res)
}

var _ function.Function = PadRight{}

type PadRight struct{}

func NewPadRight() function.Function {
	return PadRight{}
}

func (p PadRight) Definition(ctx context.Context, req function.DefinitionRequest, res *function.DefinitionResponse) {
	res.Definition = function.Definition{
		Summary:    "Pad to the right a string to a given size using a given character",
		Parameters: paddingParams,
	}
}

func (p PadRight) Metadata(ctx context.Context, req function.MetadataRequest, res *function.MetadataResponse) {
	res.Name = "pad_right"
}

func (p PadRight) Run(ctx context.Context, req function.RunRequest, res *function.RunResponse) {
	padInDirection(false, ctx, req, res)
}

func padInDirection(leftDir bool, ctx context.Context, req function.RunRequest, res *function.RunResponse) {
	var str, pad string
	var reqLen int
	res.Error = function.ConcatFuncErrors(res.Error, req.Arguments.Get(ctx, &str, &reqLen, &pad))

	if len(str) < reqLen {
		padding := strings.Repeat(pad, reqLen-len(str))
		if leftDir {
			str = padding + str
		} else {
			str = str + padding
		}
	}

	res.Error = function.ConcatFuncErrors(res.Error, res.Result.Set(ctx, str))
}
