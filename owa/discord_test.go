package owa

import "testing"

func TestDiscord(t *testing.T) {

	d := &Discord{}
	if e := d.Send("This is test"); e != nil {
		t.Error(e)
	}
}
