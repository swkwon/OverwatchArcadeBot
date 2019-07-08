#!/bin/bash
echo "pull from https://github.com/swkwon/OverwatchArcadeBot"
git pull
echo "building..."
go build
echo "copy..."
cp ./OverwatchArcadeBot /home/swkwon04/bin/
echo "finish"

