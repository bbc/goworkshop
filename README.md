# Go URL Shortener

This repository is the code to run a [Go](https://golang.org/) workshop. The code is separated into 4 goals with tests which pass when the goals are met. 

## Environment
To install the Go tools you should be able to run `brew install go`. For more in-depth instructions see [Install the Go tools](https://golang.org/doc/install#install)

## Goals
1. Accept `HTTP:GET` to "/" and return 501/Not Implemented 
2. Accept `HTTP:PUT` to "/" containing JSON with a URL and parse:
```
{ 
    "URL": "https://www.bbc.co.uk/iplayer" 
}
```
3. Return a shortened URL (using a very simple scheme that doesn't need to be cryptographically secure) in a JSON response:
```
{ 
    "ShortURL": "http://localhost:8080/1" 
}
```
4. Accept HTTP:GET to "/1" and redirect to correct URL

Stretch goals:
1. Count times the URL has been decoded
2. Delete URL from local database
3. Store key/values in a local JSON database
4. Make the encoded URL "cryptographically secure" while still short

## Running

### Test

```
go test ./goalN
```

### Run

```
go run ./goalN
```
