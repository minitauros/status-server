# Status server

Small server that will reply with the status code you tell it to. Useful for locally testing if your application responds correctly to different response status codes.

## Usage

```
go run main.go
````

Then navigate to http://localhost:8080/{status_code}, for example http://localhost:8080/404.

Optionally a JSON response body can be specified. For example loading http://localhost:8080/404?a=b will respond with `{"a":"b"}`.