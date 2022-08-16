import (
    "fmt"
    "log"
    "net/http"
    "net/http/httptest"
    "net/http/httputil"
    "github.com/gorilla/mux"
)
:

func logHandler(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        x, err := httputil.DumpRequest(r, true)
        if err != nil {
            http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
            return
        }
        log.Println(fmt.Sprintf("%q", x))
        rec := httptest.NewRecorder()
        fn(rec, r)
        log.Println(fmt.Sprintf("%q", rec.Body))            
    }
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "A message was received")
}