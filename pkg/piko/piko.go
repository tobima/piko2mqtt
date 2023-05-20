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
