package input

// NoteCreate структура для создания заметки.
type NoteCreate struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
