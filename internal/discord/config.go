package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/bytedance/sonic"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type Config struct {
	Token   string    `conf:"token"`
	GuildID models.ID `conf:"guild_id"`
}

func Connect(token string) (*discordgo.Session, error) {
	discordgo.Marshal = sonic.Marshal
	discordgo.Unmarshal = sonic.Unmarshal

	cs := fmt.Sprintf("Bot %s", token)
	session, err := discordgo.New(cs)
	if err != nil {
		return nil, fmt.Errorf("failed to create Discord client: %w", err)
	}

	session.Identify.Intents = discordgo.IntentGuilds |
		discordgo.IntentGuildPresences |
		discordgo.IntentGuildMembers |
		discordgo.IntentGuildVoiceStates

	err = session.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open Discord session: %w", err)
	}

	return session, nil
}

func RegisterCommands(
	session *discordgo.Session,
	commands map[string]*Command,
	guildID models.ID,
) error {
	appID := session.State.Application.ID

	for name, cmd := range commands {
		createdCmd, err := session.ApplicationCommandCreate(
			appID, guildID.String(), cmd.Description,
		)
		if err != nil {
			return fmt.Errorf("failed to register command '%s': %w", name, err)
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

	return nil
}

func UnregisterCommands(
	session *discordgo.Session,
	commands map[string]*Command,
	guildID models.ID,
) error {
	appID := session.State.Application.ID

	failedCmds := make(map[string]error)
	for name, cmd := range commands {
		err := session.ApplicationCommandDelete(
			appID, guildID.String(), cmd.Description.ID,
		)
		if err != nil {
			failedCmds[name] = err
		}
	}

	if len(failedCmds) != 0 {
		return fmt.Errorf("failed to unregister handlers: %v", failedCmds)
	}

	return nil
}
