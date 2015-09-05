#!/bin/bash
# project configuration
GO_PROJECT_NAME="bukkake"
GO_PROJECT_PATH="github.com/lordjpg/bukkake"
GO_PROJECT_OS="android"
GO_PROJECT_ID="BUKKAKE.UNIVERSAL.BUKKAKE"
error="WARNING: linker: libvc1dec_sa.ca7.so has text relocations. This is wasting memory and is a security risk. Please fix."

# building and installing
# based on prev configuration 
# make sure you have installed Go & gomobile 
cd $GOPATH/src/$GO_PROJECT_PATH 

echo ". . building $GO_PROJECT_ID . . "

gomobile build -target=$GO_PROJECT_OS 

if [ -e "adb" ]; then 
	./adb install -r $GO_PROJECT_NAME.apk
	./adb shell am start -n $GO_PROJECT_ID/org.golang.app.GoNativeActivity 
	./adb logcat
else
	adb install -r $GO_PROJECT_NAME.apk
	adb shell am start -n $GO_PROJECT_ID/org.golang.app.GoNativeActivity 
	adb logcat
fi

# todo: checking adb permission & installing w/o building