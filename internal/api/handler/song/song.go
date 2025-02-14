package song

import (
	"eMobile/internal/domain/interfaces"
	_ "eMobile/internal/domain/models/response"
	model "eMobile/internal/domain/models/song"
	"eMobile/pkg/logger/slogError"
	responseApi "eMobile/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	Svc interfaces.SongService
	Log *slog.Logger
}

func (this *Handler) NewSongHandler(r chi.Router) {
	r.Route("/song", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/", this.Create)
			r.Delete("/{id}", this.Delete)
			r.Put("/{id}", this.Update)
			r.Get("/{id}", this.Get)
			r.Get("/text/{id}", this.Text)
			r.Get("/search", this.Search)
		})
	})
}

// Create
// @Summary Create a new song
// @Description Create a new song in the system
// @Tags songs
// @Accept json
// @Produce json
// @Param song body model.Create true "Song data"
// @Success 201 {object} model.Song
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /song [post]
func (this *Handler) Create(w http.ResponseWriter, r *http.Request) {
	const op = "handler.song.Create"

	this.Log = this.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var song model.Create

	err := render.DecodeJSON(r.Body, &song)
	if err != nil {
		this.Log.Error("failed to decode form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	result, err := this.Svc.Create(r.Context(), song)
	if err != nil {
		this.Log.Error("failed to create song", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, result)
}

// Delete
// @Summary Delete a song by ID
// @Description Delete a song from the system by its ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 204
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /song/{id} [delete]
func (this *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	const op = "handler.song.Delete"

	this.Log = this.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		this.Log.Error("failed to parse id parameter", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = this.Svc.Delete(r.Context(), id)
	if err != nil {
		this.Log.Error("failed to delete song", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusNoContent, nil)
}

// Update
// @Summary Update a song by ID
// @Description Update an existing song in the system by its ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body model.Song true "Updated song data"
// @Success 200 {object} model.Song
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /song/{id} [put]
func (this *Handler) Update(w http.ResponseWriter, r *http.Request) {
	const op = "handler.song.Update"

	this.Log = this.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		this.Log.Error("failed to parse id parameter", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	var song model.Song
	err = render.DecodeJSON(r.Body, &song)
	if err != nil {
		this.Log.Error("failed to decode form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	update, err := this.Svc.Update(r.Context(), id, song)
	if err != nil {
		this.Log.Error("failed to update song", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, update)
}

func (this *Handler) Get(w http.ResponseWriter, r *http.Request) {
	const op = "handler.song.Get"

	this.Log = this.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		this.Log.Error("failed to parse id parameter", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	result, err := this.Svc.Get(r.Context(), id)
	if err != nil {
		this.Log.Error("failed to get song", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, result)
}

func (this *Handler) Text(w http.ResponseWriter, r *http.Request) {
	const op = "handler.song.Text"

	this.Log = this.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		this.Log.Error("failed to parse id parameter", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	var pagination model.Pagination
	err = render.DecodeJSON(r.Body, &pagination)
	if err != nil {
		this.Log.Error("failed to decode form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	result, err := this.Svc.GetText(r.Context(), id, pagination)
	if err != nil {
		this.Log.Error("failed to get song", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, result)
}

func (this *Handler) Search(w http.ResponseWriter, r *http.Request) {
	const op = "handler.song.Search"

	this.Log = this.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var search model.Search
	err := render.DecodeJSON(r.Body, &search)
	if err != nil {
		this.Log.Error("failed to decode form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	result, err := this.Svc.Search(r.Context(), search)
	if err != nil {
		this.Log.Error("failed to search song", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, result)
}
