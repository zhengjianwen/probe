package img_cap

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestInitDir(t *testing.T) {
	if err := InitDir(); err != nil {
		t.Fatal(err)
	}
}

func TestCap(t *testing.T) {
	if err := Init(ImageLocalStoreDir); err != nil {
		t.Fatal(err)
	}

	if err := Cap("http://www.baidu.com", "baidu.png"); err != nil {
		t.Fatal(err)
	}
}

func TestCaps(t *testing.T) {
	for i := 0; i <= 30; i++ {
		go func(i int) {
			if err := Cap("http://www.baidu.com?fsdf", fmt.Sprintf("baidu_%d.png", i)); err != nil {
				t.Fatal(err)
			}
		}(i)
	}

}
