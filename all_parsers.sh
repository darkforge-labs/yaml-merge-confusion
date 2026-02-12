#!/bin/bash

echo -e "\r\n======\r\n[+]Ruby\r\n======\r\n"
ruby rubyparse.rb

echo -e "\r\n======\r\n[+]GO\r\n======\r\n"
./goparser

echo -e "\r\n======\r\n[+]Python\r\n======\r\n"
python3.11 pyparse.py

echo -e "\r\n======\r\n[+]Node\r\n======\r\n"
node nodeparse.js
