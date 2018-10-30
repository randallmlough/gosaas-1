#!/bin/bash

clear
rm gosaas-dev
go build -o gosaas-dev
./gosaas-dev -driver bolt -datasource "db.bolt"