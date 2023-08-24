// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of uuid.is_version3
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-is-version3/
func Test_Uuid_is_version3(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{input: "3f22bcdf-f888-31a6-9575-d1588cb14ff4", expect: true},  // version 3
		{input: "02201c6d-57a6-479f-8e83-7d7a6f55e2bd", expect: false}, // version 4
		{input: "86573da0-058f-5871-a5b7-f3cb33447360", expect: false}, // version 5
	}

	for i, tt := range tests {
		ret, err := Uuid_is_version3(&context.Context{}, &value.String{Value: tt.input})
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if ret.Type() != value.BooleanType {
			t.Errorf("[%d] Unexpected return type, expect=BOOL, got=%s", i, ret.Type())
		}
		v := value.Unwrap[*value.Boolean](ret)
		if diff := cmp.Diff(tt.expect, v.Value); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}