# Task Management REST API Documentation

This API allows basic management of tasks, including creating, retrieving, updating, and deleting tasks.

---

## Base URL

```
http://localhost:8080
```

---

## Endpoints

### ✅ GET `/tasks`

**Description:**  
Fetch all tasks in the system.

**Response:** `200 OK`
```json
[
  {
    "id": 1,
    "title": "Task 4",
    "description": "Read API documentation and complete task 4",
    "due_date": "2025-07-15",
    "status": "pending"
  }
]
```

---

### ✅ GET `/tasks/:id`

**Description:**  
Fetch a single task by its ID.

**Path Parameter:**
- `id` (int): Task ID

**Response:**
- `200 OK`: Task found
- `404 Not Found`: Task does not exist

**Example Response:**
```json
{
    "id": 1,
    "title": "Go Through Task 4",
    "description": "Read API documentation and complete task 4",
    "due_date": "2025-07-15",
    "status": "pending"
}
```

---

### ✅ POST `/tasks`

**Description:**  
Create a new task.

**Request Body:**
```json
{
  "title": "Read Go docs",
  "description": "Focus on the Gin framework",
  "due_date": "2025-07-15",
  "status": "pending"
}
```

**Response:**
- `201 Created`: Task created successfully

**Example Response:**
```json
{
  "id": 2,
  "title": "Read Go docs",
  "description": "Focus on the Gin framework",
  "due_date": "2025-07-15",
  "status": "pending"
}
```

---

### ✅ PUT `/tasks/:id`

**Description:**  
Update an existing task by ID.

**Path Parameter:**
- `id` (int): Task ID

**Request Body:**
```json
{
  "title": "Start reading Mongo DB docs",
  "description": "Go through Mongo DB docs",
  "due_date": "2025-07-22",
  "status": "completed"
}
```

**Response:**
- `200 OK`: Task updated
- `400 Bad Request`: Invalid payload
- `404 Not Found`: Task not found

---

### ✅ DELETE `/tasks/:id`

**Description:**  
Delete a task by ID.

**Path Parameter:**
- `id` (int): Task ID

**Response:**
- `200 OK`: Task deleted
- `404 Not Found`: Task not found

**Example Response:**
```json
{
  "message": "Task deleted"
}
```

---

## Error Responses

All error responses will return a JSON object with an `error` key:

```json
{
  "error": "Task not found"
}
```

---

## Notes

- All dates should be in string format (`YYYY-MM-DD`).
- The `status` field should be either:
  - `"pending"`
  - `"completed"`

---

## Testing

You can use [Postman](https://www.postman.com/) or `curl` to test the API:

```bash
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{
  "id": "5",
  "title": "Learn Go",
  "description": "Study Go REST API development",
  "due_date": "2025-08-15T09:00:00Z",
  "status": "Pending"
}'

curl -X DELETE http://localhost:8080/tasks/1

curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{
  "id": "1",
  "title": "Updated Task 1",
  "description": "Updated description",
  "due_date": "2025-08-20T12:00:00Z",
  "status": "In Progress"
}'

```

---
