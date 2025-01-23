package listUserQuotes_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/listUserQuotes"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/services/listUserQuotes"
)

func TestService_Exec(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	i := &discordgo.Interaction{
		GuildID: "1",
	}
	author := &discordgo.Member{
		User: &discordgo.User{
			ID: "2",
		},
	}

	type args struct {
		page int
	}
	type setup func(t *testing.T, args args, err error) *listUserQuotes.Service

	tests := map[string]struct {
		args  args
		err   error
		setup setup
	}{
		"author_guild_member_error": {
			err: errors.New("guild member error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *listUserQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(nil, err)

				return listUserQuotes.New(d, db)
			},
		},
		"quote_count_db_error": {
			err: errors.New("db error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *listUserQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				db.EXPECT().CountUserQuotesInGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(0, err)

				return listUserQuotes.New(d, db)
			},
		},
		"empty_response_error": {
			err: errors.New("response error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *listUserQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				db.EXPECT().CountUserQuotesInGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(0, nil)

				d.EXPECT().Respond(
					ctx, i, responses.QuoteUserNoQuotes(author),
				).Return(err)

				return listUserQuotes.New(d, db)
			},
		},
		"quotes_page_db_error": {
			err: errors.New("db error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *listUserQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				count := 1
				db.EXPECT().CountUserQuotesInGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(count, nil)

				db.EXPECT().GetUserQuotesInGuild(
					ctx,
					models.ID(i.GuildID),
					models.ID(author.User.ID),
					0,
					commands.QuoteMaxQuotesPerPage,
				).Return(nil, err)

				return listUserQuotes.New(d, db)
			},
		},
		"quotes_page_response_error": {
			err: errors.New("response error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *listUserQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				count := 1
				db.EXPECT().CountUserQuotesInGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(count, nil)

				quotes := []*models.Quote{
					{AddedByID: "3"},
				}
				db.EXPECT().GetUserQuotesInGuild(
					ctx,
					models.ID(i.GuildID),
					models.ID(author.User.ID),
					0,
					commands.QuoteMaxQuotesPerPage,
				).Return(quotes, nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(err)

				return listUserQuotes.New(d, db)
			},
		},
		"success_empty": {
			err: nil,
			setup: func(
				t *testing.T, _ args, _ error,
			) *listUserQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				db.EXPECT().CountUserQuotesInGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(0, nil)

				d.EXPECT().Respond(
					ctx, i, responses.QuoteUserNoQuotes(author),
				).Return(nil)

				return listUserQuotes.New(d, db)
			},
		},
		"success_single_page": {
			err: nil,
			setup: func(
				t *testing.T, _ args, _ error,
			) *listUserQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				count := 1
				db.EXPECT().CountUserQuotesInGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(count, nil)

				quotes := []*models.Quote{
					{
						AuthorID:  models.ID(author.User.ID),
						Text:      "text",
						GuildID:   models.ID(i.GuildID),
						AddedByID: "3",
						Timestamp: time.Now().UTC(),
					},
				}
				db.EXPECT().GetUserQuotesInGuild(
					ctx,
					models.ID(i.GuildID),
					models.ID(author.User.ID),
					0,
					commands.QuoteMaxQuotesPerPage,
				).Return(quotes, nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).RunAndReturn(
					func(
						_ context.Context,
						_ *discordgo.Interaction,
						r *discordgo.InteractionResponse,
					) error {
						assert.Len(t, r.Data.Embeds, 2)
						assert.Equal(t, quotes[0].Text, r.Data.Embeds[1].Title)
						assert.Equal(
							t, quotes[0].Timestamp.Format(time.RFC3339),
							r.Data.Embeds[1].Timestamp,
						)

						return nil
					},
				)

				return listUserQuotes.New(d, db)
			},
		},
		"success_page_1": {
			args: args{
				page: 1,
			},
			err: nil,
			setup: func(
				t *testing.T, args args, _ error,
			) *listUserQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				count := 6
				db.EXPECT().CountUserQuotesInGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(count, nil)

				quotes := []*models.Quote{
					{
						AuthorID:  models.ID(author.User.ID),
						Text:      "text",
						GuildID:   models.ID(i.GuildID),
						AddedByID: "3",
						Timestamp: time.Now().UTC(),
					},
				}
				db.EXPECT().GetUserQuotesInGuild(
					ctx,
					models.ID(i.GuildID),
					models.ID(author.User.ID),
					args.page*commands.QuoteMaxQuotesPerPage,
					commands.QuoteMaxQuotesPerPage,
				).Return(quotes, nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).RunAndReturn(
					func(
						_ context.Context,
						_ *discordgo.Interaction,
						r *discordgo.InteractionResponse,
					) error {
						assert.Len(t, r.Data.Embeds, 2)
						assert.Equal(t, quotes[0].Text, r.Data.Embeds[1].Title)
						assert.Equal(
							t, quotes[0].Timestamp.Format(time.RFC3339),
							r.Data.Embeds[1].Timestamp,
						)
						assert.Equal(
							t, "quote_get_2_0",
							r.Data. //nolint:errcheck // panic is ok
								Components[0].(discordgo.ActionsRow).
								Components[0].(discordgo.Button).CustomID,
						)
						assert.True(
							t,
							r.Data. //nolint:errcheck // panic is ok
								Components[0].(discordgo.ActionsRow).
								Components[1].(discordgo.Button).Disabled,
						)

						return nil
					},
				)

				return listUserQuotes.New(d, db)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.args, tt.err)
				err := s.Exec(ctx, i, models.ID(author.User.ID), tt.args.page)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
