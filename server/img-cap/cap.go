package img_cap

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
driver.close()
driver.quit()
print("ok")
sys.exit(0)
`

var ImageLocalStoreDir string = "task_capture"
var ImageRequestBasePath string = "task_capture"

func Init(dir string) error {
	if len(dir) > 0 {
		ImageLocalStoreDir = dir
		log.Printf("network image caputre init store dir %s\n", dir)
	}

	if err := InitDir(); err != nil {
		return err
	}

	var fd *os.File
	if _, err := os.Stat("capture.py"); os.IsNotExist(err) {
		if fd, err = os.Create("capture.py"); err != nil {
			return err
		}
	} else if err = os.Remove("capture.py"); err != nil {
		return err
	}

	if _, err := fd.Write(getCaptureCodes()); err != nil {
		return err
	}

	if err := fd.Close(); err != nil {
		return err
	}

	return os.Chmod("capture.py", 0777)
}

func InitDir() error {
	_, err := os.Stat(ImageLocalStoreDir)
	if os.IsNotExist(err) {
		return os.Mkdir(ImageLocalStoreDir, 0777)
	}

	return nil
}

func Cap(url, image string) error {
	return exec.Command("./capture.py", url, image).Run()
}

func getCaptureCodes() []byte {
	return []byte(CapturePy)
}

func GetReqImgName(img string) string {
	return fmt.Sprintf("%s/%s", filepath.Clean(ImageRequestBasePath), img)
}

func GetLocalImgName(img string) string {
	return fmt.Sprintf("%s%s%s", filepath.Clean(ImageLocalStoreDir), string(filepath.Separator), img)
}
