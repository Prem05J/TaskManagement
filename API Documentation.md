# API Documentation

This API provides endpoints for user authentication (sign-up, sign-in, sign-out) and task management (creating, fetching , updating, and deleting tasks). The API is built using the Go Fiber framework , uses JWT (JSON Web Tokens) for user authentication and MongoDb Atlas for Database.

---

## 1. User Authentication

### 1.1 Sign Up

- **Endpoint:** `/signUp`
- **Method:** POST
- **Request Body:**
    ```json
    {
        "fullName" : "Prem Kumar",
        "userName" : "prem310509@gmail.com",
        "password" : "cricket@guitar123"    
    }
    ```
    - `fullName` (string): Full name of the user.
    - `userName` (string, required): Email of the user (must be a valid email).
    - `password` (string, required): Password for the account.

- **Response:**
    - **Success (200):**
        ```json
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InByZW0zMTA1MDlAZ21haWwuY29tIiwiZXhwIjoxNzMxMDM1NTczLCJpZCI6IjY3MmQyYTg4NWJmOGMwZDdkZmYzNmYzMCIsIm5hbWUiOiJQcmVtIEt1bWFyIn0.Ta_QIzktFduK2Nof4ZWoHSA3k6JdPUR8IjWg7K5pgnE"
        }
        ```
    - **Failure (400, 500):** Error message with status.

- **Authentication:** No authentication required (sign-up process).

---

### 1.2 Sign In

- **Endpoint:** `/signIn`
- **Method:** POST
- **Request Body:**
    ```json
    {
        "username" : "prem310509@gmail.com",
        "password" : "cricket@guitar123"
    }
    ```
    - `userName` (string, required): Email of the user.
    - `password` (string, required): Password for the user.

- **Response:**
    - **Success (200):**
        ```json
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InByZW0zMTA1MDlAZ21haWwuY29tIiwiZXhwIjoxNzMxMDM1NTczLCJpZCI6IjY3MmQyYTg4NWJmOGMwZDdkZmYzNmYzMCIsIm5hbWUiOiJQcmVtIEt1bWFyIn0.Ta_QIzktFduK2Nof4ZWoHSA3k6JdPUR8IjWg7K5pgnE"
        }
        ```
    - **Failure (400, 401, 500):** Error message with status.

- **Authentication:** No authentication required (sign-in process).

---

### 1.3 Sign Out

- **Endpoint:** `/signOut`
- **Method:** POST
- **Request Headers:**
    - `Authorization`: Bearer `<JWT_TOKEN>`
  
- **Response:**
    - **Success (200):**
        ```json
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDI0LTExLTA4VDA2OjQzOjU3LjgwNDMxNjA0MiswNTozMCJ9.y9-7FD95BiGGhLiklFO_YaxGLXXJ_lahVLZAzWI0FLM"
        }
        ```
    - **Failure (401):** Error message (if JWT token is missing or invalid).

- **Authentication:** Requires valid JWT token in the Authorization header.

---

## 2. Task Management

### 2.1 Create Task

- **Endpoint:** `/tasks`
- **Method:** POST
- **Request Body:**
    ```json
    {
        "title": "Hellowin Party Arraement",
        "description": "The login page is not responding after submitting credentials.",
        "status": "in-progress",
        "assignedTo": "John Doe",
        "createdAt": "2024-11-07T14:30:00Z"
    }
    ```
    - `title` (string, required): Title of the task.
    - `description` (string): Description of the task.
    - `status` (string): Status of the task (e.g., pending, in-progress, completed).
    - `assignedTo` (string): Username of the user assigned to the task.
    - `createdAt` (string): Date and time of task creation (ISO 8601 format).

- **Response:**
    - **Success (200):**
        ```json
        {
            "id": "672d74c15628c152d7c3da53"
        }
        ```
    - **Failure (400, 500):** Error message with status.

- **Authentication:** Requires valid JWT token in the Authorization header.

---

### 2.2 Get All Tasks

- **Endpoint:** `/tasks`
- **Method:** GET
- **Response:**
    - **Success (200):**
        ```json
        [
          {
            "id": "task_id_here",
            "title": "Task title",
            "description": "Task description",
            "status": "pending",
            "assignedTo": "username@example.com",
            "createdAt": "2024-11-08T10:00:00Z"
          },
          {
            "id": "task_id_here",
            "title": "Task title",
            "description": "Task description",
            "status": "pending",
            "assignedTo": "username@example.com",
            "createdAt": "2024-11-08T10:00:00Z"
          },
          
        ]
        ```
    - **Failure (500):** Error message with status.

- **Authentication:** Requires valid JWT token in the Authorization header.

---

### 2.3 Get Task by ID

- **Endpoint:** `/tasks/:id`
- **Method:** GET
- **Response:**
    - **Success (200):**
        ```json
        {
          "id": "task_id_here",
          "title": "Task title",
          "description": "Task description",
          "status": "pending",
          "assignedTo": "username@example.com",
          "createdAt": "2024-11-08T10:00:00Z"
        }
        ```
    - **Failure (400, 404):** Error message (e.g., Task not found).

- **Authentication:** Requires valid JWT token in the Authorization header.

---

### 2.4 Update Task

- **Endpoint:** `/tasks/:id`
- **Method:** PUT
- **Request Body:**
    ```json
    {
        "title": "Hellowin Party Arraement",
        "description": "The login page is not responding after submitting credentials.",
        "status": "in-progress",
        "assignedTo": "John Doe",
        "createdAt": "2024-11-07T14:30:00Z"
    }
    ```
    - `title` (string, optional): Updated title of the task.
    - `description` (string, optional): Updated description of the task.
    - `status` (string, optional): Updated status of the task.
    - `assignedTo` (string, optional): Updated task assigned to.
    - `createdAt` (string, optional): Updated timestamp of the task.

- **Response:**
    - **Success (200):**
        ```json
        {
            "id": "672d74c15628c152d7c3da53"
        }
        ```
    - **Failure (400, 500):** Error message with status.

- **Authentication:** Requires valid JWT token in the Authorization header.

---

### 2.5 Delete Task

- **Endpoint:** `/tasks/:id`
- **Method:** DELETE
- **Response:**
    - **Success (200):**
        ```json
        {
          "result": "Task Deleted"
        }
        ```
    - **Failure (400, 500):** Error message with status.

- **Authentication:** Requires valid JWT token in the Authorization header.

---

## 3. Authentication & Security

- **JWT Token:**
  - After a successful sign-in, a JWT token is generated and returned in the response.
  - The token must be included in the Authorization header for all subsequent requests to protected endpoints (e.g., Bearer `<JWT_TOKEN>`).
  - The token will be used for user identification and authentication.

---

## 4. Error Handling

- **Error Structure:**
    All errors follow the same format:
    ```json
    {
      "error": "Error message here"
    }
    ```

- **Common Status Codes:**
  - **200 OK:** Successful request.
  - **400 Bad Request:** Invalid request or parameters.
  - **401 Unauthorized:** Missing or invalid authentication.
  - **404 Not Found:** Resource not found.
  - **500 Internal Server Error:** Unexpected server error.

---