package interfaces

import (
	"context"
	model "eMobile/internal/domain/models/song"
	"github.com/go-chi/chi/v5"
	"net/http"
)

//go:generate mockery --name SongRepo
type SongRepo interface {
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, song model.Create) (int64, error)
	Update(ctx context.Context, id int64, song model.Song) error
	Get(ctx context.Context, id int64) (model.Entity, error)
	GetText(ctx context.Context, id int64) (string, error)
	Search(ctx context.Context, search model.Search) ([]model.Entity, error)
}

//go:generate mockery --name SongService
type SongService interface {
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, song model.Create) (int64, error)
	Update(ctx context.Context, id int64, song model.Song) (model.Song, error)
	Get(ctx context.Context, id int64) (model.Entity, error)
	GetText(ctx context.Context, id int64, pagination model.Pagination) (string, error)
	Search(ctx context.Context, search model.Search) ([]model.Entity, error)
}

type SongHandler interface {
	NewSongHandler(r chi.Router)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Text(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}
