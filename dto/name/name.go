package name

// Type Тип/Часть ФИО
type Type string

func (t Type) Values() []Type {
	return []Type{TypeUnknown, TypeSurname, TypeName, TypePatronymic}
}

// Gender Пол
type Gender string

func (g Gender) Values() []Gender {
	return []Gender{GenderUnknown, GenderMale, GenderFemale}
}

const (
	// TypeUnknown Неизвестно
	TypeUnknown Type = "UNKNOWN"
	// TypeSurname Фамилия
	TypeSurname Type = "SURNAME"
	// TypeName мя
	TypeName Type = "NAME"
	// TypePatronymic Отчество
	TypePatronymic Type = "NAME"

	// GenderUnknown Неизвестно
	GenderUnknown = "UNKNOWN"
	// GenderMale Мужской пол
	GenderMale = "MALE"
	// GenderFemale Женский пол
	GenderFemale = "FEMALE"
)

type CleanRequest struct {
	Query string `json:"query"`
}

type CleanResponse struct {
	Gender        Gender   `json:"gender"`
	Original      string   `json:"original"`
	Result        string   `json:"result"`
	LastName      string   `json:"lastName"`
	FirstName     string   `json:"firstName"`
	MiddleName    string   `json:"middleName"`
	UnparsedParts []string `json:"unparsedParts"`
	Possible      bool     `json:"possible"`
	Valid         bool     `json:"valid"`
}

type SuggestRequest struct {
	Type  Type   `json:"type"`
	Query string `json:"query"`
	Count int    `json:"count"`
}

type SuggestResponse struct {
	Suggestions []*Suggest `json:"suggestions"`
}

type Suggest struct {
	Result string `json:"result"`
	Lang   string `json:"lang"`
	Gender string `json:"gender"`
}
