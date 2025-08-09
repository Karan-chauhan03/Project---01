package main
import ("fmt"
        "log"
	"net/http"
)

func main(){
	fmt.Println("PROJECT-NAME = BLACKWALL")

	
	log.Println("Server started at http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

