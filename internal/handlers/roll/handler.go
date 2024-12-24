package roll

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/discord/responses"
	"github.com/nijeti/cinema-keeper/internal/pkg/die"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordutils"
)

type Handler struct {
	ctx     context.Context
	log     *slog.Logger
	session *discordgo.Session
	utils   discordutils.Utils
}

func New(
	ctx context.Context,
	log *slog.Logger,
	session *discordgo.Session,
) *Handler {
	log = log.With("command", discord.RollName)

	return &Handler{
		ctx:     ctx,
		log:     log,
		session: session,
		utils:   discordutils.New(ctx, log, session),
	}
}

func (h *Handler) Handle(i *discordgo.InteractionCreate) {
	size := discord.RollOptionSizeDefault
	count := discord.RollOptionCountDefault

	optionsMap := discordutils.OptionsMap(i)
	if opt, ok := optionsMap[discord.RollOptionSize]; ok {
		size = die.Size(opt.IntValue())
	}
	if opt, ok := optionsMap[discord.RollOptionCount]; ok {
		count = int(opt.IntValue())
	}

	results := make([]int, 0, count)
	for range count {
		results = append(results, die.Roll(size))
	}

	var response *discordgo.InteractionResponse
	if len(results) == discord.RollOptionCountDefault {
		if size == discord.RollOptionSizeDefault {
			response = responses.RollSingleDefault(results[0])
		} else {
			response = responses.RollSingle(size, results[0])
		}
	} else {
		response = responses.RollMultiple(size, results)
	}

	_ = h.utils.Respond(i, response)
}
