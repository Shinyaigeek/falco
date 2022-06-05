package ast

import (
	"bytes"
	"encoding/json"

	"github.com/ysugimoto/falco/token"
)

// VCL is a root of program
type VCL struct {
	Statements []Statement
}

func (v *VCL) String() string {
	var buf bytes.Buffer

	for i := range v.Statements {
		buf.WriteString(v.Statements[i].String())
	}

	return buf.String()
}

func (v *VCL) JSONString() string {
	var buf bytes.Buffer

	buf.WriteString("[")

	for i := range v.Statements {
		if i != 0 {
			buf.WriteString(",")
		}
		s, err := json.Marshal(v.Statements[i])
		if err != nil {
			panic(err)
		}
		json.Indent(&buf, []byte(s), "", "  ")
	}

	buf.WriteString("]")

	return buf.String()
}

func (v *VCL) GetMeta() *Meta {
	return New(token.Null, 0)
}
