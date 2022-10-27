package piko

import (
	"fmt"

	"github.com/tobima/piko2mqtt/pkg/dxs"
)

type MPPT struct {
	Voltage float32
	Current float32
	Power   float32
}

func (m MPPT) String() string {
	return fmt.Sprintf("U: %6.2f V, I: %5.2f A, P: %7.2f W", m.Voltage, m.Current, m.Power)
}

func ParseMPPT(resp dxs.Response, no int) MPPT {
	var curID, volID, pID dxs.DxsID
	switch no {
	case 1:
		curID = dxs.DC1_I
		volID = dxs.DC1_U
		pID = dxs.DC1_P
	case 2:
		curID = dxs.DC2_I
		volID = dxs.DC2_U
		pID = dxs.DC2_P
	case 3:
		curID = dxs.DC3_I
		volID = dxs.DC3_U
		pID = dxs.DC3_P
	}
	var data MPPT
	for _, entry := range resp.DxsEntries {
		switch entry.ID {
		case curID:
			curVal, ok := entry.Value.(float64)
			if ok {
				data.Current = float32(curVal)
			}
		case volID:
			volVal, ok := entry.Value.(float64)
			if ok {
				data.Voltage = float32(volVal)
			}
		case pID:
			powerVal, ok := entry.Value.(float64)
			if ok {
				data.Power = float32(powerVal)
			}
		}
	}
	return data
}
