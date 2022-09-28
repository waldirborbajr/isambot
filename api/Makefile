BOT_TOKEN := 5343272189:AAF5_yv9adxzqsNrYCqAY5jakgb4GqZFGBc
export BOT_TOKEN

BINARY_NAME=shoplist

help : Makefile
	@echo "\nAVAILABLE COMMANDS\n"
	@cat Makefile | grep "##" | sed -n 's/^## /make /p' | column -t -s ':' && echo ""

.DEFAULT_GOAL := help

## runbot: execute from bot/main.go
runbot:
	@echo "> "$(BOT_TOKEN)
	go run bot/main.go -token=$(BOT_TOKEN)

## buildbot: make binary version to test
buildbot:
	@echo "> building bot"
	go build bot/main.go -o ${BINARY_NAME}

## runsql: execute from sql/main.go
runsql:
	@echo "> running sql"
	go run sql/main.go

## buildsql: make a binary version to test
buildsql:
	@echo "> building SQL"
	go build sql/main.go bot/shoplist


