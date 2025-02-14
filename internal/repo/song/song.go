package song

import (
	"context"
	"database/sql"
	model "eMobile/internal/domain/models/song"
	"fmt"
)

type Repo struct {
	DB *sql.DB
}

func (this *Repo) Delete(ctx context.Context, id int64) error {
	const op = "repo.song.Delete"

	query := "DELETE FROM song WHERE id = $1"

	stmt, err := this.DB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: preparing statement: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: executing statement: %w", op, err)
	}

	return nil
}

func (this *Repo) Create(ctx context.Context, song model.Create) (int64, error) {
	const op = "repo.song.Create"

	query := `
		INSERT INTO song (name, "group")
		VALUES ($1, $2)
		RETURNING id
	`

	stmt, err := this.DB.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("%s: preparing statement: %w", op, err)
	}
	defer stmt.Close()

	var id int64

	err = stmt.QueryRowContext(ctx, song.Name, song.Group).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: executing query: %w", op, err)
	}

	return id, nil
}

func (this *Repo) Update(ctx context.Context, id int64, song model.Song) error {
	const op = "repo.song.Update"

	query := `
		UPDATE song
		SET name = $1, "group" = $2, release_date = $3, text = $4, link = $5
		WHERE id = $6
	`

	stmt, err := this.DB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: preparing statement: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, song.Name, song.Group, song.Detail.ReleaseDate, song.Detail.Text, song.Detail.Link, id)
	if err != nil {
		return fmt.Errorf("%s: executing statement: %w", op, err)
	}

	return nil
}

func (this *Repo) Get(ctx context.Context, id int64) (model.Entity, error) {
	const op = "repo.song.Get"

	query := `
		SELECT id, name, "group", TO_CHAR(release_date, 'YYYY.DD.MM') AS release_date, text, link
		FROM song
		WHERE id = $1
	`

	stmt, err := this.DB.PrepareContext(ctx, query)
	if err != nil {
		return model.Entity{}, fmt.Errorf("%s: preparing statement: %w", op, err)
	}
	defer stmt.Close()

	var song model.Entity
	err = stmt.QueryRowContext(ctx, id).Scan(&song.ID, &song.Name, &song.Group, &song.ReleaseDate, &song.Text, &song.Link)
	if err != nil {
		return model.Entity{}, fmt.Errorf("%s: executing query: %w", op, err)
	}

	return song, nil
}

func (this *Repo) GetText(ctx context.Context, id int64) (string, error) {
	const op = "repo.song.GetText"

	query := `
		SELECT text
		FROM song
		WHERE id = $1
	`

	stmt, err := this.DB.PrepareContext(ctx, query)
	if err != nil {
		return "", fmt.Errorf("%s: preparing statement: %w", op, err)
	}
	defer stmt.Close()

	var text string
	err = stmt.QueryRowContext(ctx, id).Scan(&text)
	if err != nil {
		return "", fmt.Errorf("%s: executing query: %w", op, err)
	}

	return text, nil
}

func (this *Repo) Search(ctx context.Context, search model.Search) ([]model.Entity, error) {
	const op = "repo.song.Search"

	query := `
	SELECT id, name, "group", TO_CHAR(release_date, 'YYYY.DD.MM') AS release_date, text, link
	FROM song
	WHERE 1=1` // Используем 1=1 для упрощения добавления условий

	var args []interface{}
	argIndex := 1 // Индекс для параметров запроса

	if search.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		args = append(args, "%"+search.Name+"%")
		argIndex++
	}
	if search.Group != "" {
		query += fmt.Sprintf(" AND \"group\" ILIKE $%d", argIndex)
		args = append(args, "%"+search.Group+"%")
		argIndex++
	}
	if search.ReleaseDate != "" {
		query += fmt.Sprintf(" AND release_date::text ILIKE $%d", argIndex)
		args = append(args, "%"+search.ReleaseDate+"%")
		argIndex++
	}
	if search.Text != "" {
		query += fmt.Sprintf(" AND text ILIKE $%d", argIndex)
		args = append(args, "%"+search.Text+"%")
		argIndex++
	}
	if search.Link != "" {
		query += fmt.Sprintf(" AND link ILIKE $%d", argIndex)
		args = append(args, "%"+search.Link+"%")
		argIndex++
	}

	stmt, err := this.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: preparing statement: %w", op, err)
	}
	defer stmt.Close()

	var songs []model.Entity
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: executing query: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var song model.Entity
		if err := rows.Scan(&song.ID, &song.Name, &song.Group, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			return nil, fmt.Errorf("%s: scanning row: %w", op, err)
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: iterating rows: %w", op, err)
	}

	return songs, nil
}
