#!/bin/bash
# setting up the environment
# as from v1.10, GOROOT is optional and is deduced from
# the binary path
#export GOROOT=/usr/local/opt/go/libexec

# as from v1.10, GOTMPDIR was added to control where temporary files
# are created
#export GOTMPDIR=/tmp/go

# this is where your Go Lanugage project workspace resides
# note that, in contrary to other programming lanuages where
# a project consists of src/main and src/test dirs
# and the project itself is located in a single repository
# a "Go Workspace" consists of several projects, it is one big
# codebase if you will
#
# as from v1.8, GOPATH is optional, if not set, it uses ~/go
export GOPATH=~/projects/golang