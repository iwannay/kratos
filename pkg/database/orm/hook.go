package orm

import (
	"context"

	"github.com/iwannay/kratos/pkg/net/trace"
	"xorm.io/xorm/contexts"
)

var _ contexts.Hook = &OtelHook{}

const (
	family = "sql_client"
)

type OtelHook struct {
	Tenant string
	t      trace.Trace
}

func NewOtelHook(tenant string) *OtelHook {
	if tenant == "" {
		tenant = "default"
	}
	hk := &OtelHook{
		Tenant: tenant,
	}
	return hk
}

func (oh *OtelHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	client := c.Ctx.Value(clientInstance).(*Client)
	t, ok := trace.FromContext(c.Ctx)
	if ok {
		t = t.Fork(family, "begin")
		t.SetTag(trace.String(trace.TagAddress, client.GetOpts().Tenant), trace.String(trace.TagComment, ""))
	}

	return c.Ctx, nil
}
func (oh *OtelHook) AfterProcess(c *contexts.ContextHook) error {
	client := c.Ctx.Value(clientInstance).(*Client)
	if t, ok := trace.FromContext(c.Ctx); ok {
		t = t.Fork(family, "exec")
		t.SetTag(trace.String(trace.TagAddress, client.GetOpts().Tenant), trace.String(trace.TagComment, c.SQL))
		defer t.Finish(&c.Err)
	}
	return nil
}
