# MICROBLOG
A microblog example for [interstellar](https://github.com/anddimario/interstellar)

### Requirements
GoLang, redis, interstellar, sqlite

### Installation
Clone the repository, then build:
```
go build create.go
go build retrieve.go
go build delete.go
go build form.go
```
Create an sqlite database with the table:    
`create table posts (title varchar(25), text varchar(50));`    
Then configure redis:    
```
set config:localhost:3000:db_path /path/where/your/db/is/foo.sqlite3
hset interstellar:vhost:localhost:3000:/create method POST
hset interstellar:vhost:localhost:3000:/create commands "cd /path/where/your/build/is && ./create"
hset interstellar:vhost:localhost:3000:/ method GET
hset interstellar:vhost:localhost:3000:/ commands "cd /path/where/your/build/is && ./retrieve"
hset interstellar:vhost:localhost:3000:/delete method DELETE
hset interstellar:vhost:localhost:3000:/delete commands "cd /path/where/your/build/is && ./delete"
hset interstellar:vhost:localhost:3000:/form method GET
hset interstellar:vhost:localhost:3000:/form commands "cd /path/where/your/build/is && ./form"
```
Now you can test:    
In browser: http://localhost:3000/form   
`curl http://localhost:3000/?title=title`     
`curl -XDELETE http://localhost:3000/delete?title=title`
