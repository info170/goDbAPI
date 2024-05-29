## goDbAPI

MicroService for fetching data from database. (Postgres)

It create web-server for use as REST API.

#### Install

- cp .env.sample .env
- edit .env
- make request

#### Using:

Fetch all data from the table:

`GET http://localhost:8001/api/data/{table}?limit={limit}&offset={offset}`

Fetch specific entity: 

`GET http://localhost:8001/api/data/{table}/{id}`