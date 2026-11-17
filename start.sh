#!/usr/bin/env bash

set -e

# заходим в дерикторию проекта
cd "$(dirname "${BASH_SOURCE}")" || exit 1

# собираем бинарник
CGO_ENABLED=0 GOOS=linux go build -o bot ./cmd/bot

# запускаем через докер
docker build -t clarity .
docker run -d --restart unless-stopped --name clarity --env-file .env clarity
