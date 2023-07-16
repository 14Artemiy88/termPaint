#!/bin/bash

go install github.com/14Artemiy88/termPaint@latest
conf_dir="$HOME/.config/termPaint"
if [[ ! -d "$conf_dir" ]]; then
    mkdir $conf_dir
fi
if [[ ! -f "$conf_dir/config.yaml" ]]; then
    curl -O https://raw.githubusercontent.com/14Artemiy88/termPaint/main/config.yaml "$conf_dir"
fi
