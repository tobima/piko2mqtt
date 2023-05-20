package piko

import "github.com/tobima/piko2mqtt/pkg/dxs"

type inverter struct {
	host string
}

func NewInverter(host string) inverter {
	return inverter{
		host: host,
	}
}

func (inv inverter) GetTypePlate() dxs.Response {
	entries := []dxs.DxsID{
		dxs.InvName,
		dxs.TypeName,
		dxs.Serial,
		dxs.ArticleNo,
		dxs.Mac,
	}

	data := dxs.Gather(inv.host, entries)
	return data
}

func (inv inverter) GatherAC() dxs.Response {
	entries := []dxs.DxsID{
		dxs.TotalYield,
		dxs.DailyYield,
		dxs.AC_P,
		dxs.OpHours,
	}

	data := dxs.Gather(inv.host, entries)
	return data
}

func (inv inverter) GatherMPPT(no int) MPPT {
	var entries []dxs.DxsID
	switch no {
	case 1:
		entries = []dxs.DxsID{
			dxs.DC1_I,
			dxs.DC1_U,
			dxs.DC1_P,
		}
	case 2:
		entries = []dxs.DxsID{
			dxs.DC2_I,
			dxs.DC2_U,
			dxs.DC2_P,
		}
	case 3:
		entries = []dxs.DxsID{
			dxs.DC3_I,
			dxs.DC3_U,
			dxs.DC3_P,
		}
	}

	data := dxs.Gather(inv.host, entries)
	return ParseMPPT(data, no)
}
