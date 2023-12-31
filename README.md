# REST Example in GO

Use `go get .` to get dependencies for code in the current directory.

## Run the code

Use `go run .` to running HTTP server to which you can send requests.

## Testing

From a new command line window, use curl to make a request to your running web service.

### GET Albums

``` shell
curl http://localhost:8080/albums
``` 

Response:

``` json 
[
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        }
]
```

### GET Albums By ID

``` shell
curl http://localhost:8080/albums/2
``` 

Response:

``` json 
{
	"id": "2",
	"title": "Jeru",
	"artist": "Gerry Mulligan",
	"price": 17.99
}
```

### POST Albums

``` shell
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
``` 

Response:

``` json 
{
	"id": "4",
	"title": "The Modern Sound of Betty Carter",
	"artist": "Betty Carter",
	"price": 49.99
}
```
