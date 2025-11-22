export PORT=:3000

styles:
	npx @tailwindcss/cli -i ./web/static/css/input.css -o ./web/static/css/output.css --watch

run:
	PORT=$(PORT) go run ./cmd/app/main.go

r:
	wgo -file=.go -file=.html go run ./cmd/app/main.go

pre-commit:
	pip install pre-commit
	pre-commit install --install-hooks --overwrite

status:
	goose -dir "./scripts/migrations" sqlite3 ./example.sqlite status

up:
	goose -dir "./scripts/migrations" sqlite3 ./example.sqlite up

down:
	goose -dir "./scripts/migrations" sqlite3 ./example.sqlite down
