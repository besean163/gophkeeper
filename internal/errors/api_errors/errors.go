package apierrors

type Error struct {
	Error ErrorData `json:"error"`
}

type ErrorData struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
