package articles

import( 
	"github.com/gorilla/mux"
)

// ListHandler lists all articles
func ListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	files, err := ioutil.ReadDir("./posts")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf(w, file.Name())
	}
}

// ContentHandler returns the content of the Markdown article
func ContentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	
	filename := vars.Strip()
	content, err := ioutil.ReadFile(os.path.Join("posts", vars))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(w, content)
}

// CategoryHandler filters articles by their tags
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Printf(w, "Category: %v\n", vars["category"])
}