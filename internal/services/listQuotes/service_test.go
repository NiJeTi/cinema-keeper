package listQuotes_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/listQuotes"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/services/listQuotes"
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

	type setup func(t *testing.T, err error) *listQuotes.Service

	tests := map[string]struct {
		err   error
		setup setup
	}{
		"guild_member": {
			err: errors.New("guild member error"),
			setup: func(t *testing.T, err error) *listQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(nil, err)

				return listQuotes.New(d, db)
			},
		},
		"db_error": {
			err: errors.New("db error"),
			setup: func(t *testing.T, err error) *listQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				db.EXPECT().GetUserQuotesOnGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(nil, err)

				return listQuotes.New(d, db)
			},
		},
		"empty_response_error": {
			err: errors.New("response error"),
			setup: func(t *testing.T, err error) *listQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				quotes := make([]*models.Quote, 0)
				db.EXPECT().GetUserQuotesOnGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(quotes, nil)

				d.EXPECT().Respond(
					ctx, i, responses.QuotesEmpty(author),
				).Return(err)

				return listQuotes.New(d, db)
			},
		},
		"guild_member_error": {
			err: errors.New("guild member error"),
			setup: func(t *testing.T, err error) *listQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				quotes := []*models.Quote{
					{
						Author:  author,
						Text:    "text",
						GuildID: models.ID(i.GuildID),
						AddedBy: &discordgo.Member{
							User: &discordgo.User{ID: "3"},
						},
						Timestamp: time.Time{},
					},
				}
				db.EXPECT().GetUserQuotesOnGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(quotes, nil)

				d.EXPECT().GuildMember(
					ctx,
					models.ID(i.GuildID),
					models.ID(quotes[0].AddedBy.User.ID),
				).Return(nil, err)

				return listQuotes.New(d, db)
			},
		},
		"header_response_error": {
			err: errors.New("header error"),
			setup: func(t *testing.T, err error) *listQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				quotes := []*models.Quote{
					{
						Author:  author,
						Text:    "text",
						GuildID: models.ID(i.GuildID),
						AddedBy: &discordgo.Member{
							User:   &discordgo.User{ID: "3"},
							Nick:   "addedBy",
							Avatar: "avatar",
						},
						Timestamp: time.Time{},
					},
				}
				db.EXPECT().GetUserQuotesOnGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(quotes, nil)

				d.EXPECT().GuildMember(
					ctx,
					models.ID(i.GuildID),
					models.ID(quotes[0].AddedBy.User.ID),
				).Return(quotes[0].AddedBy, nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(err)

				return listQuotes.New(d, db)
			},
		},
		"send_embeds_error": {
			err: errors.New("page error"),
			setup: func(t *testing.T, err error) *listQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				quotes := []*models.Quote{
					{
						Author:  author,
						Text:    "text",
						GuildID: models.ID(i.GuildID),
						AddedBy: &discordgo.Member{
							User:   &discordgo.User{ID: "3"},
							Nick:   "addedBy",
							Avatar: "avatar",
						},
						Timestamp: time.Time{},
					},
				}
				db.EXPECT().GetUserQuotesOnGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(quotes, nil)

				d.EXPECT().GuildMember(
					ctx,
					models.ID(i.GuildID),
					models.ID(quotes[0].AddedBy.User.ID),
				).Return(quotes[0].AddedBy, nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(nil)

				d.EXPECT().SendEmbeds(
					ctx, models.ID(i.ChannelID), mock.Anything,
				).Return(err)

				return listQuotes.New(d, db)
			},
		},
		"success_empty": {
			err: nil,
			setup: func(t *testing.T, _ error) *listQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				quotes := make([]*models.Quote, 0)
				db.EXPECT().GetUserQuotesOnGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(quotes, nil)

				d.EXPECT().Respond(
					ctx, i, responses.QuotesEmpty(author),
				).Return(nil)

				return listQuotes.New(d, db)
			},
		},
		"success": {
			err: nil,
			setup: func(t *testing.T, _ error) *listQuotes.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(author, nil)

				quotes := []*models.Quote{
					{
						Author:  author,
						Text:    "text",
						GuildID: models.ID(i.GuildID),
						AddedBy: &discordgo.Member{
							User:   &discordgo.User{ID: "3"},
							Nick:   "addedBy",
							Avatar: "avatar",
						},
						Timestamp: time.Time{},
					},
				}
				db.EXPECT().GetUserQuotesOnGuild(
					ctx, models.ID(i.GuildID), models.ID(author.User.ID),
				).Return(quotes, nil)

				d.EXPECT().GuildMember(
					ctx,
					models.ID(i.GuildID),
					models.ID(quotes[0].AddedBy.User.ID),
				).Return(quotes[0].AddedBy, nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).RunAndReturn(
					func(
						_ context.Context,
						_ *discordgo.Interaction,
						r *discordgo.InteractionResponse,
					) error {
						assert.Equal(t, "Quotes", r.Data.Embeds[0].Title)
						return nil
					},
				)

				d.EXPECT().SendEmbeds(
					ctx, models.ID(i.ChannelID), mock.Anything,
				).RunAndReturn(
					func(
						_ context.Context,
						_ models.ID,
						embeds []*discordgo.MessageEmbed,
					) error {
						assert.Len(t, embeds, 1)
						assert.Equal(t, quotes[0].Text, embeds[0].Title)
						assert.Equal(
							t,
							quotes[0].Timestamp.Format(time.RFC3339),
							embeds[0].Timestamp,
						)
						assert.Equal(
							t,
							quotes[0].AddedBy.DisplayName(),
							embeds[0].Footer.Text,
						)
						return nil
					},
				)

				return listQuotes.New(d, db)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.err)
				err := s.Exec(ctx, i, models.ID(author.User.ID))
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
