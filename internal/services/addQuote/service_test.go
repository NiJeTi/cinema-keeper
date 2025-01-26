package addQuote_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/addQuote"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/services/addQuote"
)

func TestService_Exec(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	i := &discordgo.Interaction{
		GuildID: "1",
		Member: &discordgo.Member{
			User: &discordgo.User{
				ID: "2",
			},
			Nick:   "addedBy",
			Avatar: "avatar",
		},
	}
	author := &discordgo.Member{
		User: &discordgo.User{
			ID: "1",
		},
		Nick:   "author",
		Avatar: "avatar",
	}
	const text = "text"

	type setup func(t *testing.T, err error) *addQuote.Service

	tests := map[string]struct {
		err   error
		setup setup
	}{
		"guild_member_error": {
			err: errors.New("guild member error"),
			setup: func(t *testing.T, err error) *addQuote.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(author.User.ID),
				).Return(nil, err)

				return addQuote.New(d, db)
			},
		},
		"db_error": {
			err: errors.New("db error"),
			setup: func(t *testing.T, err error) *addQuote.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(author.User.ID),
				).Return(author, nil)

				db.EXPECT().AddUserQuoteInGuild(ctx, mock.Anything).Return(err)

				return addQuote.New(d, db)
			},
		},
		"respond_error": {
			err: errors.New("respond error"),
			setup: func(t *testing.T, err error) *addQuote.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(author.User.ID),
				).Return(author, nil)

				db.EXPECT().AddUserQuoteInGuild(ctx, mock.Anything).Return(nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(err)

				return addQuote.New(d, db)
			},
		},
		"success": {
			err: nil,
			setup: func(t *testing.T, _ error) *addQuote.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().GuildMember(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(author.User.ID),
				).Return(author, nil)

				db.EXPECT().AddUserQuoteInGuild(
					ctx, mock.Anything,
				).RunAndReturn(
					func(_ context.Context, quote *models.Quote) error {
						assert.Equal(t, author.User.ID, string(quote.AuthorID))
						assert.Equal(t, text, quote.Text)
						assert.Equal(t, i.GuildID, string(quote.GuildID))
						assert.Equal(
							t, i.Member.User.ID, string(quote.AddedByID),
						)

						return nil
					},
				)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(nil)

				return addQuote.New(d, db)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.err)
				err := s.Exec(ctx, i, models.DiscordID(author.User.ID), text)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
