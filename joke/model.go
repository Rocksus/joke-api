package joke

const (
	defaultLang = "en"
)

var jokeCategories = map[string]bool{
	"general":     true,
	"knock-knock": true,
	"programming": true,
}

type Joke struct {
	ID        int64  `json:"id"`
	Category  string `json:"category"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Language  string `json:"lang"`
}

type jokeResponse struct {
	ID        int64  `json:"id"`
	Category  string `json:"category"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Language  string `json:"lang"`
	Status    int    `json:"status"`
	Error     string `json:"Error"`
}
