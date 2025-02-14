//go:build wireinject
// +build wireinject

package di

import (
	"eMobile/internal/api"
	"eMobile/internal/config"
	"eMobile/internal/db"
	songProvider "eMobile/internal/domain/providers/song"
	"github.com/google/wire"
	"log/slog"
)

func InitializeAPI(cfg *config.Config, log *slog.Logger) (*api.ServerHTTP, error) {
	panic(wire.Build(
		songProvider.SongProviderSet,

		db.ConnectToBD,
		api.NewServerHTTP,
	))
}
