# Dockerfile for GO application

Simple Docker to dockerize golang application.

For building the image :
```
docker build --rm -t my-golang-app:latest .
```
Than run it :

```
docker run my-golang-app -p 800:8000
```
