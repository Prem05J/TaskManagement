TASK MANAGMENT 

Fiber API - Go
This project demonstrates how to set up and run a Fiber-based API with two methods:

Using Docker for containerized deployment.
Using the traditional Go run method for local development.

## Installation
    - Close the Repository
        - git clone https://github.com/Prem05J/TaskManagement.git
        - cd TaskManagement

    - Install Dependencies
        - go mod tidy

## Running the Project

    - Option 1: Run with Docker
        - docker build -t taskmangement-api .
        - docker run -d -p 8080:8080 taskmangement-api

    - Option 2: Run with Go (Traditional Method)
        - Unit Testing
            - go test ./Test 
        - Run the Application
            - go run main.go


