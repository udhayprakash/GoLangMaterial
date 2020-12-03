package conv

import "testing"

func TestConvert(t *testing.T) {
	str, err := Convert("50 mi" , "km")
	if err != nil {
		t.Log("error should be nil",err)
		t.Fail()
	}
	if str != "80.47km" {
		t.Log("error should be 80.47km, but got",err)
		t.Fail()
	}
}