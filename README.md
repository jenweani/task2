## HNGx stage two task
### Simple CRUD REST API for Managing a Person Data
This is a simple API that allows you to perform CRUD operations on a "Person" data structure. 
You can create, read, update, and delete person records using HTTP requests POST, GET, PATCH and DELETE respectively.

#### Live URL
Live URL to interact with the api: https://hng-task2-q5k1.onrender.com/api

#### UML Diagram of the Person Model
![UML diagram](https://github.com/jenweani/task2/blob/main/uml_diagram_person.jpg?raw=true)

#### Postman Test JSON File
The test file is in the repo
Link to [Test File](https://github.com/jenweani/task2/blob/main/hng-task2.postman_collection.json) 

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

### <a name="prereq"></a> Prerequisites
Before you begin, ensure you have met the following requirements:

* Go (Golang) installed on your machine.
* An IDE of your choice.
* An HTTP client (e.g., curl, Postman, or your web browser) to interact with the API.

### <a name="install"></a> Installation
* Clone the repository
```bash
git clone `https://github.com/jenweani/task2.git`
```

### <a name="run"></a> Run
* build and run the api
```bash
go run main.go 
```
The api should run locally on port 80

#### <a name="usage"></a> Usage
You can interact with the API using HTTP requests. Here are examples of how to perform CRUD operations on the "Person" data structure:

#### <a name="create"></a> Create a Person
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

#### <a name="retrieve"></a> Retrieve a Person
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

#### <a name="update"></a> Update a Person
Endpoint: PATCH /api/{user_id}

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

#### <a name="delete"></a> Delete a Person
Endpoint: DELETE /api/{user_id}

Response: 
```json
{
    "status": 204,
    "message": "successfully deleted user with id: 1",
    "data": null
}
```

### <a name="endpoints"></a> Endpoints
* POST /api: Create a new person.
* GET /api/{user_id}: Retrieve a person.
* PUT /api/{user_id}: Update a person.
* DELETE /api/{user_id}: Delete a person.

### <a name="examples"></a> Example Requests
The following are some requests that you can use to interact with the api

* Create a new person:
```bash
curl -X POST https://hng-task2-q5k1.onrender.com/api -H "Content-Type: application/json" -d '{"name": "Mark Essien"}'
```
* Retrieve a particular person:
```bash
curl https://hng-task2-q5k1.onrender.com/api/1
```
* Update a Person:
```bash
curl -X PATCH https://hng-task2-q5k1.onrender.com/api/1 -H "Content-Type: application/json" -d '{"name": "Elon musk"}'
```
* delete a Person:
```bash
curl -X DELETE https://hng-task2-q5k1.onrender.com/api/1
```