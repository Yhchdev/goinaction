package search

import (
	"encoding/json"
	"os"
)

//将数据从文件读入到特定的数据结构中

const datafile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(datafile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var feeds []*Feed

	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err

}
