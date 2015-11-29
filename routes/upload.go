package routes

import (
	"encoding/hex"
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		// Create ./tmp if it doesn't exist
		if _, err := os.Stat("./tmp"); err != nil {
			if os.IsNotExist(err) {
				os.Mkdir("./tmp", 0777)
			}
		}

		hashed := SHA1(handler.Filename)
		hashed_base := hashed[0:8]

		extensions := strings.Split(handler.Filename, ".")
		extension := extensions[len(extensions)-1]

		hashed_filename := hashed_base + "." + extension

		f, err := os.OpenFile("./tmp/" + hashed_filename, os.O_WRONLY | os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}

		if (!is_valid_image(f)) {
			http.Error(w, "Invalid image", 500)
		} else {
			// fmt.Fprintf(w, "%v", handler.Header)
			http.Redirect(w, r, "/i/" + hashed_filename, 302) // 302 Found
		}

		defer f.Close()
		io.Copy(f, file)
	}
}

func SHA1(s string) (string) {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func is_valid_image(file *os.File) (bool) {
	return true
}
