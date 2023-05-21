package dxs

import (
	"encoding/json"
	"fmt"
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

func (e Entry) String() string {
	var value string
	valF, ok := e.Value.(float64)
	if ok {
		switch e.ID {
		case OpHours:
			value = fmt.Sprintf("%.0f h", valF)
		case TotalYield:
			value = fmt.Sprintf("%.3f KWh", valF)
		case DailyYield:
			value = fmt.Sprintf("%.3f Wh", valF)
		case AC_P:
			value = fmt.Sprintf("%.2f W", valF)
		default:
			value = fmt.Sprintf("%f", valF)
		}
	} else {
		value = fmt.Sprintf("%s", e.Value)
	}

	return fmt.Sprintf("%s: %s", e.ID, value)
}

type Response struct {
	DxsEntries []Entry
}

func PrintEntries(entries []Entry) {
	for _, entry := range entries {
		fmt.Printf("%s\n", entry)
	}
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
