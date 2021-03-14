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
	return NthBadAssPage(1)
}

func NthBadAssPage(n int) (*BadAss, error) {
	fmt.Printf("%d First Page\n", n)
	badAss := &BadAss{}
	url := fmt.Sprintf(baseURL, n)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(badAss)
	fmt.Printf("%d First Page\n", n)
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

func GetAllBadAssAsFunCSS() ([]*FunCSS, error) {
	funCSSES := make([]*FunCSS, 0, 0)
	page := 1
	badAss, err := FirstBadAssPage()
	if err != nil {
		return nil, err
	}
	to := 0
	total := badAss.Total
	for total != to {
		badAss, err = NthBadAssPage(page)
		if err != nil {
			return nil, err
		}
		page += 1
		to = badAss.To
		funCSSES = append(funCSSES, badAssAsFunCSS(badAss)...)
	}
	return funCSSES, nil
}
