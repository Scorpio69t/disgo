package rest

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/rest/route"
	"github.com/disgoorg/snowflake"
)

var _ GuildTemplates = (*guildTemplateImpl)(nil)

func NewGuildTemplates(restClient Client) GuildTemplates {
	return &guildTemplateImpl{restClient: restClient}
}

type GuildTemplates interface {
	GetGuildTemplate(templateCode string, opts ...RequestOpt) (*discord.GuildTemplate, error)
	GetGuildTemplates(guildID snowflake.Snowflake, opts ...RequestOpt) ([]discord.GuildTemplate, error)
	CreateGuildTemplate(guildID snowflake.Snowflake, guildTemplateCreate discord.GuildTemplateCreate, opts ...RequestOpt) (*discord.GuildTemplate, error)
	CreateGuildFromTemplate(templateCode string, createGuildFromTemplate discord.GuildFromTemplateCreate, opts ...RequestOpt) (*discord.Guild, error)
	SyncGuildTemplate(guildID snowflake.Snowflake, templateCode string, opts ...RequestOpt) (*discord.GuildTemplate, error)
	UpdateGuildTemplate(guildID snowflake.Snowflake, templateCode string, guildTemplateUpdate discord.GuildTemplateUpdate, opts ...RequestOpt) (*discord.GuildTemplate, error)
	DeleteGuildTemplate(guildID snowflake.Snowflake, templateCode string, opts ...RequestOpt) (*discord.GuildTemplate, error)
}

type guildTemplateImpl struct {
	restClient Client
}

func (s *guildTemplateImpl) GetGuildTemplate(templateCode string, opts ...RequestOpt) (guildTemplate *discord.GuildTemplate, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.GetGuildTemplate.Compile(nil, templateCode)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, nil, &guildTemplate, opts...)
	return
}

func (s *guildTemplateImpl) GetGuildTemplates(guildID snowflake.Snowflake, opts ...RequestOpt) (guildTemplates []discord.GuildTemplate, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.GetGuildTemplates.Compile(nil, guildID)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, nil, &guildTemplates, opts...)
	return
}

func (s *guildTemplateImpl) CreateGuildTemplate(guildID snowflake.Snowflake, guildTemplateCreate discord.GuildTemplateCreate, opts ...RequestOpt) (guildTemplate *discord.GuildTemplate, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.CreateGuildTemplate.Compile(nil, guildID)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, guildTemplateCreate, &guildTemplate, opts...)
	return
}

func (s *guildTemplateImpl) CreateGuildFromTemplate(templateCode string, createGuildFromTemplate discord.GuildFromTemplateCreate, opts ...RequestOpt) (guild *discord.Guild, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.CreateGuildFromTemplate.Compile(nil, templateCode)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, createGuildFromTemplate, &guild, opts...)
	return
}

func (s *guildTemplateImpl) SyncGuildTemplate(guildID snowflake.Snowflake, templateCode string, opts ...RequestOpt) (guildTemplate *discord.GuildTemplate, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.SyncGuildTemplate.Compile(nil, guildID, templateCode)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, nil, &guildTemplate, opts...)
	return
}

func (s *guildTemplateImpl) UpdateGuildTemplate(guildID snowflake.Snowflake, templateCode string, guildTemplateUpdate discord.GuildTemplateUpdate, opts ...RequestOpt) (guildTemplate *discord.GuildTemplate, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.UpdateGuildTemplate.Compile(nil, guildID, templateCode)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, guildTemplateUpdate, &guildTemplate, opts...)
	return
}

func (s *guildTemplateImpl) DeleteGuildTemplate(guildID snowflake.Snowflake, templateCode string, opts ...RequestOpt) (guildTemplate *discord.GuildTemplate, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.DeleteGuildTemplate.Compile(nil, guildID, templateCode)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, nil, &guildTemplate, opts...)
	return
}
