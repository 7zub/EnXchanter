package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type IParams interface {
	GetParams(ccy Ccy) *Request
}

type Request struct {
	Exchange int
	Url      string
	Params   IParams
	Response IResponse
	ReqDate  time.Time
}

func (r *Request) SendRequest() {
	r.UrlExec(r.UrlBuild())
}

func (r *Request) UrlBuild() *http.Request {
	fields := reflect.TypeOf(r.Params)
	values := reflect.ValueOf(r.Params)

	rq, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		panic(err)
	}

	q := rq.URL.Query()

	for i := 0; i < fields.NumField(); i++ {
		q.Add(strings.ToLower(fields.Field(i).Name), fmt.Sprintf("%v", values.Field(i)))
	}

	rq.URL.RawQuery = q.Encode()
	fmt.Printf("Полный URL: %s\n", rq.URL.String())

	return rq
}

func (r *Request) UrlExec(rq *http.Request) {
	r.ReqDate = time.Now()
	client := http.Client{}
	resp, err := client.Do(rq)
	if err != nil {
		log.Fatalln(err)
	}
	json.NewDecoder(resp.Body).Decode(r.Response)
}
