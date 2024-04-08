package event

import (
	"fmt"
	"strings"
	"time"

	"github.com/git-fal7/chatwatcher/internal/config"
	"github.com/git-fal7/chatwatcher/internal/profanityfilter"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/util/uuid"
)

var cooldownMap = make(map[uuid.UUID]time.Time, 0)

func chatEvent() func(*proxy.PlayerChatEvent) {
	return func(e *proxy.PlayerChatEvent) {
		if config.ViperConfig.GetBool("profanityfilter.enabled") {
			if !e.Player().HasPermission(config.ViperConfig.GetString("profanityfilter.permission")) {
				e.SetMessage(profanityfilter.SanitizeMessage(e.Message()))
			}
		}

		if config.ViperConfig.GetBool("antispam.enabled") {
			if !e.Player().HasPermission(config.ViperConfig.GetString("antispam.permission")) {
				if time.Now().Before(cooldownMap[e.Player().ID()]) {
					e.SetAllowed(false)
					e.Player().SendMessage(&component.Text{
						Content: strings.ReplaceAll(config.ViperConfig.GetString("antispam.message"), "%duration%", fmt.Sprintf("%03v", time.Until(cooldownMap[e.Player().ID()]))),
					})
				} else {
					cooldownMap[e.Player().ID()] = time.Now().Add(time.Second * time.Duration(config.ViperConfig.GetFloat64("antispam.cooldown")))
				}
			}
		}
	}
}
