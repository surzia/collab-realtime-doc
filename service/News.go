package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/surzia/grpc-starter/conf"
	"github.com/surzia/grpc-starter/protobuf"
)

// ApiResponse is struct of response of http api
type ApiResponse struct {
	Reason string `json:"reason"`
	Result Result `json:"result"`
}

type Result struct {
	Stat string `json:"stat"`
	Data []News `json:"data"`
}

// News is hot top new's struct
type News struct {
	Uniquekey     string `json:"uniquekey"`
	Title         string `json:"title"`
	Date          string `json:"date"`
	Category      string `json:"category"`
	AuthorName    string `json:"author_name"`
	Url           string `json:"url"`
	ThumbnailPicS string `json:"thumbnail_pic_s"`
	IsContent     string `json:"is_content"`
}

// NewService contains services that fetch new and convert to grpc protobuf
type NewService struct {
	apiUri   string
	apiKey   string
	reqType  string
	page     int
	size     int
	isFilter int
}

func NewNewService(conf *conf.Config) *NewService {
	return &NewService{
		apiUri:   conf.NewsConf.ApiUri,
		apiKey:   conf.NewsConf.ApiKey,
		reqType:  conf.NewsConf.Type,
		page:     conf.NewsConf.Page,
		size:     conf.NewsConf.Size,
		isFilter: conf.NewsConf.IsFilter,
	}
}

func (s *NewService) RequestPublishAPI() []*protobuf.New {
	reqUrl := fmt.Sprintf("%s?type=%s&page=%d&page_size=%d&is_filter=%d&key=%s", s.apiUri, s.reqType, s.page, s.size, s.isFilter, s.apiKey)
	log.Printf("request url: %s", reqUrl)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, reqUrl, nil)

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var resp ApiResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		panic(err)
	}

	var ret []*protobuf.New
	for _, n := range resp.Result.Data {
		isContent, _ := strconv.Atoi(n.IsContent)
		ret = append(ret, &protobuf.New{
			Uniquekey:     n.Uniquekey,
			Title:         n.Title,
			Date:          n.Date,
			Category:      n.Category,
			AuthorName:    n.AuthorName,
			Url:           n.Url,
			ThumbnailPicS: n.ThumbnailPicS,
			IsContent:     int64(isContent),
		})
	}

	return ret
}
