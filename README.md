# Go Starter Template 🚀

A minimal and clean starter template for building Go web applications using:

- **Echo** – HTTP framework
- **SQLite** – lightweight embedded database
- **SQLC** – type-safe queries from SQL
- **Goose** – database migrations
- **Dependency-injection friendly architecture**
- **Real & in-memory services**
- **Unit tests**

---

## 📦 Project Structure

.
├── cmd/app/main.go # Application entrypoint
├── internal/
│ ├── db/ # DB init + SQLC generated code
│ ├── service/ # AuthorService + MemoryAuthorService
│ ├── handler/ # Echo HTTP handlers
│ ├── router/ # Route registration
│ └── model/ # View/data models
├── scripts/migrations/ # Goose migration files
├── web/ # Static assets
├── Makefile # Dev helpers
├── go.mod
└── README.md


---

## 🧱 Architecture

Handler (Echo)
↓
Service (business logic)
↓
DB Layer (SQLC)


### Handlers
- Handle HTTP only
- Depend on `AuthorService` interface

### Services
- Business logic
- Two implementations:
  - `authorService` (SQLite + SQLC)
  - `MemoryAuthorService` (in-memory for testing)

### DB
- SQLite
- SQLC for type-safe queries
- Goose for migrations

---

## 🗄️ Database

### Apply migrations

make up


### Roll back

make down


### Status

make status


### Generate SQLC code

sqlc generate


---

## ▶️ Run the App

make run


Server starts on:

http://localhost:3000


---

## 🧪 Tests

Includes:

- Handler tests with Echo + httptest
- Real DB tests using a test-specific DSN
- Optional in-memory service (`MemoryAuthorService`)

Run tests:

go test ./...


---

## 🛠️ Development Tools

### Tailwind (optional)

make styles


### Auto-reload with wgo

make r


### Pre-commit

make pre-commit


---

## 🧩 Features

- Clean starter template
- Simple CRUD example (Authors)
- Service interfaces + multiple implementations
- SQLite + Goose + SQLC setup
- Easy to extend and test
- Minimal dependencies
- Production-friendly structure

---

## ❤️ Contributing

PRs and issues are welcome!

---

## ⭐ License

MIT License.
