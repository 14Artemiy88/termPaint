#!/bin/bash

git clone https://github.com/14Artemiy88/termPaint
cd termPaint || exit
go install .
mkdir "$HOME"/.config/termPaint
cp config.yaml "$HOME"/.config/termPaint
cd ..
yes | rm -r termPaint