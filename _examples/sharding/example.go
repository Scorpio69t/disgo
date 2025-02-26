package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/sharding"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
)

var (
	token = os.Getenv("disgo_token")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetLevel(log.LevelInfo)
	log.Info("starting example...")
	log.Info("disgo version: ", disgo.Version)

	client, err := disgo.New(token,
		bot.WithShardManagerConfigOpts(
			sharding.WithShards(0, 1, 2),
			sharding.WithShardCount(3),
			sharding.WithGatewayConfigOpts(
				gateway.WithGatewayIntents(discord.GatewayIntentGuilds, discord.GatewayIntentGuildMessages, discord.GatewayIntentDirectMessages),
				gateway.WithCompress(true),
			),
		),
		bot.WithCacheConfigOpts(cache.WithCacheFlags(cache.FlagsDefault)),
		bot.WithEventListeners(&events.ListenerAdapter{
			OnMessageCreate: onMessageCreate,
			OnGuildReady: func(event *events.GuildReadyEvent) {
				log.Infof("guild %s ready", event.GuildID)
			},
			OnGuildsReady: func(event *events.GuildsReadyEvent) {
				log.Infof("guilds on shard %d ready", event.ShardID)
			},
		}),
	)
	if err != nil {
		log.Fatalf("error while building disgo: %s", err)
	}

	defer client.Close(context.TODO())

	if err = client.ConnectShardManager(context.TODO()); err != nil {
		log.Fatal("error while connecting to gateway: ", err)
	}

	log.Infof("example is now running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-s
}

func onMessageCreate(event *events.MessageCreateEvent) {
	if event.Message.Author.Bot {
		return
	}
	_, _ = event.Client().Rest().Channels().CreateMessage(event.ChannelID, discord.NewMessageCreateBuilder().SetContent(event.Message.Content).Build())
}
