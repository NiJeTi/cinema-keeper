package roll

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/discord/responses"
	"github.com/nijeti/cinema-keeper/internal/pkg/die"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

type Handler struct {
	ctx   context.Context
	log   *slog.Logger
	utils discordUtils.Utils
}

func New(
	ctx context.Context,
	log *slog.Logger,
) Handler {
	log = log.With("command", discord.RollName)

	return Handler{
		ctx:   ctx,
		log:   log,
		utils: discordUtils.New(log),
	}
}

func (h Handler) Handle(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	size := discord.RollOptionSizeDefault
	count := discord.RollOptionCountDefault

	optionsMap := discordUtils.OptionsMap(i)
	if opt, ok := optionsMap[discord.RollOptionSize]; ok {
		size = die.Size(opt.IntValue())
	}
	if opt, ok := optionsMap[discord.RollOptionCount]; ok {
		count = int(opt.IntValue())
	}

	results := make([]int, 0, count)
	for index := 0; index < count; index++ {
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

	_ = h.utils.Respond(h.ctx, s, i, response)
}
