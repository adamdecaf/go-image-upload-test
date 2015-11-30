package routes

import (
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func RandomImage(w http.ResponseWriter, r *http.Request) {
	files, err := filepath.Glob("./tmp/*")

	if err != nil {
		http.Error(w, "whoops", 500)
	}

	idx := rand.Intn(len(files))
	file := files[idx]

	if _, err := os.Stat(file); err != nil {
		http.Error(w, "Error reading file", 500)
	}

	fixed_filename := strings.Replace(file, "tmp", "i", -1)

	http.Redirect(w, r, fixed_filename, 302) // 302 Found
}
