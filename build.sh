#!/bin/bash
GO_PROJECT_NAME="bukkake"
GO_PROJECT_PATH="github.com/lordjpg/bukkake"
GO_PROJECT_OS="android"
GO_PROJECT_ID="BUKKAKE.UNIVERSAL.BUKKAKE"
error="WARNING: linker: libvc1dec_sa.ca7.so has text relocations. This is wasting memory and is a security risk. Please fix."

cd $GOPATH/src/$GO_PROJECT_PATH 
echo ". . . building $GO_PROJECT_ID . . ."

debug_output=$(gomobile build -target=$GO_PROJECT_OS)
while read -r line; do
	if [ "$line" != "" ]; then
		exit 1
	fi
done <<< "$debug_output"


output=$(adb devices)
while read -r line; do
   	if [ "$line" = " " ]; then
   		echo "attach device first"
   		exit 1
   	fi
done <<< "$output"

install_output=$(sudo ./adb install -r $GO_PROJECT_NAME.apk)
while read -r line; do 
	if [ ! "$line" == "WARNING: linker: libvc1dec_sa.ca7.so has text relocations. This is wasting memory and is a security risk. Please fix." ]; then 
		echo $line
	fi
done <<< "$install_output"

adb shell am start -n $GO_PROJECT_ID/org.golang.app.GoNativeActivity 
