package response

type (
	Error struct {
		Error string `json:"error"`
	} // @name ResponseError

	Success struct {
		Data interface{} `json:"data"`
	} // @name ResponseSuccess
)
