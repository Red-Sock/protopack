package rules_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yoheimuta/go-protoparser/v4/parser/meta"

	"go.redsock.ru/protopack/internal/core"
	"go.redsock.ru/protopack/internal/rules"
)

func TestRPCRequestStandardName_Message(t *testing.T) {
	t.Parallel()

	assert := require.New(t)

	const expMessage = "rpc request should have suffix 'Request'"

	rule := rules.RPCRequestStandardName{}
	message := rule.Message()

	assert.Equal(expMessage, message)
}

func TestRPCRequestStandardName_Validate(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		fileName   string
		wantIssues *core.Issue
		wantErr    error
	}{
		"invalid": {
			fileName: invalidAuthProto,
			wantIssues: &core.Issue{
				Position: meta.Position{
					Filename: "",
					Offset:   214,
					Line:     11,
					Column:   3,
				},
				SourceName: "SessionInfo",
				Message:    "rpc request should have suffix 'Request'",
				RuleName:   "RPC_REQUEST_STANDARD_NAME",
			},
		},
		"valid": {
			fileName: validAuthProto,
			wantErr:  nil,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			r, protos := start(t)

			rule := rules.RPCRequestStandardName{}
			issues, err := rule.Validate(protos[tc.fileName])
			r.ErrorIs(err, tc.wantErr)
			switch {
			case tc.wantIssues != nil:
				r.Contains(issues, *tc.wantIssues)
			case len(issues) > 0:
				r.Empty(issues)
			}
		})
	}
}
