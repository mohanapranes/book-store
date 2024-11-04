package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(req *http.Request, inter interface{}) {
	log.Printf("Parse body")
	body, err := io.ReadAll(req.Body)
	if err == nil {
		if err := json.Unmarshal([]byte(body), inter); err != nil {
			return
		}
	}
}
