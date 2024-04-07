package plugin

import (
	"context"

	"github.com/git-fal7/chatwatcher/internal/event"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func InitPlugin(ctx context.Context, proxy *proxy.Proxy) error {
	event.Init(proxy)
	return nil
}
