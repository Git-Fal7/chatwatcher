package chatwatcher

import (
	"github.com/git-fal7/chatwatcher/internal/plugin"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

var Plugin = proxy.Plugin{
	Name: "ChatWatcher",
	Init: plugin.InitPlugin,
}
