#!/bin/bash

git clone https://github.com/14Artemiy88/termPaint
cd termPaint || exit
go install .
conf_dir="$HOME/.config/termPaint"
if [[ ! -d $conf_dir ]]; then
    mkdir $conf_dir
fi
if [[ ! -f "$conf_dir/config.yaml" ]]; then
    cp config.yaml "$conf_dir"
fi
cd ..
yes | rm -r termPaint