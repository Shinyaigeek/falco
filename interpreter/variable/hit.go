package variable

import (
	"bytes"
	"io"
	"strings"
	"time"

	"net/http"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

type HitScopeVariables struct {
	Variable
	base *AllScopeVariables
	ctx  *context.Context
}

func NewHitScopeVariables(ctx *context.Context) *HitScopeVariables {
	return &HitScopeVariables{
		base: NewAllScopeVariables(ctx),
		ctx:  ctx,
	}
}

func (v *HitScopeVariables) Get(s context.Scope, name string) (value.Value, error) {
	switch name {
	case OBJ_AGE:
		// fixed value
		return &value.RTime{Value: 60 * time.Second}, nil
	case OBJ_CACHEABLE:
		// always true
		return &value.Boolean{Value: true}, nil
	case OBJ_ENTERED:
		return &value.RTime{Value: 60 * time.Second}, nil
	case OBJ_GRACE:
		return v.ctx.ObjectGrace, nil
	case OBJ_HITS:
		return &value.Integer{Value: 1}, nil
	case OBJ_IS_PCI:
		return &value.Boolean{Value: false}, nil
	case OBJ_LASTUSE:
		return &value.RTime{Value: 60 * time.Second}, nil
	case OBJ_PROTO:
		return &value.String{Value: v.ctx.BackendResponse.Proto}, nil
	case OBJ_RESPONSE:
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(v.ctx.Object.Body); err != nil {
			return value.Null, errors.WithStack(err)
		}
		v.ctx.Object.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
		return &value.String{Value: buf.String()}, nil
	case OBJ_STALE_IF_ERROR:
		// alias for obj.grace
		return v.ctx.ObjectGrace, nil
	case OBJ_STALE_WHILE_REVALIDATE:
		return &value.RTime{Value: 60 * time.Second}, nil
	case OBJ_STATUS:
		return &value.Integer{Value: int64(v.ctx.Object.StatusCode)}, nil
	case OBJ_TTL:
		return v.ctx.ObjectTTL, nil
	// Digest ratio will return fixed value
	case REQ_DIGEST_RATIO:
		return &value.Float{Value: 0.4}, nil
	}

	if val := v.getFromRegex(name); val != nil {
		return val, nil
	}

	// If not found, also look up all scope value
	val, err := v.base.Get(s, name)
	if err != nil {
		return value.Null, errors.WithStack(err)
	}
	return val, nil
}

func (v *HitScopeVariables) getFromRegex(name string) value.Value {
	// HTTP request header matching
	if match := objectHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return &value.String{
			Value: v.ctx.Object.Header.Get(match[1]),
		}
	}
	return v.base.getFromRegex(name)
}

func (v *HitScopeVariables) Set(s context.Scope, name, operator string, val value.Value) error {
	switch name {
	case OBJ_GRACE:
		if err := doAssign(v.ctx.ObjectGrace, operator, val); err != nil {
			return errors.WithStack(err)
		}
		return nil
	case OBJ_RESPONSE:
		v.ctx.Object.Body = io.NopCloser(strings.NewReader(val.String()))
		return nil
	case OBJ_STATUS:
		i := &value.Integer{Value: 0}
		if err := doAssign(i, operator, val); err != nil {
			return errors.WithStack(err)
		}
		v.ctx.Object.StatusCode = int(i.Value)
		v.ctx.Object.Status = http.StatusText(int(i.Value))
		return nil
	case OBJ_TTL:
		if err := doAssign(v.ctx.ObjectTTL, operator, val); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}

	if match := objectHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		v.ctx.Object.Header.Set(match[1], val.String())
		return nil
	}

	// If not found, pass to all scope value
	return v.base.Set(s, name, operator, val)
}

func (v *HitScopeVariables) Add(s context.Scope, name string, val value.Value) error {
	// Add statement could be use only for HTTP header
	match := objectHttpHeaderRegex.FindStringSubmatch(name)
	if match == nil {
		// Nothing values to be enable to add in PASS, pass to base
		return v.base.Add(s, name, val)
	}

	v.ctx.Object.Header.Add(match[1], val.String())
	return nil
}

func (v *HitScopeVariables) Unset(s context.Scope, name string) error {
	match := objectHttpHeaderRegex.FindStringSubmatch(name)
	if match == nil {
		// Nothing values to be enable to unset in PASS, pass to base
		return v.base.Unset(s, name)
	}
	v.ctx.Object.Header.Del(match[1])
	return nil
}