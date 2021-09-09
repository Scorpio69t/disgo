package core

import (
	"reflect"
)

// ListenerAdapter lets you override the handles for receiving events
type ListenerAdapter struct {
	// Other events
	OnHeartbeat   func(event *HeartbeatEvent)
	OnHTTPRequest func(event *HTTPRequestEvent)
	OnRawGateway  func(event *RawEvent)
	OnReadyEvent  func(event *ReadyEvent)

	// core.GetGuildChannel Events
	OnGuildChannelCreate func(event *GuildChannelCreateEvent)
	OnGuildChannelUpdate func(event *GuildChannelUpdateEvent)
	OnGuildChannelDelete func(event *GuildChannelDeleteEvent)

	// core.DMChannel Events
	OnDMChannelCreate func(event *DMChannelCreateEvent)
	OnDMChannelUpdate func(event *DMChannelUpdateEvent)
	OnDMChannelDelete func(event *DMChannelDeleteEvent)

	// core.DMChannel Reaction Events
	OnDMMessageReactionAdd         func(event *DMMessageReactionAddEvent)
	OnDMMessageReactionRemove      func(event *DMMessageReactionRemoveEvent)
	OnDMMessageReactionRemoveEmoji func(event *DMMessageReactionRemoveEmojiEvent)
	OnDMMessageReactionRemoveAll   func(event *DMMessageReactionRemoveAllEvent)

	// core.Emoji Events
	OnEmojiCreate func(event *EmojiCreateEvent)
	OnEmojiUpdate func(event *EmojiUpdateEvent)
	OnEmojiDelete func(event *EmojiDeleteEvent)

	// core.GatewayStatus Events
	OnConnected    func(event *ConnectedEvent)
	OnReconnected  func(event *ReconnectedEvent)
	OnResumed      func(event *ResumedEvent)
	OnDisconnected func(event *DisconnectedEvent)

	// core.Guild Events
	OnGuildJoin        func(event *GuildJoinEvent)
	OnGuildUpdate      func(event *GuildUpdateEvent)
	OnGuildLeave       func(event *GuildLeaveEvent)
	OnGuildAvailable   func(event *GuildAvailableEvent)
	OnGuildUnavailable func(event *GuildUnavailableEvent)
	OnGuildBan         func(event *GuildBanEvent)
	OnGuildUnban       func(event *GuildUnbanEvent)

	// core.Guild core.Invite Events
	OnGuildInviteCreate func(event *GuildInviteCreateEvent)
	OnGuildInviteDelete func(event *GuildInviteDeleteEvent)

	// core.Guild core.Member Events
	OnGuildMemberJoin   func(event *GuildMemberJoinEvent)
	OnGuildMemberUpdate func(event *GuildMemberUpdateEvent)
	OnGuildMemberLeave  func(event *GuildMemberLeaveEvent)

	// core.Guild core.Message Events
	OnGuildMessageCreate func(event *GuildMessageCreateEvent)
	OnGuildMessageUpdate func(event *GuildMessageUpdateEvent)
	OnGuildMessageDelete func(event *GuildMessageDeleteEvent)

	// core.Guild core.Message Reaction Events
	OnGuildMessageReactionAdd         func(event *GuildMessageReactionAddEvent)
	OnGuildMessageReactionRemove      func(event *GuildMessageReactionRemoveEvent)
	OnGuildMessageReactionRemoveEmoji func(event *GuildMessageReactionRemoveEmojiEvent)
	OnGuildMessageReactionRemoveAll   func(event *GuildMessageReactionRemoveAllEvent)

	// core.Guild Voice Events
	OnGuildVoiceUpdate func(event *GuildVoiceUpdateEvent)
	OnGuildVoiceJoin   func(event *GuildVoiceJoinEvent)
	OnGuildVoiceLeave  func(event *GuildVoiceLeaveEvent)

	// core.Guild core.StageInstance Events
	OnStageInstanceCreate func(event *StageInstanceCreateEvent)
	OnStageInstanceUpdate func(event *StageInstanceUpdateEvent)
	OnStageInstanceDelete func(event *StageInstanceDeleteEvent)

	// core.Guild core.Role Events
	OnRoleCreate func(event *RoleCreateEvent)
	OnRoleUpdate func(event *RoleUpdateEvent)
	OnRoleDelete func(event *RoleDeleteEvent)

	// core.Interaction Events
	OnSlashCommand     func(event *SlashCommandEvent)
	OnUserCommand      func(event *UserCommandEvent)
	OnMessageCommand   func(event *MessageCommandEvent)
	OnButtonClick      func(event *ButtonClickEvent)
	OnSelectMenuSubmit func(event *SelectMenuSubmitEvent)

	// core.Message Events
	OnMessageCreate func(event *MessageCreateEvent)
	OnMessageUpdate func(event *MessageUpdateEvent)
	OnMessageDelete func(event *MessageDeleteEvent)

	// core.Message Reaction Events
	OnMessageReactionAdd         func(event *MessageReactionAddEvent)
	OnMessageReactionRemove      func(event *MessageReactionRemoveEvent)
	OnMessageReactionRemoveEmoji func(event *MessageReactionRemoveEmojiEvent)
	OnMessageReactionRemoveAll   func(event *MessageReactionRemoveAllEvent)

	// Self Events
	OnSelfUpdate func(event *SelfUpdateEvent)

	// core.User Events
	OnUserUpdate      func(event *UserUpdateEvent)
	OnUserTyping      func(event *UserTypingEvent)
	OnGuildUserTyping func(event *GuildMemberTypingEvent)
	OnDMUserTyping    func(event *DMChannelUserTypingEvent)

	// core.User core.Activity Events
	OnUserActivityStart  func(event *UserActivityStartEvent)
	OnUserActivityUpdate func(event *UserActivityUpdateEvent)
	OnUserActivityEnd    func(event *UserActivityEndEvent)
}

