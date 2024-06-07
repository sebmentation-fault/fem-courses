#!/usr/bin/env bash

languages=$(echo "golang c cpp typescript java rust zig haskell python" | tr " " "\n")
core_utils=$(echo "echo grep read head tail git tr xargs sort cat uniq curl find xargs sed awk pacman paru" | tr " " "\n")
selected=$(echo -e "$languages\n$core_utils" | fzf)

read -p "Query: " query

if echo "$languages" | grep -qs $selected; then
    tmux split-window -h bash -c "curl cht.sh/$selected/$(echo "$query" | tr " " "+") | less"
else
    tmux split-window -h bash -c "curl cht.sh/$selected~$query | less"
fi

