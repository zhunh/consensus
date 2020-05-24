#!/bin/bash
#编写完脚本给脚本执行权限  chmod u+x test.sh


# go run test/TestMainNodeCli.go 10018
# go run test/TestMainNodeCli.go 10019

gnome-terminal -x bash -c "go run test/TestMainNodeCli.go 10017"
gnome-terminal -x bash -c "go run test/TestMainNodeCli.go 10018"
gnome-terminal -x bash -c "go run test/TestMainNodeCli.go 10019"
