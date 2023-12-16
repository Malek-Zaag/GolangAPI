git: 
	@git add .
	@git commit -m "$(msg)"
	@git push -u origin main


build:
	@go build ./main.go

run: 
	@go run main.go

exec:
	./main.exe