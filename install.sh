#!/bin/bash

go install github.com/14Artemiy88/termPaint@latest
conf_dir="$HOME/.config/termPaint"
if [[ ! -d "$conf_dir" ]]; then
    mkdir "$conf_dir"
fi

FILE="$conf_dir/config.yaml"
if [ -f "$FILE" ]; then
    mv "$FILE" "$FILE.old"
fi

curl -o "$FILE"  https://raw.githubusercontent.com/14Artemiy88/termPaint/main/config.yaml
