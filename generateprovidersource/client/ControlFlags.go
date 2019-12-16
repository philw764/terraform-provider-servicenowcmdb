package client

import (
	"fmt"
	"strings"
)

type key struct {
	CiName			string
	CiAttribute 	string
}

func GenerateControlFlagFile(CiList []CmdbCIMetaModel) error {
	var controlFlags = make(map[key]string)

	for ci := range CiList {
		for attr := range CiList[ci].Result.Attributes {

			if CiList[ci].CiName == "cmdb_ci" && strings.Contains( CiList[ci].Result.Attributes[attr].Element, "sys_") {
				controlFlags[key{"ALL",CiList[ci].Result.Attributes[attr].Element}] = "Computed:true"
				continue
			}


		}

	}
	fmt.Printf("This is Controlflag:%s", controlFlags[key{"all", "sys_changed"}])

	return nil
}
