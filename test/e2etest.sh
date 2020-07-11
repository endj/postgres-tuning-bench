#!/bin/sh
boards=$(curl -k https://localhost/board | jq '. | length')

if [ "$boards" -eq "3" ]; then
  echo "Boards: good"
else
  exit 0
fi


# Show threads
curl -k https://localhost/board/news

# Create thread
curl -k -d '{ "title":"title", "post":"Some text"}' https://localhost/board/news/thread

# List threads
curl -k https://localhost/board/news

# Show thread
curl -k https://localhost/board/news/thread/1

# Make reply
curl -k -d '{"text": "A reply"}' https://localhost/board/news/thread/1/reply

# Show thread
curl -k https://localhost/board/news/thread/1





