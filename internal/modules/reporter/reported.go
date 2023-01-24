package reporter

import (
	"fmt"
	"log"
	"time"

	"github.com/edpryk/buter/internal/modules/requester"
)

type Reporter struct{}

type Filters interface {
	Url() bool
	Status() bool
	Length() bool
	Duration() bool
}

func (r Reporter) ListenReponse(responseQ chan requester.CustomResponse, filters Filters) {
	for res := range responseQ {
		report := fmt.Sprintf("%3s", "")

		duration := res.Duration
		// url := res.Request.URL
		code := res.StatusCode
		payloads := ""
		for number, p := range res.Payloads {
			payloads += fmt.Sprintf("P_%d: %-5s", number+1, p)
		}

		report += payloads

		if filters != nil {
		}

		report += fmt.Sprintf("Status %-5d", code)
		report += fmt.Sprintf("Duration %5dms", duration/time.Millisecond)

		log.Println(report)
	}
}

func New() Reporter {
	return Reporter{}
}