// OnEvent is getting called everytime we receive an event
func (l ListenerAdapter) OnEvent(event interface{}) {
	switch e := event.(type) {
	case *HeartbeatEvent:
		if listener := l.OnHeartbeat; listener != nil {
			listener(e)
		}
	case *HTTPRequestEvent:
		if listener := l.OnHTTPRequest; listener != nil {
			listener(e)
		}
	case *RawEvent:
		if listener := l.OnRawGateway; listener != nil {
			listener(e)
		}
	case *ReadyEvent:
		if listener := l.OnReadyEvent; listener != nil {
			listener(e)
		}

	// core.GetGuildChannel Events
	case *GuildChannelCreateEvent:
		if listener := l.OnGuildChannelCreate; listener != nil {
			listener(e)
		}
	case *GuildChannelUpdateEvent:
		if listener := l.OnGuildChannelUpdate; listener != nil {
			listener(e)
		}
	case *GuildChannelDeleteEvent:
		if listener := l.OnGuildChannelDelete; listener != nil {
			listener(e)
		}

	// core.DMChannel Events// core.Category Events
	case *DMChannelCreateEvent:
		if listener := l.OnDMChannelCreate; listener != nil {
			listener(e)
		}
	case *DMChannelUpdateEvent:
		if listener := l.OnDMChannelUpdate; listener != nil {
			listener(e)
		}
	case *DMChannelDeleteEvent:
		if listener := l.OnDMChannelDelete; listener != nil {
			listener(e)
		}

	// core.DMChannel Events// core.Category Events
	case *DMMessageReactionAddEvent:
		if listener := l.OnDMMessageReactionAdd; listener != nil {
			listener(e)
		}
	case *DMMessageReactionRemoveEvent:
		if listener := l.OnDMMessageReactionRemove; listener != nil {
			listener(e)
		}
	case *DMMessageReactionRemoveEmojiEvent:
		if listener := l.OnDMMessageReactionRemoveEmoji; listener != nil {
			listener(e)
		}
	case *DMMessageReactionRemoveAllEvent:
		if listener := l.OnDMMessageReactionRemoveAll; listener != nil {
			listener(e)
		}

	// core.Emoji Events
	case *EmojiCreateEvent:
		if listener := l.OnEmojiCreate; listener != nil {
			listener(e)
		}
	case *EmojiUpdateEvent:
		if listener := l.OnEmojiUpdate; listener != nil {
			listener(e)
		}
	case *EmojiDeleteEvent:
		if listener := l.OnEmojiDelete; listener != nil {
			listener(e)
		}

	// gateway.GatewayStatus Events
	case *ConnectedEvent:
		if listener := l.OnConnected; listener != nil {
			listener(e)
		}
	case *ReconnectedEvent:
		if listener := l.OnReconnected; listener != nil {
			listener(e)
		}
	case *ResumedEvent:
		if listener := l.OnResumed; listener != nil {
			listener(e)
		}
	case *DisconnectedEvent:
		if listener := l.OnDisconnected; listener != nil {
			listener(e)
		}

	// core.Guild Events
	case *GuildJoinEvent:
		if listener := l.OnGuildJoin; listener != nil {
			listener(e)
		}
	case *GuildUpdateEvent:
		if listener := l.OnGuildUpdate; listener != nil {
			listener(e)
		}
	case *GuildLeaveEvent:
		if listener := l.OnGuildLeave; listener != nil {
			listener(e)
		}
	case *GuildAvailableEvent:
		if listener := l.OnGuildAvailable; listener != nil {
			listener(e)
		}
	case *GuildUnavailableEvent:
		if listener := l.OnGuildUnavailable; listener != nil {
			listener(e)
		}
	case *GuildBanEvent:
		if listener := l.OnGuildBan; listener != nil {
			listener(e)
		}
	case *GuildUnbanEvent:
		if listener := l.OnGuildUnban; listener != nil {
			listener(e)
		}

	// core.Guild core.Invite Events
	case *GuildInviteCreateEvent:
		if listener := l.OnGuildInviteCreate; listener != nil {
			listener(e)
		}
	case *GuildInviteDeleteEvent:
		if listener := l.OnGuildInviteDelete; listener != nil {
			listener(e)
		}

	// core.Member Events
	case *GuildMemberJoinEvent:
		if listener := l.OnGuildMemberJoin; listener != nil {
			listener(e)
		}
	case *GuildMemberUpdateEvent:
		if listener := l.OnGuildMemberUpdate; listener != nil {
			listener(e)
		}
	case *GuildMemberLeaveEvent:
		if listener := l.OnGuildMemberLeave; listener != nil {
			listener(e)
		}

	// core.Guild core.Message Events
	case *GuildMessageCreateEvent:
		if listener := l.OnGuildMessageCreate; listener != nil {
			listener(e)
		}
	case *GuildMessageUpdateEvent:
		if listener := l.OnGuildMessageUpdate; listener != nil {
			listener(e)
		}
	case *GuildMessageDeleteEvent:
		if listener := l.OnGuildMessageDelete; listener != nil {
			listener(e)
		}

	// core.Guild core.Message Reaction Events
	case *GuildMessageReactionAddEvent:
		if listener := l.OnGuildMessageReactionAdd; listener != nil {
			listener(e)
		}
	case *GuildMessageReactionRemoveEvent:
		if listener := l.OnGuildMessageReactionRemove; listener != nil {
			listener(e)
		}
	case *GuildMessageReactionRemoveEmojiEvent:
		if listener := l.OnGuildMessageReactionRemoveEmoji; listener != nil {
			listener(e)
		}
	case *GuildMessageReactionRemoveAllEvent:
		if listener := l.OnGuildMessageReactionRemoveAll; listener != nil {
			listener(e)
		}

	// core.Guild Voice Events
	case *GuildVoiceUpdateEvent:
		if listener := l.OnGuildVoiceUpdate; listener != nil {
			listener(e)
		}
	case *GuildVoiceJoinEvent:
		if listener := l.OnGuildVoiceJoin; listener != nil {
			listener(e)
		}
	case *GuildVoiceLeaveEvent:
		if listener := l.OnGuildVoiceLeave; listener != nil {
			listener(e)
		}

	// core.Guild core.StageInstance Events
	case *StageInstanceCreateEvent:
		if listener := l.OnStageInstanceCreate; listener != nil {
			listener(e)
		}
	case *StageInstanceUpdateEvent:
		if listener := l.OnStageInstanceUpdate; listener != nil {
			listener(e)
		}
	case *StageInstanceDeleteEvent:
		if listener := l.OnStageInstanceDelete; listener != nil {
			listener(e)
		}

	// core.Guild core.Role Events
	case *RoleCreateEvent:
		if listener := l.OnRoleCreate; listener != nil {
			listener(e)
		}
	case *RoleUpdateEvent:
		if listener := l.OnRoleUpdate; listener != nil {
			listener(e)
		}
	case *RoleDeleteEvent:
		if listener := l.OnRoleDelete; listener != nil {
			listener(e)
		}

	// Interaction Events
	case *SlashCommandEvent:
		if listener := l.OnSlashCommand; listener != nil {
			listener(e)
		}
	case *UserCommandEvent:
		if listener := l.OnUserCommand; listener != nil {
			listener(e)
		}
	case *MessageCommandEvent:
		if listener := l.OnMessageCommand; listener != nil {
			listener(e)
		}
	case *ButtonClickEvent:
		if listener := l.OnButtonClick; listener != nil {
			listener(e)
		}
	case *SelectMenuSubmitEvent:
		if listener := l.OnSelectMenuSubmit; listener != nil {
			listener(e)
		}

	// core.Message Events
	case *MessageCreateEvent:
		if listener := l.OnMessageCreate; listener != nil {
			listener(e)
		}
	case *MessageUpdateEvent:
		if listener := l.OnMessageUpdate; listener != nil {
			listener(e)
		}
	case *MessageDeleteEvent:
		if listener := l.OnMessageDelete; listener != nil {
			listener(e)
		}

	// core.Message Reaction Events
	case *MessageReactionAddEvent:
		if listener := l.OnMessageReactionAdd; listener != nil {
			listener(e)
		}
	case *MessageReactionRemoveEvent:
		if listener := l.OnMessageReactionRemove; listener != nil {
			listener(e)
		}
	case *MessageReactionRemoveEmojiEvent:
		if listener := l.OnMessageReactionRemoveEmoji; listener != nil {
			listener(e)
		}
	case *MessageReactionRemoveAllEvent:
		if listener := l.OnMessageReactionRemoveAll; listener != nil {
			listener(e)
		}

	// Self Events
	case *SelfUpdateEvent:
		if listener := l.OnSelfUpdate; listener != nil {
			listener(e)
		}

	// core.User Events
	case *UserUpdateEvent:
		if listener := l.OnUserUpdate; listener != nil {
			listener(e)
		}
	case *UserTypingEvent:
		if listener := l.OnUserTyping; listener != nil {
			listener(e)
		}
	case *GuildMemberTypingEvent:
		if listener := l.OnGuildUserTyping; listener != nil {
			listener(e)
		}
	case *DMChannelUserTypingEvent:
		if listener := l.OnDMUserTyping; listener != nil {
			listener(e)
		}

	// core.User core.Activity Events
	case *UserActivityStartEvent:
		if listener := l.OnUserActivityStart; listener != nil {
			listener(e)
		}
	case *UserActivityUpdateEvent:
		if listener := l.OnUserActivityUpdate; listener != nil {
			listener(e)
		}
	case *UserActivityEndEvent:
		if listener := l.OnUserActivityEnd; listener != nil {
			listener(e)
		}

	default:
		println("OUF")
		if e, ok := e.(Event); ok {
			var name string
			if t := reflect.TypeOf(e); t.Kind() == reflect.Ptr {
				name = "*" + t.Elem().Name()
			} else {
				name = t.Name()
			}
			e.Bot().Logger.Errorf("unexpected event received: \"%s\", event: \"%#e\"", name, event)
		}
	}
}
