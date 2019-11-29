#!/bin/bash
sudo apt-get update -y && sudo apt-get upgrade -y
sudo apt-get install $(grep -vE "^\s*#" deps.txt  | tr "\n" " ")
