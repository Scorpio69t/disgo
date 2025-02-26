package handlers

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

// gatewayHandlerWebhooksUpdate handles discord.GatewayEventTypeWebhooksUpdate
type gatewayHandlerWebhooksUpdate struct{}

// EventType returns the raw discord.GatewayEventType
func (h *gatewayHandlerWebhooksUpdate) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeWebhooksUpdate
}

// New constructs a new payload receiver for the raw gateway event
func (h *gatewayHandlerWebhooksUpdate) New() any {
	return &discord.GatewayEventWebhooksUpdate{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *gatewayHandlerWebhooksUpdate) HandleGatewayEvent(client bot.Client, sequenceNumber int, v any) {
	payload := *v.(*discord.GatewayEventWebhooksUpdate)

	client.EventManager().DispatchEvent(&events.WebhooksUpdateEvent{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber),
		GuildId:      payload.GuildID,
		ChannelID:    payload.ChannelID,
	})
}
