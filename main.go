/* 1
package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", snowSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
*/

/* 2
package main2

import "net/http"

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

 */

/* 3
package main2

import "net/http"

func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
*/

/* 4
package main2

import "net/http"

func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

 */

/* 5
package main2

import (
	"fmt"
	"net/http"
	"strconv"
)

func showSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w,"Display a specific snippet with ID %d....")
}

 */


/* 6
package main2

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}
func showSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w,"Display a specific snippet with ID %d...")
}

func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Allow",http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

 */

/* 7
package main2

import (
	"fmt"
	"html/template" // New import
	"log" // New import
	"net/http"
	"strconv"
)
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

*/

/* 8
package main2
import (
	"flag"
	"log"
	"net/http"
)
func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}


 */


