# Simple CRUD in Go

Just trying to make a simples CRUD in go after studying the whole documentation for two weeks

## How to Run

### Requirements

- Go >= 1.25

### Running in development

```bash
git clone <your-repo> && cd simple-crud
go mod download
go run ./cmd/simple-crud
```

### Building

```bash
go build -o simple-crud ./cmd/simple-crud
./simple-crud
```

The service runs by default at:

```
http://localhost:8080
```

---

## Endpoints

### GET /account/balance

Returns the account balance for a given user.

**Query parameters:**

- `username` — required

**Headers:**

- `Authorization: <token>`

**Example:**

```bash
curl -i "http://localhost:8080/account/balance?username=rafael"   -H "Authorization: 789GHI"
```

**Success Response (200):**

```json
{
  "code": 200,
  "balance": 3817.98
}
```

---

## Technologies Used

- **Go**
- **Chi Router** (`github.com/go-chi/chi`)
- **Gorilla Schema** (`github.com/gorilla/schema`)
- **Logrus** (`github.com/sirupsen/logrus`)
- **Mock database** (in `internal/tool/mockdb.go`)

---

## Project Structure

```
api/
cmd/simple-crud/
internal/
  handler/
  middleware/
  tool/
```

---

## Notes

- Authentication is done via middleware by validating a username + token pair.
- The current database is a mock implementation; replace it with a real database by implementing `DatabaseInterface`.
