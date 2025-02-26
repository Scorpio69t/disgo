package oauth2

import (
	"errors"
	"fmt"

	"github.com/disgoorg/snowflake"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/rest"
)

// errors returned by the OAuth2 client
var (
	ErrStateNotFound      = errors.New("state could not be found")
	ErrAccessTokenExpired = errors.New("access token expired. refresh the session")
	ErrMissingOAuth2Scope = func(scope discord.ApplicationScope) error {
		return fmt.Errorf("missing '%s' scope", scope)
	}
)

type Client interface {
	// ID returns the configured client ID
	ID() snowflake.Snowflake
	// Secret returns the configured client secret
	Secret() string
	// Rest returns the underlying rest.OAuth2
	Rest() rest.OAuth2

	// SessionController returns the configured SessionController
	SessionController() SessionController
	// StateController returns the configured StateController
	StateController() StateController

	// GenerateAuthorizationURL generates an authorization URL with the given redirect URI, permissions, guildID, disableGuildSelect & scopes. State is automatically generated
	GenerateAuthorizationURL(redirectURI string, permissions discord.Permissions, guildID snowflake.Snowflake, disableGuildSelect bool, scopes ...discord.ApplicationScope) string
	// GenerateAuthorizationURLState generates an authorization URL with the given redirect URI, permissions, guildID, disableGuildSelect & scopes. State is automatically generated & returned
	GenerateAuthorizationURLState(redirectURI string, permissions discord.Permissions, guildID snowflake.Snowflake, disableGuildSelect bool, scopes ...discord.ApplicationScope) (string, string)

	// StartSession starts a new Session with the given authorization code & state
	StartSession(code string, state string, identifier string, opts ...rest.RequestOpt) (Session, error)
	// RefreshSession refreshes the given Session with the refresh token
	RefreshSession(identifier string, session Session, opts ...rest.RequestOpt) (Session, error)

	// GetUser returns the discord.OAuth2User associated with the given Session. Fields filled in the struct depend on the Session.Scopes
	GetUser(session Session, opts ...rest.RequestOpt) (*discord.OAuth2User, error)
	// GetGuilds returns the discord.OAuth2Guild(s) the user is a member of. This requires the discord.ApplicationScopeGuilds scope in the Session
	GetGuilds(session Session, opts ...rest.RequestOpt) ([]discord.OAuth2Guild, error)
	// GetConnections returns the discord.Connection(s) the user has connected. This requires the discord.ApplicationScopeConnections scope in the Session
	GetConnections(session Session, opts ...rest.RequestOpt) ([]discord.Connection, error)
}
