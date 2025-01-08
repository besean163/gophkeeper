package entities

// NoteCreateInput структура для создания заметки.
type NoteCreateInput struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
