package events

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake"
)

type GenericGuildScheduledEventEvent struct {
	*GenericEvent
	GuildScheduledEvent discord.GuildScheduledEvent
}

type GuildScheduledEventCreateEvent struct {
	*GenericGuildScheduledEventEvent
}

type GuildScheduledEventUpdateEvent struct {
	*GenericGuildScheduledEventEvent
	OldGuildScheduledEvent discord.GuildScheduledEvent
}

type GuildScheduledEventDeleteEvent struct {
	*GenericGuildScheduledEventEvent
}

type GenericGuildScheduledEventUserEvent struct {
	*GenericEvent
	GuildScheduledEventID snowflake.Snowflake
	UserID                snowflake.Snowflake
	GuildID               snowflake.Snowflake
}

func (e *GenericGuildScheduledEventUserEvent) GuildScheduledEvent() (discord.GuildScheduledEvent, bool) {
	return e.Client().Caches().GuildScheduledEvents().Get(e.GuildID, e.GuildScheduledEventID)
}

func (e *GenericGuildScheduledEventUserEvent) Member() (discord.Member, bool) {
	return e.Client().Caches().Members().Get(e.GuildID, e.UserID)
}

type GuildScheduledEventUserAddEvent struct {
	*GenericGuildScheduledEventUserEvent
}

type GuildScheduledEventUserRemoveEvent struct {
	*GenericGuildScheduledEventUserEvent
}
