package plugin

import (
	"context"

	"github.com/git-fal7/chatwatcher/internal/config"
	"github.com/git-fal7/chatwatcher/internal/event"
	"github.com/git-fal7/chatwatcher/internal/profanityfilter"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func InitPlugin(ctx context.Context, proxy *proxy.Proxy) error {
	config.InitConfig()
	profanityfilter.InitProfanityFilter()
	event.Init(proxy)
	return nil
}
