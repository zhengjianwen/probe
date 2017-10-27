package stat

import (
	"testing"
)

func TestListWorkingWorkers(t *testing.T) {
	_, err := ListWorkingWorkers()
	if err != nil {
		t.Fatal(err)
	}
}

func TestListWorkerWithCached(t *testing.T) {
	wks, err := ListWorkingWorkers()
	if err != nil {
		t.Fatal(err)
	}

	wks2, err2 := ListWorkerWithCached()
	if err2 != nil {
		t.Fatal(err)
	}

	if len(wks) != len(wks2) {
		t.Fatal()
	}
}