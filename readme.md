# Running the Project

The project consists of 3 microservices, 1 frontend, and 1 MongoDB database.

To run the MongoDB database locally, run the following command:

``` 
docker run --name meu-mongo -d -p 27017:27017 mongo
```

Before starting, create an .env file in the backend, monitor, and sso directories, and copy the content from the .env.example file.

Make sure that the MongoDB URL matches the one you want to run.

To run the monitor microservice, execute the following commands:

```
cd ./monitor
npm install
npx tsc
node dist/main.js
```

To run the sso microservice, execute the following commands:

```
cd ./sso
go run main.go
```

To run the backend microservice, execute the following commands:

```
cd ./backend
go run main.go
```