#!/bin/bash

go build -o bookings cmd/web/* && ./bookings -dbname=bookings -dbuser=postgres -cache=false -production=false