package handlers

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type gatewayHandlerThreadCreate struct{}

func (h *gatewayHandlerThreadCreate) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeThreadCreate
}

func (h *gatewayHandlerThreadCreate) New() any {
	return &discord.GatewayEventThreadCreate{}
}

func (h *gatewayHandlerThreadCreate) HandleGatewayEvent(client bot.Client, sequenceNumber int, v any) {
	payload := *v.(*discord.GatewayEventThreadCreate)

	client.Caches().Channels().Put(payload.ID(), payload.GuildThread)
	client.Caches().ThreadMembers().Put(payload.ID(), payload.ThreadMember.UserID, payload.ThreadMember)

	client.EventManager().DispatchEvent(&events.ThreadCreateEvent{
		GenericThreadEvent: &events.GenericThreadEvent{
			GenericEvent: events.NewGenericEvent(client, sequenceNumber),
			ThreadID:     payload.ID(),
			GuildID:      payload.GuildID(),
			Thread:       payload.GuildThread,
		},
		ThreadMember: payload.ThreadMember,
	})
}
