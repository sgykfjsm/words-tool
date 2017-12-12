package words

var (
	words *WordsResponse
)

// https://www.wordsapi.com/docs
type WordsResponse struct {
	Word          string        `json:"word"`
	Results       []Result      `json:"results"`
	Syllables     Syllable      `json:"syllables"`
	Pronunciation Pronunciation `json:"pronunciation"`
	Frequency     float64       `json:"frequency"`
}

type Result struct {
	Definition   string   `json:"definition"`
	PartOfSpeech string   `json:"partOfSpeech"`
	Synonyms     []string `json:"synonyms,omitempty"`
	Examples     []string `json:"examples,omitempty"`
	Derivation   []string `json:"derivation,omitempty"`
	SimilarTo    []string `json:"similarTo,omitempty"`
	TypeOf       []string `json:"typeOf,omitempty"`
	HasTypes     []string `json:"hasTypes,omitempty"`
	VerbGroup    []string `json:"verbGroup,omitempty"`
	Entails      []string `json:"entails,omitempty"`
	Also         []string `json:"also,omitempty"`
	Attribute    []string `json:"attribute,omitempty"`
	Antonyms     []string `json:"antonyms,omitempty"`
	InCategory   []string `json:"inCategory,omitempty"`
}

type Syllable struct {
	Count int      `json:"count"`
	List  []string `json:"list"`
}

type Pronunciation struct {
	All string `json:"all"`
}
