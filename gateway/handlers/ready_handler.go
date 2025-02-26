package handlers

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

// gatewayHandlerReady handles discord.GatewayEventTypeReady
type gatewayHandlerReady struct{}

// EventType returns the gateway.EventType
func (h *gatewayHandlerReady) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeReady
}

// New constructs a new payload receiver for the raw gateway event
func (h *gatewayHandlerReady) New() any {
	return &discord.GatewayEventReady{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *gatewayHandlerReady) HandleGatewayEvent(client bot.Client, sequenceNumber int, v any) {
	readyEvent := *v.(*discord.GatewayEventReady)

	var shardID int
	if readyEvent.Shard != nil {
		shardID = readyEvent.Shard[0]
	}

	client.HandleReadyEvent(readyEvent)

	for _, guild := range readyEvent.Guilds {
		client.Caches().Guilds().SetUnready(shardID, guild.ID)
	}

	client.EventManager().DispatchEvent(&events.ReadyEvent{
		GenericEvent:      events.NewGenericEvent(client, sequenceNumber),
		GatewayEventReady: readyEvent,
	})

}
