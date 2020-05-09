package articles

import( 
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
	"path"
	"io/ioutil"
)

// ListHandler lists all articles
func ListHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	files, err := ioutil.ReadDir("./posts")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Fprintf(w, file.Name())
	}
}

// ContentHandler returns the content of the Markdown article
func ContentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	
	content, err := ioutil.ReadFile(path.Join("posts", vars["filename"]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(content[:]))
}

// CategoryHandler filters articles by their tags
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}