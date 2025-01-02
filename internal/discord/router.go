package discord

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/bytedance/sonic"
)

type Router struct {
	cfg      Config
	session  *discordgo.Session
	logger   *slog.Logger
	ctx      context.Context
	commands map[string]Command
}

type RouterOpt func(r *Router)

func WithLogger(logger *slog.Logger) RouterOpt {
	return func(router *Router) {
		router.logger = logger
	}
}

func WithContext(ctx context.Context) RouterOpt {
	return func(router *Router) {
		router.ctx = ctx
	}
}

func NewRouter(cfg Config, opts ...RouterOpt) (*Router, error) {
	r := &Router{
		cfg:      cfg,
		logger:   slog.Default(),
		ctx:      context.Background(),
		commands: make(map[string]Command),
	}

	for _, opt := range opts {
		opt(r)
	}

	session, err := r.configureSession(cfg.Token)
	if err != nil {
		return nil, err
	}

	r.session = session

	return r, nil
}

func (r *Router) Close() error {
	if err := r.UnsetCommands(); err != nil {
		return fmt.Errorf("failed to unset commands: %w", err)
	}

	if err := r.session.Close(); err != nil {
		return fmt.Errorf("failed to close Discord session: %w", err)
	}

	return nil
}

func (r *Router) Session() *discordgo.Session {
	return r.session
}

func (r *Router) SetCommands(commands ...Command) error {
	if err := r.UnsetCommands(); err != nil {
		return fmt.Errorf("failed to unset previous commands: %w", err)
	}

	for _, cmd := range commands {
		name := cmd.Description.Name

		createdCmd, err := r.session.ApplicationCommandCreate(
			r.session.State.Application.ID,
			r.cfg.GuildID,
			cmd.Description,
		)
		if err != nil {
			return fmt.Errorf("failed to set command '%s': %w", name, err)
		}

		r.commands[name] = Command{
			Description: createdCmd,
			Handler:     cmd.Handler,
		}
	}

	return nil
}

func (r *Router) UnsetCommands() error {
	for name, cmd := range r.commands {
		err := r.session.ApplicationCommandDelete(
			r.session.State.Application.ID,
			r.cfg.GuildID,
			cmd.Description.ID,
		)
		if err != nil {
			return fmt.Errorf("failed to remove command '%s': %w", name, err)
		}
	}

	r.commands = make(map[string]Command)

	return nil
}

func (r *Router) configureSession(token string) (*discordgo.Session, error) {
	discordgo.Marshal = sonic.Marshal
	discordgo.Unmarshal = sonic.Unmarshal

	session, err := discordgo.New("Bot " + token)
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

	session.AddHandler(r.handle)

	return session, nil
}

func (r *Router) handle(_ *discordgo.Session, i *discordgo.InteractionCreate) {
	defer func() {
		if err := recover(); err != nil {
			r.logger.ErrorContext(
				r.ctx, "panic in command handler", "error", err,
			)
		}
	}()

	cmdName := i.ApplicationCommandData().Name
	cmd, ok := r.commands[cmdName]
	if !ok {
		r.logger.WarnContext(r.ctx, "unknown command", "command", cmdName)
		return
	}

	err := cmd.Handler.Handle(r.ctx, i)
	if err != nil {
		r.logger.ErrorContext(
			r.ctx, "failed to handle command", "command", cmdName, "error", err,
		)
	}
}
