package console

import (
	"context"
	"crypto/tls"

	"github.com/jackc/pgx/v5/pgproto3"
	"github.com/pg-sharding/spqr/pkg/client"
	"github.com/pg-sharding/spqr/pkg/clientinteractor"
	"github.com/pg-sharding/spqr/pkg/config"
	"github.com/pg-sharding/spqr/pkg/coord"
	"github.com/pg-sharding/spqr/pkg/meta"
	"github.com/pg-sharding/spqr/pkg/models/datashards"
	"github.com/pg-sharding/spqr/pkg/models/kr"
	"github.com/pg-sharding/spqr/pkg/models/shrule"
	"github.com/pg-sharding/spqr/pkg/spqrlog"
	"github.com/pg-sharding/spqr/pkg/txstatus"
	"github.com/pg-sharding/spqr/router/qlog"
	qlogprovider "github.com/pg-sharding/spqr/router/qlog/provider"
	"github.com/pg-sharding/spqr/router/rulerouter"
	spqrparser "github.com/pg-sharding/spqr/yacc/console"
	"google.golang.org/grpc"
)

type Console interface {
	Serve(ctx context.Context, cl client.Client) error
	ProcessQuery(ctx context.Context, q string, cl client.Client) error
	Shutdown() error
}

type Local struct {
	cfg     *tls.Config
	Coord   meta.EntityMgr
	RRouter rulerouter.RuleRouter
	qlogger qlog.Qlog

	stchan chan struct{}
}

var _ Console = &Local{}

func (l *Local) Shutdown() error {
	return nil
}

func NewConsole(cfg *tls.Config, coord meta.EntityMgr, rrouter rulerouter.RuleRouter, stchan chan struct{}) (*Local, error) {
	return &Local{
		Coord:   coord,
		RRouter: rrouter,
		qlogger: qlogprovider.NewLocalQlog(),
		cfg:     cfg,
		stchan:  stchan,
	}, nil
}

type TopoCntl interface {
	kr.KeyRangeMgr
	shrule.ShardingRulesMgr
	datashards.ShardsMgr
}

func (l *Local) processQueryInternal(ctx context.Context, cli *clientinteractor.PSQLInteractor, q string) error {
	tstmt, err := spqrparser.Parse(q)
	if err != nil {
		spqrlog.Zero.Error().Err(err).Msg("")
		return err
	}

	spqrlog.Zero.Debug().
		Str("query", q).
		Type("type", tstmt).
		Msg("processQueryInternal: parsed query with type")

	return l.proxyProc(ctx, tstmt, cli)
}

func (l *Local) proxyProc(ctx context.Context, tstmt spqrparser.Statement, cli *clientinteractor.PSQLInteractor) error {
	var mgr meta.EntityMgr = l.Coord

	if !config.RouterConfig().WithCoordinator {
		return meta.Proc(ctx, tstmt, mgr, l.RRouter, cli)
	}

	switch tstmt := tstmt.(type) {
	case *spqrparser.Show:
		switch tstmt.Cmd {
		case spqrparser.RoutersStr:
			coordAddr, err := l.Coord.GetCoordinator(ctx)
			if err != nil {
				return err
			}
			conn, err := grpc.Dial(coordAddr, grpc.WithInsecure()) //nolint:all
			if err != nil {
				return err
			}
			defer conn.Close()
			mgr = coord.NewAdapter(conn)
		}

	default:
		coordAddr, err := l.Coord.GetCoordinator(ctx)
		if err != nil {
			return err
		}
		conn, err := grpc.Dial(coordAddr, grpc.WithInsecure()) //nolint:all
		if err != nil {
			return err
		}
		defer conn.Close()
		mgr = coord.NewAdapter(conn)
	}

	spqrlog.Zero.Debug().Type("mgr type", mgr).Msg("proxy proc")
	return meta.Proc(ctx, tstmt, mgr, l.RRouter, cli)
}

func (l *Local) ProcessQuery(ctx context.Context, q string, cl client.Client) error {
	return l.processQueryInternal(ctx, clientinteractor.NewPSQLInteractor(cl), q)
}

const greeting = `
		SQPR router admin console
	Here you can configure your routing rules
------------------------------------------------
	You can find documentation here 
https://github.com/pg-sharding/spqr/tree/master/docs
`

func (l *Local) Serve(ctx context.Context, cl client.Client) error {
	msgs := []pgproto3.BackendMessage{
		&pgproto3.AuthenticationOk{},
	}

	params := []string{"client_encoding", "standard_conforming_strings"}
	for _, p := range params {
		if v, ok := cl.Params()[p]; ok {
			msgs = append(msgs, &pgproto3.ParameterStatus{Name: p, Value: v})
		}
	}

	msgs = append(msgs, []pgproto3.BackendMessage{
		&pgproto3.ParameterStatus{Name: "integer_datetimes", Value: "on"},
		&pgproto3.ParameterStatus{Name: "server_version", Value: "console"},
		&pgproto3.NoticeResponse{
			Message: greeting,
		},
		&pgproto3.ReadyForQuery{
			TxStatus: byte(txstatus.TXIDLE),
		},
	}...)

	for _, msg := range msgs {
		if err := cl.Send(msg); err != nil {
			spqrlog.Zero.Error().Err(err).Msg("")
			return err
		}
	}

	spqrlog.Zero.Info().Msg("console.ProcClient start")

	for {
		msg, err := cl.Receive()

		if err != nil {
			return err
		}

		switch v := msg.(type) {
		case *pgproto3.Query:
			if err := l.ProcessQuery(ctx, v.String, cl); err != nil {
				_ = cl.ReplyErrMsg(err.Error())
				// continue to consume input
			}
		case *pgproto3.Terminate:
			return nil
		default:
			spqrlog.Zero.Info().
				Type("message type", v).
				Msg("got unexpected postgresql proto message with type")
		}
	}
}

func (l *Local) Qlog() qlog.Qlog {
	return l.qlogger
}
