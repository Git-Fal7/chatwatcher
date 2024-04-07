package plugin

import (
	"context"

	"github.com/git-fal7/chatwatcher/internal/event"
	profanityfilter "github.com/git-fal7/chatwatcher/internal/profanityfilter"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func InitPlugin(ctx context.Context, proxy *proxy.Proxy) error {
	profanityfilter.InitProfanityFilter()
	event.Init(proxy)
	return nil
}
