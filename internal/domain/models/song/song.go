package song

type (
	Entity struct {
		ID int64 `json:"id"`
		Song
	}

	Song struct {
		Create
		Detail
	}

	Create struct {
		Name  string `json:"name" validate:"required"`
		Group string `json:"group" validate:"required"`
	}

	Detail struct {
		ReleaseDate string `json:"releaseDate" validate:"omitempty"`
		Text        string `json:"text" validate:"omitempty"`
		Link        string `json:"link" validate:"omitempty"`
	}

	Pagination struct {
		Page     int `json:"page" validate:"required"`
		PageSize int `json:"page_size" validate:"required"`
	}

	Search struct {
		Pagination
		Name  string `json:"name" validate:"omitempty"`
		Group string `json:"group" validate:"omitempty"`
		Detail
	}
)
