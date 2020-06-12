#!/bin/sh
boards=$(curl -k https://localhost/board | jq '. | length')

if [ "$boards" -eq 3 ]; then
  echo "Boards: good"
else
  exit 0
fi

curl -k -d '{ "title":"title", "post":"post"}' https://localhost/board/news/thread







