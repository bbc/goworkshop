# Go URL Shortener

## Goals
1. Accept HTTP connections and return 200/OK 
2. Accept JSON containing a URL and parse:
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
4. Avoid generating short URLs for duplicate URLs
5. Accept a shortened URL and redirect to correct URL

Stretch goals:
1. Store key/values in a local JSON database
2. Count times the URL has been decoded
3. Delete URL from local database
4. Make the encoded URL "cryptographically secure"

## Running

### Test

```
go test ./goalN
```

### Run

```
go run ./goalN
```
