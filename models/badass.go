package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const baseURL = "http://bada55.io/lazyload?page=%d"

type BadAss struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	From        int `json:"from"`
	To          int `json:"to"`
	Data        []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Label     string `json:"label"`
		Author    string `json:"author"`
		Twitter   string `json:"twitter"`
		Active    string `json:"active"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Lumen     string `json:"lumen"`
		Token     string `json:"token"`
	} `json:"data"`
}

// URL to use http://bada55.io/lazyload?page=?
func FirstBadAssPage() (*BadAss, error) {
	fmt.Println("Start First Page")
	badAss := &BadAss{}
	url := fmt.Sprintf(baseURL, 1)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(badAss)
	fmt.Println("End First Page")
	return badAss, err
}

func badAssAsFunCSS(badAss *BadAss) []*FunCSS {
	funCSSList := make([]*FunCSS, 0, 0)
	for _, value := range badAss.Data {
		funCSS := &FunCSS{}
		// I should do something smart but I am okay with 'id' being zero
		id, _ := strconv.Atoi(value.ID)
		funCSS.ID = id
		funCSS.Author = value.Author
		funCSS.CSSHex = value.Name
		funCSS.Name = value.Label
		funCSSList = append(funCSSList, funCSS)
	}

	return funCSSList
}
