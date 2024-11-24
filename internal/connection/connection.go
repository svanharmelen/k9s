package connection

import (
	"context"
	"sync"

	"github.com/derailed/k9s/internal/config"
)

var mgr *manager

type manager struct {
	context   context.Context
	cancelFn  context.CancelFunc
	waitGroup sync.WaitGroup
}

func init() {
	ctx, cancelFn := context.WithCancel(context.Background())
	mgr = &manager{
		context:  ctx,
		cancelFn: cancelFn,
	}
}

func Start(conn *config.Connection) (string, error) {
	for _, cmd := range conn.Commands {
		mgr.waitGroup.Add(1)
		if err := RunCommand(mgr.context, cmd.Command, mgr.waitGroup.Done); err != nil {
			return "", err
		}

		if cmd.WaitForPort != 0 {
			if err := WaitForPort(mgr.context, cmd.WaitForPort); err != nil {
				return "", err
			}
		}
	}

	return conn.KubeConfig, nil
}

func Stop() {
	mgr.cancelFn()
	mgr.waitGroup.Wait()
}
