package v

import "testing"

func TestString(t *testing.T) {
	Version = "devel"
	Revision = "1.0"
	GitSHA1 = "abcdefg"
	BuildDate = "2018-10-29T14:22:26+01:00"

	v := &version{}
	if v.String() != "devel (abcdefg), 1.0, 2018-10-29T14:22:26+01:00" {
		t.Errorf("version expectation not met, got: %s", v.String())
	}
}
