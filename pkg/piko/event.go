package piko

import (
	"encoding/binary"
	"errors"
	"time"

	"github.com/tobima/piko2mqtt/pkg/dxs"
	"golang.org/x/exp/slices"
)

var LogEntriesIDs = []dxs.DxsID{
	dxs.LogEntry0,
	dxs.LogEntry1,
	dxs.LogEntry2,
	dxs.LogEntry3,
	dxs.LogEntry4,
	dxs.LogEntry5,
	dxs.LogEntry6,
	dxs.LogEntry7,
	dxs.LogEntry8,
	dxs.LogEntry9,
}

type Event struct {
	Date time.Time
	Code int16
	Env  int16
}

// inLocTime adjust a time zone (subtract the difference to utc)
func inLocTime(inLoc time.Time) time.Time {
	_, offset := inLoc.Zone()
	ts := inLoc.Add(time.Duration(-offset) * time.Second)
	return ts
}

func ParseEvents(resp dxs.Response) []Event {
	// TODO: add error handling
	res := resp.GetEntryValue(dxs.LogEntries)
	if res == nil {
		return nil
	}
	f, ok := res.(float64)
	if !ok {
		return nil
	}
	num := int(f)
	if num > 10 {
		return nil
	}
	if num < 1 {
		return nil
	}
	ids := LogEntriesIDs[0:num]
	events := make([]Event, num)
	for _, entry := range resp.DxsEntries {
		idx := slices.Index(ids, entry.ID)
		if (idx >= -1) && (idx <= num) {
			value, ok := entry.Value.([]interface{})
			if !ok {
				continue
			}
			event, err := parseEventEntry(value)
			if err != nil {
				continue
			}
			events[idx] = event
		}
	}

	return events
}

func parseEventEntry(data []interface{}) (Event, error) {
	var evt Event
	if len(data) != 8 {
		return evt, errors.New("invalid length  for event")
	}
	d, err := convData(data)
	if err != nil {
		return evt, err
	}
	evt.Date = inLocTime(time.Unix(int64(binary.LittleEndian.Uint32(d[0:4])), 0))
	evt.Code = int16(binary.LittleEndian.Uint16(d[4:6]))
	evt.Env = int16(binary.LittleEndian.Uint16(d[6:8]))
	return evt, nil
}

func convData(data []interface{}) ([]byte, error) {
	res := make([]byte, len(data))
	for idx, el := range data {
		val, ok := el.(float64)
		if !ok || val > 255 || val < 0 {
			return nil, errors.New("invalid input")
		}
		res[idx] = byte(int8(val))
	}
	return res, nil
}
