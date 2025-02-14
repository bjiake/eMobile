package song

type (
	Entity struct {
		ID int64 `json:"id"`
		Song
	} // @name Entity

	Song struct {
		Create
		Detail
	} // @name Song

	Create struct {
		// Название песни
		Name string `json:"name" validate:"required"`
		// Группа, к которой принадлежит песня
		Group string `json:"group" validate:"required"`
	} // @name Create

	Detail struct {
		// Дата релиза в формате yyyy.dd.mm
		ReleaseDate string `json:"releaseDate" validate:"omitempty" example:"2023.15.01"`
		// Текст песни *ПРИПЕВЫ выделять -Начало припева- .... -Конец припева-
		Text string `json:"text" validate:"omitempty" example:""`
		// Ссылка на песню
		Link string `json:"link" validate:"omitempty"`
	} // @name Detail

	Pagination struct {
		// Номер страницы
		Page int `json:"page" validate:"required"`
		// Количество элементов на странице
		PageSize int `json:"page_size" validate:"required"`
	} // @name Pagination

	Search struct {
		Pagination
		// Название песни для поиска
		Name string `json:"name" validate:"omitempty"`
		// Группа для поиска
		Group string `json:"group" validate:"omitempty"`
		Detail
	} // @name Search
)
