package orm

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var mux sync.RWMutex
var clients = make(map[string]*Client)
var tracerName = "github.com/kratos/pkg/database/db"

type Client struct {
	telnet    string
	opts      atomic.Value
	incr      int64
	currCount int64
	*xorm.Engine
	lastUsed time.Time
}

type Options struct {
	Tenant string
	Driver string
	Dsn    string
	Cb     func(*Client)
}

func (opt *Options) Sign() string {
	return fmt.Sprintf("%s-%s-%s", opt.Tenant, opt.Driver, opt.Dsn)
}

func (c *Client) Switch(opts *Options) *Client {
	var (
		cli *Client
		ok  bool
	)
	index := opts.Sign()
	mux.RLock()
	cli, ok = clients[index]
	mux.RUnlock()
	if !ok {
		cli = New(opts)
	}
	cli.telnet = opts.Tenant
	cli.lastUsed = time.Now()
	atomic.AddInt64(&cli.incr, 1)
	return cli
}

func (c *Client) GetOpts() *Options {
	return c.opts.Load().(*Options)
}

type ctxKey string

var clientInstance = ctxKey("__ci__")

func New(opts *Options) *Client {
	var (
		err error
		cli Client
	)
	cli.Engine, err = xorm.NewEngine(opts.Driver, opts.Dsn)
	if err != nil {
		panic(err)
	}
	cli.SetDefaultContext(
		context.WithValue(context.Background(), clientInstance, &cli))

	cli.opts.Store(opts)

	cli.AddHook(NewOtelHook(opts.Tenant))
	if opts.Cb != nil {
		opts.Cb(&cli)
	}
	mux.Lock()
	clients[opts.Sign()] = &cli
	mux.Unlock()
	return &cli
}
