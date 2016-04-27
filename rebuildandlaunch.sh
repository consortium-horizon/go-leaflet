#!/bin/bash
screen -X quit
go build go-leaflet.go
screen -d -m /home/vagrant/code/go-leaflet/go-leaflet