package song

import (
	"database/sql"
	songHdl "eMobile/internal/api/handler/song"
	"eMobile/internal/domain/interfaces"
	songRepo "eMobile/internal/repo/song"
	songSvc "eMobile/internal/service/song"
	"github.com/google/wire"
	"log/slog"
	"sync"
)

var (
	hdl     *songHdl.Handler
	hdlOnce sync.Once

	svc     *songSvc.Service
	svcOnce sync.Once

	repository     *songRepo.Repo
	repositoryOnce sync.Once
)

var SongProviderSet = wire.NewSet(
	ProvideSongHandler,
	ProvideSongService,
	ProvideSongRepository,

	wire.Bind(new(interfaces.SongHandler), new(*songHdl.Handler)),
	wire.Bind(new(interfaces.SongService), new(*songSvc.Service)),
	wire.Bind(new(interfaces.SongRepo), new(*songRepo.Repo)),
)

func ProvideSongHandler(svc interfaces.SongService, log *slog.Logger) *songHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &songHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})
	return hdl
}

func ProvideSongService(repo interfaces.SongRepo) *songSvc.Service {
	svcOnce.Do(func() {
		svc = &songSvc.Service{
			Repo: repo,
		}
	})

	return svc
}

func ProvideSongRepository(db *sql.DB) *songRepo.Repo {
	repositoryOnce.Do(func() {
		repository = &songRepo.Repo{
			DB: db,
		}
	})

	return repository
}
