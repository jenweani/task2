## API Documentation
The Person API allows you to perform CRUD (Create, Read, Update, Delete) operations on person records. This API provides endpoints for managing individual person records.

#### Table of Contents
* [Introduction](#intro)
* [Base URL](#base-url)
* [Build and Run](#run)
* [Endpoints](#endpoints)
    * [Create a Person](#create)
    * [Retrieve a Person](#retrieve)
    * [Update a Person](#update)
    * [Delete a Person](#delete)
* [Request and Response Formats](#examples)

### <a name="intro"></a> Introduction
The Person API is a simple RESTful API that allows you to manage person records. Each person record has the following attributes:

ID: Unique identifier for the person, it auto increments.
Name: person's full name, should be unique.

#### UML Diagram of the Person Model
![UML diagram](https://github.com/jenweani/task2/blob/main/uml_diagram_person.jpg?raw=true)

#### <a name="base-url"></a> Base URL
The base URL for all API endpoints is:

```url
https://.com/api
```

### <a name="run"></a> Run
* build and run the api locally
```bash
go run main.go 
```
The api should run locally on port 80

### <a name="endpoints"></a> Endpoints

#### <a name="create"></a> Create a Person
Endpoint: POST /api
Description: Create a new person record.
Request Body: JSON object with the following fields:
* name (string, required): The person's full name.
Response: JSON object representing the created person record, HTTP status code 201.

#### <a name="retrieve"></a> Retrieve a Person
Endpoint: GET /api/{user_id}
Description: Retrieve a person record by ID. You can also retrieve a person using the name field passed as a JSON object but the user_id must be set to zero to do so.
Request Body: JSON object with the following fields:
* name (string, required): The person's full name.
Response: JSON object representing the retrieved person record, HTTP status code 200.

#### <a name="update"></a> Update a Person
Endpoint: PUT /api/{user_id}
Description: Update a person record by ID.
Request Body: JSON object with the following fields:
* name (string): The person's full name.
Response: JSON object representing the updated person record, HTTP status code 200.

#### <a name="delete"></a> Delete a Person
Endpoint: DELETE /api/{user_id}
Description: Delete a person record by ID.
Response: HTTP status code 204 (No Content) upon successful deletion.

### <a name="examples"></a> Request and Response Formats
These are examples showing the request formats allowed and their responses.
All responses use the JSON format.

#### Create a new person:

Request:
```bash
curl -X POST example.com/api -H "Content-Type: application/json" -d '{"name": "Mark Essien"}'
```

Response:

```json
{
    "status": 201,
    "message": "successfully created new user",
    "data": {
        "id": 1,
        "name": "Mark Essien"
    }
}
```

#### Retrieve a particular person by id:

Request:
```bash
curl example.com/api/2
```

Response:
```json
{
    "status": 200,
    "message": "successfully retrieved user from db",
    "data": {
        "id": 2,
        "name": "John Enweani"
    }
}
```

#### Retrieve a particular person by name:
Request:
```bash
curl example.com/api/0 -H "Content-Type: application/json" -d '{"name": "Mark Essien}
```
Response:
```json
{
    "status": 200,
    "message": "successfully retrieved user from db",
    "data": {
        "id": 1,
        "name": "Mark Essien"
    }
}
```

#### Update a Person:
Request:
```bash
curl -X PATCH example.com/api/1 -H "Content-Type: application/json" -d '{"name": "Elon musk"}'
```

Response:
```json
{
    "status": 200,
    "message": "successfully updated user",
    "data": {
        "id": 1,
        "name": "Elon musk"
    }
}
```

#### delete a Person:
Request:
```bash
curl -X DELETE example.com/api/1
```

Response:
```json
{
    "status": 204,
    "message": "successfully deleted user with id: 1",
    "data": null
}
```