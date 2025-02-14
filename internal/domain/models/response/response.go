package response

type (
	Error struct {
		Error string `json:"error"`
	} // @name Error

	Success struct {
		Data interface{} `json:"data"`
	} // @name Success
)
