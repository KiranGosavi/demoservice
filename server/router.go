package server
import(
	"demoservice/server/handler"
	"github.com/gorilla/mux"

)
func loadRoutes()*mux.Router {
	root :=mux.NewRouter()
	root.StrictSlash(true)
	root.HandleFunc("/", handler.DefaultPathHandler)
	return root
}
