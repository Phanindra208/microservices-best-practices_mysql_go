
# Best practices for Golang microservices

This program helps in describing the go language microservice based  application architectures and using go packages in the most effective manner. 


## Features
- Make file 
- OpenAPI documentation
- Swagger integrated
- Tracing Enbled by default 
- logging enabled
- Multi-stage Docker file 
- Configurations load upon Environment 
- Unittest coverage 
- Docker Compose


## Installation

Install Mysql Docker server 

```bash
docker run --name mysql -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=pass123 \
    --restart unless-stopped \
    mysql:latest
```
    
Install zipkin server for tracing

```bash
docker run -d -p 9411:9411 openzipkin/zipkin
```
 Create Database    

```bash
CREATE DATABASE school
```
Genarate API from Swagger 
```bash
make genswag
```
## Demo

run code

```bash
run main.go or run from VS code
```
Documentation Avilabele on 

http://localhost:8888/docs
