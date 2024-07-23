package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/bytedance/sonic"

	"github.com/nijeti/cinema-keeper/internal/types"
)

type Config struct {
	Token string   `conf:"token"`
	Guild types.ID `conf:"guild"`
}

func Connect(token string) *discordgo.Session {
	discordgo.Marshal = sonic.Marshal
	discordgo.Unmarshal = sonic.Unmarshal

	cs := fmt.Sprintf("Bot %s", token)
	discord, err := discordgo.New(cs)
	if err != nil {
		log.Fatalln("failed to create discord client:", err)
	}

	discord.Identify.Intents = discordgo.IntentGuilds |
		discordgo.IntentGuildPresences |
		discordgo.IntentGuildMembers |
		discordgo.IntentGuildVoiceStates

	err = discord.Open()
	if err != nil {
		log.Fatalln("failed to open Discord session:", err)
	}

	return discord
}

func RegisterCommands(
	discord *discordgo.Session,
	cmds map[string]*Command,
	guildID types.ID,
) {
	appID := discord.State.Application.ID

	for name, cmd := range cmds {
		createdCmd, err := discord.ApplicationCommandCreate(
			appID, guildID.String(), cmd.Description,
		)
		if err != nil {
			log.Fatalln("failed to register command:", name, err)
		}

		cmd.Description = createdCmd
	}

	discord.AddHandler(
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			defer func() {
				if err := recover(); err != nil {
					log.Println("panic:", err)
				}
			}()

			if cmd, ok := cmds[i.ApplicationCommandData().Name]; ok {
				cmd.Handler.Handle(s, i)
			}
		},
	)
}

func UnregisterCommands(
	discord *discordgo.Session,
	cmds map[string]*Command,
	guildID types.ID,
) {
	appID := discord.State.Application.ID

	failedCmds := map[string]error{}
	for name, cmd := range cmds {
		err := discord.ApplicationCommandDelete(
			appID, guildID.String(), cmd.Description.ID,
		)
		if err != nil {
			failedCmds[name] = err
		}
	}

	if len(failedCmds) != 0 {
		log.Fatalln("failed to unregister handlers:", failedCmds)
	}
}
