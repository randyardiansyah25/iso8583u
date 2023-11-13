package tcpengine

import (
	"fmt"
	"strings"
	"testing"

	iso8583uParser "github.com/randyardiansyah25/iso8583u/parser"
)

func TestIso8583Engine(t *testing.T) {
	//support multiple iso field values
	isoEngine := GetEngine(40, 104) //assign multiple key

	//support multiple iso field values
	isoEngine.AddHandler(func(iso *iso8583uParser.ISO8583U) {
		fmt.Println(iso.GetField(3))

	}, "EDUPCR")
	//err := isoEngine.RunInBackground("3301") //run on routine

	err := isoEngine.Run("3301") //run with blocking
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
