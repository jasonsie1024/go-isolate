package isolate

import "testing"

func Test(t *testing.T) {
	sandbox, err := New(0, true)
	if err != nil {
		t.Error(err)
	}

	if sandbox.Id != "0" {
		t.Error("Wrong ID")
	}

	err = sandbox.CleanUp()
	if err != nil {
		t.Error(err)
	}

	t.Log("success")
}
