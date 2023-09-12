## HNGx stage two task
### Simple CRUD REST API for Managing a Person Data
This is a simple API that allows you to perform CRUD operations on a "Person" data structure. 
You can create, read, update, and delete person records using HTTP requests POST, GET, PATCH and DELETE respectively.

#### Live URL
Click [link]() to interact with the api

#### UML Diagram of the Person Model
[UML diagram](https://github.com/jenweani/task2/blob/main/uml_diagram_person.jpg?raw=true)

### Table of Contents
* [Prerequisites](#prereq)
* [Installation](#install)
* [Running the Api](#run)
* [Usage](#usage)
    * [Create a Person](#create)
    * [Retrieve a Person](#retrieve)
    * [Update a Person](#update)
    * [Delete a Person](#delete)
* [Endpoints](#endpoints)
* [Example Requests](#examples)

### Prerequisites {#prereq}
Before you begin, ensure you have met the following requirements:

* Go (Golang) installed on your machine.
* An IDE of your choice.
* An HTTP client (e.g., curl, Postman, or your web browser) to interact with the API.

### Installation {#install}
* Clone the repository
```bash
git clone `https://github.com/jenweani/task2.git`
```

### Run {#run}
* build and run the api
```bash
go run main.go 
```
The api should run locally on port 80

#### Usage {#usage}
You can interact with the API using HTTP requests. Here are examples of how to perform CRUD operations on the "Person" data structure:

#### Create a Person {#create}
Endpoint: POST /api

Request Body:

```json
{
    "name": "John Enweani",
}
```
Response:

```json
{
    "status": 201,
    "message": "successfully created new user",
    "data": {
        "id": 1,
        "name": "John Doe"
    }
}
```

#### Retrieve a Person {#retrieve}
Endpoint: GET /api/{user_id}

Response:

```json
{
    "status": 200,
    "message": "successfully retrieved user from db",
    "data": {
        "id": 1,
        "name": "John Doe"
    }
}
```

#### Update a Person {#update}
Endpoint: PATCH /api/{id}

Request Body:

```json
{
    "name": "John E"
}
```
Response:

```json
{
    "status": 200,
    "message": "successfully updated user",
    "data": {
        "id": 1,
        "name": "John E"
    }
}
```

#### Delete a Person {#delete}
Endpoint: DELETE /api/{user_id}

Response: 
```json
{
    "status": 204,
    "message": "successfully deleted user with id: 1",
    "data": null
}
```

### Endpoints {#endpoints}
* POST /api: Create a new person.
* GET /api/{user_id}: Retrieve a person.
* PUT /api/{user_id}: Update a person.
* DELETE /api/{user_id}: Delete a person.

### Example Requests {#examples}
The following are some requests that you can use to interact with the api

* Create a new person:
```bash
curl -X POST example.com/api -H "Content-Type: application/json" -d '{"name": "Mark Essien"}'
```
* Retrieve a particular person:
```bash
curl example.com/api/1
```
* Update a Person:
```bash
curl -X PATCH example.com/api/1 -H "Content-Type: application/json" -d '{"name": "Elon musk"}'
```
* delete a Person:
```bash
curl -X DELETE example.com/api/1
```