package event

import (
	goaway "github.com/TwiN/go-away"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func chatEvent() func(*proxy.PlayerChatEvent) {
	return func(e *proxy.PlayerChatEvent) {
		message := e.Message()

		e.SetMessage(goaway.Censor(message))
	}
}
