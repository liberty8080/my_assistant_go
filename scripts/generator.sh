#! /bin/bash
xorm reverse mysql "aide:jacob_aide@(192.168.98.100:3306)/jacob_aide?charset=utf8" \
../venv/pkg/mod/github.com/go-xorm/cmd/xorm@v0.0.0-20190426080617-f87981e709a1/templates/goxorm \
../app/model



