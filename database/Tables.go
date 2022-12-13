package db

type WordsTable struct {
	Id        uint32 `json:"id"`
	Id_normal uint32 `json:"id_normal"`
	Word      string `json:"word"`
	Detail    string `json:"detail"`
}

type OtherTables struct {
	Id      uint32 `json:"id"`
	Id_word uint32 `json:"id_word"`
	Name    string `json:"name"`
}

var NameTables = map[string][]string{
	"words":        {"id", "id_normal", "word", "detail"},
	"parts_speech": {"id", "id_word", "name"},
	"grammar":      {"id", "id_word", "name"},
	"cases":        {"id", "id_word", "name"},
}

var TableIds = map[string]int{
	"words":        1,
	"parts_speech": 0,
	"grammar":      0,
	"cases":        0,
}
