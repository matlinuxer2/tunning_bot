all:
	@./tunning_bot

go:
	go fmt latency.go
	go build latency.go
	./latency

clean:
	@ [ -e latency ] && rm -v latency || true
	@ [ -e tunning_bot.db ] && rm -v tunning_bot.db || true
