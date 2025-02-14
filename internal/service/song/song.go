package song

import (
	"context"
	"eMobile/internal/domain/interfaces"
	model "eMobile/internal/domain/models/song"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Service struct {
	Repo interfaces.SongRepo
}

func (this *Service) Delete(ctx context.Context, id int64) error {
	const op = "service.song.Delete"

	err := this.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (this *Service) Create(ctx context.Context, song model.Create) (int64, error) {
	const op = "service.song.Create"

	res, err := this.Repo.Create(ctx, song)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (this *Service) Update(ctx context.Context, id int64, song model.Song) (model.Song, error) {
	const op = "service.song.Update"

	if song.ReleaseDate != "" {
		_, err := time.Parse("2006.01.02", song.ReleaseDate)
		if err != nil {
			return model.Song{}, fmt.Errorf("%s: invalid release date format, expected dd.mm.yyyy: %w", op, err)
		}
	}

	err := this.Repo.Update(ctx, id, song)
	if err != nil {
		return model.Song{}, err
	}

	return song, nil
}

func (this *Service) Get(ctx context.Context, id int64) (model.Entity, error) {
	const op = "service.song.Get"

	song, err := this.Repo.Get(ctx, id)
	if err != nil {
		return model.Entity{}, err
	}
	return song, nil
}

func (this *Service) GetText(ctx context.Context, id int64, pagination model.Pagination) (string, error) {
	const op = "service.song.GetText"

	text, err := this.Repo.GetText(ctx, id)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`(?s)-Начало припева-(.*?)-Конец припева-`)
	cleanedText := re.ReplaceAllString(text, "")

	lines := strings.Split(cleanedText, "\n")

	start := (pagination.Page - 1) * pagination.PageSize
	end := start + pagination.PageSize
	if start >= len(lines) {
		return "", fmt.Errorf("%s: Out of bounds: %w", op, err) // Возвращаем пустую строку, если страница вне диапазона
	}
	if end > len(lines) {
		end = len(lines) // Ограничиваем конец до длины массива
	}

	paginatedLines := lines[start:end] // Получаем нужные строки

	paginatedText := strings.Join(paginatedLines, "\n")

	return paginatedText, nil
}

func (this *Service) Search(ctx context.Context, search model.Search) ([]model.Entity, error) {
	const op = "service.song.Search"

	songs, err := this.Repo.Search(ctx, search)
	if err != nil {
		return nil, err
	}

	start := (search.Page - 1) * search.PageSize
	end := start + search.PageSize
	if start >= len(songs) {
		return nil, nil // Возвращаем пустую строку, если страница вне диапазона
	}
	if end > len(songs) {
		end = len(songs) // Ограничиваем конец до длины массива
	}

	return songs[start:end], nil
}
