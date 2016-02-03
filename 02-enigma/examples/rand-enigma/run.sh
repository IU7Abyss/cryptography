#!/bin/bash

go run rand-enigma.go text.txt cryptedtext.txt
go run rand-enigma.go cryptedtext.txt uncryptedtext.txt
cmp text.txt uncryptedtext.txt

go run rand-enigma.go image.png cryptedimage.png
go run rand-enigma.go cryptedimage.png uncryptedimage.png
cmp image.png uncryptedimage.png

go run rand-enigma.go archive.zip cryptedarchive.zip
go run rand-enigma.go cryptedarchive.zip uncryptedarchive.zip
cmp archive.zip uncryptedarchive.zip

