package rest

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/rest/route"
	"github.com/disgoorg/snowflake"
)

var _ StageInstances = (*stageInstanceImpl)(nil)

func NewStageInstances(restClient Client) StageInstances {
	return &stageInstanceImpl{restClient: restClient}
}

type StageInstances interface {
	GetStageInstance(guildID snowflake.Snowflake, opts ...RequestOpt) (*discord.StageInstance, error)
	CreateStageInstance(stageInstanceCreate discord.StageInstanceCreate, opts ...RequestOpt) (*discord.StageInstance, error)
	UpdateStageInstance(guildID snowflake.Snowflake, stageInstanceUpdate discord.StageInstanceUpdate, opts ...RequestOpt) (*discord.StageInstance, error)
	DeleteStageInstance(guildID snowflake.Snowflake, opts ...RequestOpt) error
}

type stageInstanceImpl struct {
	restClient Client
}

func (s *stageInstanceImpl) GetStageInstance(guildID snowflake.Snowflake, opts ...RequestOpt) (stageInstance *discord.StageInstance, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.GetStageInstance.Compile(nil, guildID)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, nil, &stageInstance, opts...)
	return
}

func (s *stageInstanceImpl) CreateStageInstance(stageInstanceCreate discord.StageInstanceCreate, opts ...RequestOpt) (stageInstance *discord.StageInstance, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.CreateStageInstance.Compile(nil)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, stageInstanceCreate, &stageInstance, opts...)
	return
}

func (s *stageInstanceImpl) UpdateStageInstance(guildID snowflake.Snowflake, stageInstanceUpdate discord.StageInstanceUpdate, opts ...RequestOpt) (stageInstance *discord.StageInstance, err error) {
	var compiledRoute *route.CompiledAPIRoute
	compiledRoute, err = route.UpdateStageInstance.Compile(nil, guildID)
	if err != nil {
		return
	}
	err = s.restClient.Do(compiledRoute, stageInstanceUpdate, &stageInstance, opts...)
	return
}

func (s *stageInstanceImpl) DeleteStageInstance(guildID snowflake.Snowflake, opts ...RequestOpt) error {
	compiledRoute, err := route.DeleteStageInstance.Compile(nil, guildID)
	if err != nil {
		return err
	}
	return s.restClient.Do(compiledRoute, nil, nil, opts...)
}
