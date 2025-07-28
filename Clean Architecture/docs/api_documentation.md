# Task Management API Documentation

## Overview

This API supports user authentication and authorization using JWT, with role-based access for admins and regular users.

---

## Authentication & Authorization

- **Register:** `POST /api/register`
  - JSON: `{ "username": "yourname", "password": "yourpassword" }`
  - First user is an admin. Others are regular users by default.

- **Login:** `POST /api/login`
  - JSON: `{ "username": "yourname", "password": "yourpassword" }`
  - Response: `{ "token": "JWT_TOKEN" }`
  - Use this token in the `Authorization: Bearer JWT_TOKEN` header for all other endpoints.

---

## Roles

- **admin:** Can create, update, delete tasks, and promote users.
- **user:** Can only view tasks.

---

## Endpoints

### Register

- `POST /api/register`
- JSON: `{ "username": "...", "password": "..." }`
- Creates a new user.

### Login

- `POST /api/login`
- JSON: `{ "username": "...", "password": "..." }`
- Returns JWT token.

### Promote User (admin only)

- `POST /api/promote`
- JSON: `{ "username": "userToPromote" }`
- Promotes the user to admin.

---

### Tasks

- `GET /api/tasks` — List all tasks (auth required)
- `GET /api/tasks/:id` — Get task by ID (auth required)
- `POST /api/tasks` — Create task (admin only)
- `PUT /api/tasks/:id` — Update task (admin only)
- `DELETE /api/tasks/:id` — Delete task (admin only)

---

## JWT Usage

- After login, copy the token.
- For protected endpoints, set header: `Authorization: Bearer YOUR_TOKEN_HERE`

---

## Password Security

- Passwords are **hashed** using bcrypt before storage.
- Passwords are never returned or logged.

---

## Error Responses

- `401 Unauthorized`: Missing or invalid token.
- `403 Forbidden`: Insufficient permissions (e.g., non-admin accessing admin endpoint).
- `400 Bad Request`: Invalid input (e.g., duplicate username).
- `404 Not Found`: Resource not found.

---

## Testing

- Use Postman or curl to register, login, and access endpoints with or without JWT tokens.
- Only admins can create/update/delete tasks and promote users.

---

## Example Workflow

1. Register a user (becomes admin if first).
2. Login as that user, receive JWT token.
3. Use token to create tasks.
4. Register another user, login.
5. Only admin can promote the second user to admin.
