# Go Starter Template 🚀

A minimal and clean starter template for building Go web applications using:

- **Echo** – HTTP framework
- **SQLite** – lightweight embedded database
- **SQLC** – type-safe queries from SQL
- **Goose** – database migrations
- **Dependency-injection friendly architecture**
- **Unit tests**

---

## 📦 Project Structure

.
├── cmd/app/main.go         # Application entrypoint
├── internal/
│   ├── db/                 # DB init + SQLC generated code
│   ├── service/            # AuthorService + MemoryAuthorService
│   ├── handler/            # Echo HTTP handlers
│   ├── router/             # Route registration
│   └── model/              # View/data models
├── scripts/migrations/     # Goose SQL migration files
├── web/                    # Front-end static assets
├── example.sqlite          # SQLite database (auto-created)
├── Makefile
├── go.mod
└── README.md


---

## 🧱 Architecture
Handler (Echo)
     ↓
Service (business logic)
     ↓
Database (SQLC + SQLite)



### Handlers
- Handle HTTP only
- Depend on `AuthorService` interface

### Services
- Business logic

### DB
- SQLite
- SQLC for type-safe queries
- Goose for migrations

---

## 🗄️ Database

### Apply migrations
```shell
make up
```

### Roll back
```shell
make down
```


### Status
```shell
make status
```

### Generate SQLC code

```shell
sqlc generate

```

---

## ▶️ Run the App

```shell
make run

```

Server starts on:

http://localhost:3000


---

## 🧪 Tests

Includes:

- Handler tests with Echo + httptest
- Real DB tests using a test-specific DSN

Run tests:

```shell
go test ./...

```

---

## 🛠️ Development Tools

### Tailwind (optional)

```shell
make styles

```

### Auto-reload with wgo

```shell
make r

```

### Pre-commit

```shell
make pre-commit

```

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
