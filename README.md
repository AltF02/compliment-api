# Compliment API
This api is build in go and uses the mux library, it returns a random compliment from the database 
## How to use

### Go
```go
import (
    "net/http"
)

resp, err := http.get("")
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(resp)) // Prints the received compliment
```


### Python
```python
import requests
resp = requests.get("")
print(resp.body) # Prints the compliment
```

### JavaScript
```javascript
const Http = new XMLHttpRequest();
const url='';
Http.open("GET", url);
Http.send();

Http.onreadystatechange = (e) => {
  console.log(Http.responseText) // Prints the compliment
}
```