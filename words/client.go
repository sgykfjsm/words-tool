package words

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

const (
	endpoint = "https://wordsapiv1.p.mashape.com/words/"
)

type Words struct {
	once       sync.Once
	Client     *http.Client
	MashapeKey string
	Endpoint   string
}

// Words
// Original API Doc: https://www.wordsapi.com/docs#words
// Mashape API Doc: https://market.mashape.com/wordsapi/wordsapi#word
func (w *Words) Words(word string) (*WordsResponse, error) {
	url := fmt.Sprintf("%s%s", w.Endpoint, word)
	res, err := w.get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get the result from ", url)
	}

	var wr WordsResponse
	if err := json.Unmarshal(res, &wr); err != nil {
		return nil, errors.Wrap(err, "failed to convert the response to json")
	}

	return &wr, nil
}

func (w *Words) get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create new request")
	}
	req.Header.Set("X-Mashape-Key", w.MashapeKey)
	req.Header.Set("Accept", "application/json")

	w.Client = new(http.Client)
	res, err := w.Client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get the response from %s", w.Endpoint)
	}

	if !strings.Contains(res.Header.Get("Content-Type"), "application/json") {
		return nil, errors.New(fmt.Sprint("unexpected content-type: ", res.Header.Get("Content-Type")))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read the response from %s", w.Endpoint)
	}
	defer res.Body.Close()

	return body, nil
}

func New(key string) *Words {
	return &Words{
		MashapeKey: key,
		Endpoint:   endpoint,
	}
}
