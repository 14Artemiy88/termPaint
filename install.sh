#!/bin/bash

go install github.com/14Artemiy88/termPaint@latest
conf_dir="$HOME/.config/termPaint"
if [[ ! -d "$conf_dir" ]]; then
    mkdir $conf_dir
fi
file="$conf_dir/config.yaml"
if [[ ! -f "$file" ]]; then
    curl -o $file  https://raw.githubusercontent.com/14Artemiy88/termPaint/main/config.yaml
fi
