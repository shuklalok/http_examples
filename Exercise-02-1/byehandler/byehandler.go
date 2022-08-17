package byehandler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (b *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Bye Handler")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "an error occurred", http.StatusBadRequest)
	}
	rw.Write([]byte(fmt.Sprintf("Bye %s", data)))
}
