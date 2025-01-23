package printRandomQuote_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/printRandomQuote"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/services/printRandomQuote"
)

func TestService_Exec(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	i := &discordgo.Interaction{
		GuildID: "1",
	}

	type setup func(t *testing.T, err error) *printRandomQuote.Service

	tests := map[string]struct {
		err   error
		setup setup
	}{
		"db_error": {
			err: errors.New("db error"),
			setup: func(t *testing.T, err error) *printRandomQuote.Service {
				db := mocks.NewMockDb(t)
				d := mocks.NewMockDiscord(t)

				db.EXPECT().GetRandomQuoteInGuild(
					ctx, models.ID(i.GuildID),
				).Return(nil, err)

				return printRandomQuote.New(db, d)
			},
		},
		"no_quotes_response_error": {
			err: errors.New("response error"),
			setup: func(t *testing.T, err error) *printRandomQuote.Service {
				db := mocks.NewMockDb(t)
				d := mocks.NewMockDiscord(t)

				db.EXPECT().GetRandomQuoteInGuild(
					ctx, models.ID(i.GuildID),
				).Return(nil, nil)

				d.EXPECT().Respond(
					ctx, i, responses.QuoteGuildNoQuotes(),
				).Return(err)

				return printRandomQuote.New(db, d)
			},
		},
		"guild_member_error": {
			err: errors.New("guild member error"),
			setup: func(t *testing.T, err error) *printRandomQuote.Service {
				db := mocks.NewMockDb(t)
				d := mocks.NewMockDiscord(t)

				q := &models.Quote{
					AuthorID: "1",
				}
				db.EXPECT().GetRandomQuoteInGuild(
					ctx, models.ID(i.GuildID),
				).Return(q, nil)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), q.AuthorID,
				).Return(nil, err)

				return printRandomQuote.New(db, d)
			},
		},
		"response_error": {
			err: errors.New("response error"),
			setup: func(t *testing.T, err error) *printRandomQuote.Service {
				db := mocks.NewMockDb(t)
				d := mocks.NewMockDiscord(t)

				quote := &models.Quote{
					AuthorID: "1",
				}
				db.EXPECT().GetRandomQuoteInGuild(
					ctx, models.ID(i.GuildID),
				).Return(quote, nil)

				author := &discordgo.Member{
					User: &discordgo.User{
						ID: "1",
					},
					Nick: "author",
				}
				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), quote.AuthorID,
				).Return(author, nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(err)

				return printRandomQuote.New(db, d)
			},
		},
		"success_no_quotes": {
			err: nil,
			setup: func(t *testing.T, _ error) *printRandomQuote.Service {
				db := mocks.NewMockDb(t)
				d := mocks.NewMockDiscord(t)

				db.EXPECT().GetRandomQuoteInGuild(
					ctx, models.ID(i.GuildID),
				).Return(nil, nil)

				d.EXPECT().Respond(
					ctx, i, responses.QuoteGuildNoQuotes(),
				).Return(nil)

				return printRandomQuote.New(db, d)
			},
		},
		"success": {
			err: nil,
			setup: func(t *testing.T, _ error) *printRandomQuote.Service {
				db := mocks.NewMockDb(t)
				d := mocks.NewMockDiscord(t)

				quote := &models.Quote{
					AuthorID:  "1",
					Text:      "text",
					Timestamp: time.Now().UTC(),
				}
				db.EXPECT().GetRandomQuoteInGuild(
					ctx, models.ID(i.GuildID),
				).Return(quote, nil)

				author := &discordgo.Member{
					User: &discordgo.User{
						ID: "1",
					},
					Nick: "author",
				}
				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), quote.AuthorID,
				).Return(author, nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).RunAndReturn(
					func(
						_ context.Context,
						_ *discordgo.Interaction,
						r *discordgo.InteractionResponse,
					) error {
						assert.Equal(t, quote.Text, r.Data.Embeds[0].Title)
						assert.Equal(
							t, quote.Timestamp.Format(time.RFC3339),
							r.Data.Embeds[0].Timestamp,
						)
						assert.Equal(
							t, author.DisplayName(),
							r.Data.Embeds[0].Footer.Text,
						)

						return nil
					},
				)

				return printRandomQuote.New(db, d)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.err)
				err := s.Exec(ctx, i)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
