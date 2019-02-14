package tcpengine

import (
	"fmt"
	"github.com/randyardiansyah25/libpkg/iso8583uparser"
	"testing"
)

func TestIso8583Engine(t *testing.T) {
	isoEngine := GetEngine(40, 3) // fieldNumberKey = 3; current project using processing code as key
	isoEngine.AddHandler("<iso_field_change_me>", func(iso *iso8583uParser.ISO8583U) {
		fmt.Println(iso.GetField(3))
	})
	err := isoEngine.RunInBackground("3301") //run on routine
	//err := isoEngine.Run("3301") //run with blocking
	if err != nil {
		fmt.Println("error")
	}

}
