package img_cap

import (
	"os"
	"os/exec"
)

var CapturePy = `#!/usr/bin/python
# -*- coding:utf8 -*-

from selenium import webdriver
import sys

args = sys.argv[1:]
if len(args) < 2:
    print("param need url and gen file")
    sys.exit(1)

driver = webdriver.PhantomJS(executable_path='/usr/bin/phantomjs')

driver.get(args[0])
driver.save_screenshot(args[1])

print("ok")
sys.exit(0)

`

func Init() error {
	var fd *os.File
	if _, err := os.Stat("capture.py"); os.IsNotExist(err) {
		if fd, err = os.Create("capture.py"); err != nil {
			return err
		}
	} else {
		if fd, err = os.Open("capture.py"); err != nil {
			return err
		}
	}

	if _, err := fd.Write(getCaptureCodes()); err != nil {
		return err
	}

	if err := fd.Close(); err != nil {
		return err
	}

	return os.Chmod("capture.py", 0777)
}

func Cap(url, target string) error {
	return exec.Command("./capture.py", url, target).Run()
}

func getCaptureCodes() []byte {
	return []byte(CapturePy)
}