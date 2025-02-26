package handlers

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

// gatewayHandlerGuildBanAdd handles discord.GatewayEventTypeIntegrationUpdate
type gatewayHandlerIntegrationUpdate struct{}

// EventType returns the discord.GatewayEventType
func (h *gatewayHandlerIntegrationUpdate) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeIntegrationUpdate
}

// New constructs a new payload receiver for the raw gateway event
func (h *gatewayHandlerIntegrationUpdate) New() any {
	return &discord.GatewayEventIntegrationCreate{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *gatewayHandlerIntegrationUpdate) HandleGatewayEvent(client bot.Client, sequenceNumber int, v any) {
	payload := *v.(*discord.GatewayEventIntegrationUpdate)

	client.EventManager().DispatchEvent(&events.IntegrationUpdateEvent{
		GenericIntegrationEvent: &events.GenericIntegrationEvent{
			GenericEvent: events.NewGenericEvent(client, sequenceNumber),
			GuildID:      payload.GuildID,
			Integration:  payload.Integration,
		},
	})
}
