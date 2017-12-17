package jsonx

//Auto-generated code; DONT EDIT THIS CODE

import (
	"bytes"
	"testing"

	"github.com/rimpo/go-null/examples/example1/null"
	"github.com/tidwall/gjson"
)

func TestNullBuiltInTypes(t *testing.T) {
	type A struct {
		X null.String `json:"x"`
		B struct {
			Y null.Int `json:"y"`
			C struct {
				Z null.Float64 `json:"z"`
			} `json:"c"`
		} `json:"b"`
	}

	var a A
	a.X.Set("X")
	a.B.Y.Set(10)
	a.B.C.Z.Set(50.320924)

	var jsonData bytes.Buffer
	Marshal(&a, &jsonData)

	j := jsonData.String()

	x := gjson.Get(j, "x")

	if x.String() != a.X.Get() {
		t.Errorf("Expected:%v Got:%v", a.X.Get(), x.String())
	}

	y := gjson.Get(j, "b.y")

	if y.Int() != int64(a.B.Y.Get()) {
		t.Errorf("Expected:%v Got:%v", a.B.Y.Get(), y.Int())
	}

	z := gjson.Get(j, "b.c.z")

	if z.Float() != a.B.C.Z.Get() {
		t.Errorf("Expected:%v Got:%v", a.B.C.Z.Get(), z.Float())
	}

}

type Photo struct {
	Url null.String `json:"url"`
	ID  null.Int    `json:"id"`
}

type testProfile struct {
	Account struct {
		Memberlogin null.String `json:"memberlogin"`
	} `json:"account"`
	Basic struct {
		Username  null.String `json:"username"`
		FirstName null.String `json:"first_name"`
	} `json:"basic"`
	PhotoDetails struct {
		Photos []Photo `json:"photos"`
	} `json:"photo_details"`
	Traits struct {
		AboutMe null.String `json:"about_me"`
	} `json:"traits"`
}

func fillDummyData(p *testProfile) {
	p.Account.Memberlogin.Set("IDJDLKJS")
	p.Basic.Username.Set("rimpo")
	p.Basic.FirstName.Set("Manoj")
	var ph Photo
	ph.Url.Set("./test/abc")
	ph.ID.Set(1)
	p.PhotoDetails.Photos = append(p.PhotoDetails.Photos, ph)
	ph.Url.Set("./test/xyz")
	ph.ID.Set(2)
	p.PhotoDetails.Photos = append(p.PhotoDetails.Photos, ph)
	p.Traits.AboutMe.Set("Hi.. I'm a professional, independent and successful woman who would love to have that special person sharing my life and vice versa. My work life is not my only life by any means.... my passion for travel has taken me to many destinations far and wide to explore new things, people and places. I am physically fit, mentally astute and spiritually inclined having explored the magic of yoga. I am an active person with a zest for life enjoying all the various facets it has to offer. \r\nI am looking for an intelligent, good looking, strong minded and genuine guy ready for a committment who will enrich my life and vice versa (not talking about money here... though wouldn't say no!!!). \r\nIf you think you can be my inter-dependent partner... potentially... then no doubt you will reply back........ lets explore life together!!!!!! \r\n (for a response, would appreciate a photo - thanks!).")
}

func TestMarshal(t *testing.T) {
	var p testProfile
	fillDummyData(&p)

	var jsonData bytes.Buffer
	Marshal(&p, &jsonData)

	if len(jsonData.String()) == 0 {
		t.Errorf("Json output is empty")
	}
}
