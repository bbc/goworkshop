# Go URL Shortener

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
4. Accept `HTTP:GET` to "/1" and redirect to the correct URL

Stretch goals:
1. Avoid generating short URLs for duplicate URLs
2. Store key/values in a local JSON database 
3. Delete URL from local database (HTTP:DELETE)
4. Count times the URL has been decoded 
5. Make the encoded URL "cryptographically secure"


## Running

### Test

```
go test ./goalN
```

### Run

```
go run ./goalN
```
