package main

import (
	"fmt"

	"github.com/tobima/piko2mqtt/pkg/dxs"
	"github.com/tobima/piko2mqtt/pkg/piko"
)

func main() {
	entries := []dxs.DxsID{
		dxs.InvName,
		dxs.TypeName,
		dxs.Serial,
		dxs.Mac,
		dxs.DC1_I,
		dxs.DC1_P,
		dxs.DC1_U,
		dxs.DC2_I,
		dxs.DC2_P,
		dxs.DC2_U,
		dxs.DC3_I,
		dxs.DC3_P,
		dxs.DC3_U,
	}
	data := dxs.Gather("192.168.178.185", entries)

	//	for _, entry := range data.DxsEntries {
	//		fmt.Printf("%s: %s \n", entry.ID, entry.Value)
	//	}

	fmt.Println(piko.ParseMPPT(data, 1))
	fmt.Println(piko.ParseMPPT(data, 2))
	fmt.Println(piko.ParseMPPT(data, 3))

	entries = piko.LogEntriesIDs
	entries = append(entries, dxs.LogEntries)
	data = dxs.Gather("192.168.178.185", entries)
	fmt.Println(piko.ParseEvents(data))

}
