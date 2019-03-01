package tcpengine

import (
	"fmt"
	"github.com/randyardiansyah25/iso8583u/parser"
	"strings"
	"testing"
)

func TestIso8583Engine(t *testing.T) {
	//support multiple iso field values
	isoEngine := GetEngine(40, 3, 41, 104) //assign multiple key

	//support multiple iso field values
	isoEngine.AddHandler(func(iso *iso8583uParser.ISO8583U) {
		fmt.Println(iso.GetField(3))
	}, "<iso_field_change_me>", "<iso_field_change_me>", "<iso_field_change_me>")
	err := isoEngine.RunInBackground("3301") //run on routine
	//err := isoEngine.Run("3301") //run with blocking
	if err != nil {
		fmt.Println("error")
	}

	s := "3:100700"
	if strings.Contains(s, ":") {
		fmt.Println(" `:` char is exist")
	} else {
		fmt.Println("doesn't exist")
	}
}
