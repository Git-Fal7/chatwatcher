package event

import (
	"github.com/git-fal7/chatwatcher/internal/profanityfilter"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func chatEvent() func(*proxy.PlayerChatEvent) {
	return func(e *proxy.PlayerChatEvent) {
		message := e.Message()

		e.SetMessage(profanityfilter.SanitizeMessage(message))
	}
}
