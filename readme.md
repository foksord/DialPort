## PortChecker Go Library

# Using

func PortChecker() (int64, error) 


Example   
```go

package main

// [START import]
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/foksord/portChecker"
)

func main() {

	port, err := portChecker.PortChecker()
	if err != nil {
		panic(err.Error())
	}
	var sb strings.Builder
	sb.WriteString(":")
	sb.WriteString(strconv.FormatInt(int64(port), 10))
	log.Printf("Listening on port %s", sb.String())

	http.HandleFunc("/", indexHandler)

	log.Printf("Listening on port %s", sb.String())
	if err := http.ListenAndServe(sb.String(), nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(writeRes http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(writeRes, req)
		return
	}
	fmt.Fprint(writeRes, "Hello, World!")
}
```
