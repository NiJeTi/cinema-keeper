package discord

type Config struct {
	Token   string `conf:"token"`
	GuildID string `conf:"guild_id"`
}
