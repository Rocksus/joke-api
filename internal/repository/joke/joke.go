package joke

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// the joke repo will load all jokes from the json file to here
var jokes map[int64]*Joke

// Load is a function helper that loads the joke file to the repo
func Load(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	var jokeRaw []Joke

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &jokeRaw)

	jokes = make(map[int64]*Joke, len(jokeRaw))

	for _, v := range jokeRaw {
		jokes[v.ID] = &Joke{
			ID:        v.ID,
			Category:  v.Category,
			Setup:     v.Setup,
			Punchline: v.Punchline,
			Language:  v.Language,
		}
	}

	log.Print("[Joke] Successfully loaded the horrible jokes")
}

// InitHandler returns the joke handler
func InitHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseInt(vars["id"], 10, 64)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if jokes[id] == nil {
			raw := jokeResponse{
				Error: "No joke found from given ID.",
			}

			byteData, err := json.Marshal(raw)

			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, string(byteData))

		} else {
			data := jokes[id]
			raw := jokeResponse{
				ID:        data.ID,
				Category:  data.Category,
				Setup:     data.Setup,
				Punchline: data.Punchline,
				Language:  data.Language,
				Error:     "",
				Status:    http.StatusOK,
			}
			byteData, err := json.Marshal(raw)

			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, string(byteData))
		}

	}
}

func InitRandomHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jokeCategory := "all"
		vars := mux.Vars(r)
		if vars["category"] != "" {
			jokeCategory = vars["category"]
		}

		keys := r.URL.Query()
		lang := keys.Get("lang")
		if lang == "" {
			lang = defaultLang
		}

		// Create a priority list, if anyone knows a more
		// efficient way of doing this, please do make a PR
		pList := rand.Perm(len(jokes))

		found := false
		var jokeData *Joke
		for i := int(0); !found && i < len(pList); i++ {
			if jokes[int64(pList[i])].Category == jokeCategory || jokeCategory == "all" {
				if jokes[int64(pList[i])].Language == lang {
					found = true
					jokeData = jokes[int64(pList[i])]
				}
			}
		}

		raw := jokeResponse{
			ID:        jokeData.ID,
			Category:  jokeData.Category,
			Setup:     jokeData.Setup,
			Punchline: jokeData.Punchline,
			Language:  jokeData.Language,
			Error:     "",
			Status:    http.StatusOK,
		}
		byteData, err := json.Marshal(raw)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(byteData))
	}
}
