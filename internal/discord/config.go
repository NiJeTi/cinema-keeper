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
	session, err := discordgo.New(cs)
	if err != nil {
		log.Fatalln("failed to create Discord client:", err)
	}

	session.Identify.Intents = discordgo.IntentGuilds |
		discordgo.IntentGuildPresences |
		discordgo.IntentGuildMembers |
		discordgo.IntentGuildVoiceStates

	err = session.Open()
	if err != nil {
		log.Fatalln("failed to open Discord session:", err)
	}

	return session
}

func RegisterCommands(
	session *discordgo.Session,
	commands map[string]*Command,
	guildID types.ID,
) {
	appID := session.State.Application.ID

	for name, cmd := range commands {
		createdCmd, err := session.ApplicationCommandCreate(
			appID, guildID.String(), cmd.Description,
		)
		if err != nil {
			log.Fatalln("failed to register command:", name, err)
		}

		cmd.Description = createdCmd
	}

	session.AddHandler(
		func(_ *discordgo.Session, i *discordgo.InteractionCreate) {
			defer func() {
				if err := recover(); err != nil {
					log.Println("panic:", err)
				}
			}()

			if cmd, ok := commands[i.ApplicationCommandData().Name]; ok {
				cmd.Handler.Handle(i)
			}
		},
	)
}

func UnregisterCommands(
	session *discordgo.Session,
	commands map[string]*Command,
	guildID types.ID,
) {
	appID := session.State.Application.ID

	failedCmds := map[string]error{}
	for name, cmd := range commands {
		err := session.ApplicationCommandDelete(
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
