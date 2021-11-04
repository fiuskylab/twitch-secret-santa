package internal

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Common struct {
	Env
	L *zap.Logger
}

type Env struct {
	BOT_AUTH       string
	BOT_USERNAME   string
	TWITCH_CHANNEL string
}

func NewCommon(l *zap.Logger) *Common {
	if err := godotenv.Load(); err != nil {
		l.Error(err.Error())
		return &Common{}
	}

	return &Common{
		L: l,
		Env: Env{
			BOT_AUTH:       os.Getenv("BOT_AUTH"),
			BOT_USERNAME:   os.Getenv("BOT_USERNAME"),
			TWITCH_CHANNEL: os.Getenv("TWITCH_CHANNEL"),
		},
	}
}
