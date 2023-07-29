.PHONY: clean run

clean:
	rm -rf pb_data
	@echo "cleaned!"

run:
	go run *.go serve --http=127.0.0.1:8080