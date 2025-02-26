package handlers

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/snowflake"
)

// gatewayHandlerGuildDelete handles discord.GatewayEventTypeGuildDelete
type gatewayHandlerGuildDelete struct{}

// EventType returns the discord.GatewayEventType
func (h *gatewayHandlerGuildDelete) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeGuildDelete
}

// New constructs a new payload receiver for the raw gateway event
func (h *gatewayHandlerGuildDelete) New() any {
	return &discord.UnavailableGuild{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *gatewayHandlerGuildDelete) HandleGatewayEvent(client bot.Client, sequenceNumber int, v any) {
	unavailableGuild := *v.(*discord.UnavailableGuild)

	guild, _ := client.Caches().Guilds().Remove(unavailableGuild.ID)
	client.Caches().VoiceStates().RemoveAll(unavailableGuild.ID)
	client.Caches().Presences().RemoveAll(unavailableGuild.ID)
	client.Caches().Channels().RemoveIf(func(channel discord.Channel) bool {
		if guildChannel, ok := channel.(discord.GuildChannel); ok {
			return guildChannel.GuildID() == unavailableGuild.ID
		}
		return false
	})
	client.Caches().Emojis().RemoveAll(unavailableGuild.ID)
	client.Caches().Stickers().RemoveAll(unavailableGuild.ID)
	client.Caches().Roles().RemoveAll(unavailableGuild.ID)
	client.Caches().StageInstances().RemoveAll(unavailableGuild.ID)
	client.Caches().ThreadMembers().RemoveIf(func(groupID snowflake.Snowflake, threadMember discord.ThreadMember) bool {
		// TODO: figure out how to remove thread members from cache via guild id
		return false
	})
	client.Caches().Messages().RemoveIf(func(channelID snowflake.Snowflake, message discord.Message) bool {
		return message.GuildID != nil && *message.GuildID == unavailableGuild.ID
	})

	if unavailableGuild.Unavailable {
		client.Caches().Guilds().SetUnavailable(unavailableGuild.ID)
	}

	genericGuildEvent := &events.GenericGuildEvent{
		GenericEvent: events.NewGenericEvent(client, sequenceNumber),
		GuildID:      unavailableGuild.ID,
		Guild:        guild,
	}

	if unavailableGuild.Unavailable {
		client.EventManager().DispatchEvent(&events.GuildUnavailableEvent{
			GenericGuildEvent: genericGuildEvent,
		})
	} else {
		client.EventManager().DispatchEvent(&events.GuildLeaveEvent{
			GenericGuildEvent: genericGuildEvent,
		})
	}
}
