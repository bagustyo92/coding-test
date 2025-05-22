package main

import (
	"encoding/json"
	_ "encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http"
	"net/url"
	"strings"
	"sync"
)

/*
 * Complete the 'bestUniversityByCountry' function below.
 * Base URL for copy/paste: https://jsonmock.hackerrank.com/api/universities
 *
 * The function is expected to return a STRING.
 * The function accepts STRING country as parameter.
 */

type DataResp struct {
	University string `json:"university"`
	Rank       string `json:"rank_display"`
	Type       string `json:"string"`
	Location   struct {
		Country string `json:"country"`
	} `location`
}

type Resp struct {
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
	Data       []DataResp `json:"data"`
}

func bestUniversityByCountry(country string) string {
	var wg sync.WaitGroup
	var mu sync.Mutex
	page := 1
	resp, err := invokeRequest(page)
	if err != nil {
		return ""
	}
	if resp == nil {
		return ""
	}
	var datas []DataResp
	datas = append(datas, resp.Data...)
	if resp.TotalPages > 1 {
		for page = 2; page <= resp.TotalPages; page++ {
			wg.Add(1)

			go func(p int) {
				defer wg.Done()

				resp, err := invokeRequest(page)
				if err != nil {
					return
				}

				mu.Lock()
				datas = append(datas, resp.Data...)
				mu.Unlock()
			}(page)

		}
		wg.Wait()

		// for page = 2; page <= resp.TotalPages; page++ {
		// 	resp, err := invokeRequest(page)
		// 	if err != nil {
		// 		return ""
		// 	}
		// 	datas = append(datas, resp.Data...)
		// }
	}

	bestUniversityOnEachCountry := make(map[string]DataResp)
	for _, data := range datas {
		currentData, exist := bestUniversityOnEachCountry[strings.ToLower(data.Location.Country)]
		if !exist || data.Rank < currentData.Rank {
			bestUniversityOnEachCountry[strings.ToLower(data.Location.Country)] = data
		}
	}

	best, ok := bestUniversityOnEachCountry[country]
	if !ok {
		return ""
	}
	return best.University
}

func invokeRequest(page int) (*Resp, error) {
	baseUrl := "https://jsonmock.hackerrank.com/api/universities"
	params := url.Values{}
	params.Add("page", fmt.Sprintf("%d", page))

	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Printf("Error when GET request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error when reading body resp: %v", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Printf("Error wiht status http code %d", resp.StatusCode)
		return nil, errors.New("Err with http code not 200")
	}

	response := Resp{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("Error while marshal the json %v", err)
		return nil, err
	}

	return &response, nil
}

func main() {
	fmt.Println(bestUniversityByCountry(strings.ToLower("UNITED STATES")))
}
