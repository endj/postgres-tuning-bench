#!/bin/sh
boards=$(curl -k https://localhost/boards | jq '. | length')

if [ "$boards" -eq "3" ]; then
  echo "Boards: good"
else
  exit 0
fi


# Show threads
curl -k https://localhost/boards/news

# Create thread
curl -k -d '{ "title":"title", "post":"Some text"}' https://localhost/boards/news/threads

# List threads
curl -k https://localhost/boards/news

# Show thread
curl -k https://localhost/boards/news/threads/1

# Make reply
curl -k -d '{"text": "A reply"}' https://localhost/boards/news/threads/1/reply

# Show thread
curl -k https://localhost/boards/news/threads/1





