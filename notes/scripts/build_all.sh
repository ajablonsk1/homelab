#!/bin/bash

cd "../inbox" && go build -o /usr/local/bin/inbox
cd "../zettel" && go build -o /usr/local/bin/zettel
cd "../material" && go build -o /usr/local/bin/material

echo "All builds finished successfully!"
