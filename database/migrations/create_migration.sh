#!/bin/bash

printf ">>> "
read name
filename=$(goose create $name sql 2>&1 | cut -d' ' -f6)
mv $filename ./database/migrations
