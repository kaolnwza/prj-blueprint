package api

import (
	"fmt"
	"net/url"
	"path"
	"reflect"
	"strings"
)

type ApiUrl struct {
	BaseUrl  string
	Endpoint string
	Param    []ApiQuery
	Query    []ApiQuery
}

type ApiQuery struct {
	Key   string
	Value string
}

func (h httpClient) BuildEndpoint(apiUrl ApiUrl) (*url.URL, error) {
	generatedUrl, err := url.Parse(apiUrl.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse baseUrl, err = %v", err)
	}

	if len(apiUrl.Param) == 0 {
		generatedUrl.Path = path.Join(generatedUrl.Path, apiUrl.Endpoint)
	} else {
		pathWithParam := apiUrl.Endpoint
		for _, p := range apiUrl.Param {
			pathWithParam = strings.Replace(pathWithParam, fmt.Sprint(":", p.Key), p.Value, 1)
		}
		generatedUrl.Path = path.Join(generatedUrl.Path, pathWithParam)
	}

	query := generatedUrl.Query()
	for _, q := range apiUrl.Query {
		if q.Key == "" || q.Value == "" {
			continue
		}
		query.Set(q.Key, q.Value)
	}

	generatedUrl.RawQuery = query.Encode()
	return generatedUrl, err
}

func NewApiQuery(i interface{}) []ApiQuery {
	var queries []ApiQuery
	val := reflect.ValueOf(i)
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		jsonField := t.Field(i).Tag.Get("json")
		field := val.Field(i)

		queries = append(queries, ApiQuery{
			Key:   jsonField,
			Value: fmt.Sprint(field.Interface()),
		})
	}
	return queries
}
