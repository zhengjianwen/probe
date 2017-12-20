#!/usr/bin/python
# -*- coding:utf8 -*-

from selenium import webdriver
import sys

# args[0] = url args[1] target file

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
