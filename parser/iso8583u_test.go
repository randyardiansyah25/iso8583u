package iso8583uParser

import (
	"fmt"
	"testing"
)

var isomsg string = "2200BF38404109E200080000000013000000100700000000000001000000000000000000000000000012130905420000000019996324761520181213090542121300170020040002200073247615000628112072277    86118903227013312001.01.00023004test0403d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d06990001160002001.01.00001006TINTCR"

func TestISO8583U_GoUnMarshal(t *testing.T) {
	//iso,_ := NewISO8583U("../isopackager.yml")
	iso, _ := NewISO8583U()
	err := iso.GoUnMarshal(isomsg)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log("Message :\n", isomsg)
	t.Log("Parse : \n", iso.PrettyPrint())
	t.Log("Get Amount : ", iso.GetField(4))
}

func TestISO8583U_GoMarshal(t *testing.T) {

	iso, err := NewISO8583U("../isopackager.yml")
	//iso, err := NewISO8583U()
	if err != nil {
		fmt.Println("load package error", err.Error())
		return
	}
	iso.SetMti("2200")
	iso.SetField(3, "100700")
	iso.SetField(4, 10000)
	iso.SetField(5, 0)
	iso.SetField(6, 0)
	iso.SetField(7, "1213090542")
	iso.SetField(8, 0)
	iso.SetField(11, "199963247615")
	iso.SetField(12, "20181213090542")
	iso.SetField(13, "1213")
	iso.SetField(18, "0017")
	iso.SetField(26, "0020")
	iso.SetField(32, "0002")
	iso.SetField(37, "200073247615")
	iso.SetField(40, "000")
	iso.SetField(41, "628112072277")
	iso.SetField(42, "861189032270133")
	iso.SetField(43, "001.01.00023")
	iso.SetField(47, "test")
	iso.SetField(61, "3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d")
	iso.SetField(100, "990001")
	iso.SetField(103, "0002001.01.00001")
	iso.SetField(104, "TINTCR")
	msgiso, err := iso.GoMarshal()
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("Set Fields : \n", iso.PrettyPrint())
	t.Log("Result : ", msgiso)
}
