#!/bin/bash

export ADDRESS=":9000"

go run ./cmd/web -address=$ADDRESS
