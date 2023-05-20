package dxs

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type DxsID int

func (i DxsID) Val() string {
	return strconv.Itoa(int(i))
}

type Entry struct {
	ID    DxsID `json:"dxsId"`
	Value interface{}
}

type Response struct {
	DxsEntries []Entry
}

func Gather(host string, entries []DxsID) Response {
	base, err := url.Parse("http://" + host + "/api/dxs.json")
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	for _, entry := range entries {
		params.Add("dxsEntries", entry.Val())
	}
	base.RawQuery = params.Encode()
	resp, err1 := http.Get(base.String())
	if err1 != nil {
		panic(err1)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	//fmt.Printf("Url: %s \n", base)
	//fmt.Printf("Resp: %s \n", body)
	return ParseRespone(body)
}

func ParseRespone(bar []byte) Response {
	var dat Response
	if err := json.Unmarshal(bar, &dat); err != nil {
		panic(err)
	}
	return dat
}

func (r Response) GetEntryValue(id DxsID) (value interface{}) {
	for _, entry := range r.DxsEntries {
		if entry.ID == id {
			value = entry.Value
			break
		}
	}
	return value
}
