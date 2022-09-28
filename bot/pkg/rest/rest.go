package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

func HeathCheck() bool {
	resp, err := http.Get("https://my-json-server.typicode.com/waldirborbajr/mockrest/profile")

	if err != nil {
		log.Info().Msg("Error calling api")
		return false
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)

	contentStr := string(content)
	fmt.Println("Content: \n", contentStr)

	return true
}
