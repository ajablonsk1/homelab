#!/bin/bash

cd "../inbox" && go build -o /usr/local/bin/inbox
cd "../zettel" && go build -o /usr/local/bin/zettel
cd "../media" && go build -o /usr/local/bin/media
cd "../todo" && go build -o /usr/local/bin/todo

echo "All builds finished successfully!"
