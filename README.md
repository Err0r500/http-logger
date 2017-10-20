# http-logger

Get the details of an HTTP request as a JSON response, useful for testing http calls

## Docker Image
download Docker image [dockerHub](https://hub.docker.com/r/matthieujacquot/http-logger/)

## Usage
- run the container, exposing the port you want and route it to :8080 :

```
docker run --rm -p 3000:8080 matthieujacquot/http-logger:0.1
```

- make an http call :

```
curl -X POST \
-H "Content-Type: application/json" \
-H "Authorization: myAuthToken" \
-d '{"string":"xyz","int":1234, "anObject": {"a":12345, "b":"someString"}, "anArray": [1,2,3]}' \
http://localhost:3000/anyPath
```

- parse the JSON response to check if what you just sent actually fits your expectations :
```
{
   "method":"POST",
   "path":"/anyPath",
   "headers":{
      "Accept":"*/*",
      "Authorization":"myAuthToken",
      "Content-Length":"90",
      "Content-Type":"application/json",
      "User-Agent":"curl/7.55.1"
   },
   "body":{
      "anArray":[
         1,
         2,
         3
      ],
      "anObject":{
         "a":12345,
         "b":"someString"
      },
      "int":1234,
      "string":"xyz"
   }
}
```