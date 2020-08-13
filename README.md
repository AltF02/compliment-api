# Compliment API
This api is build in go and uses the mux library, it returns a random compliment from the database 
## How to use

### Python
```python
import requests
resp = requests.get("")
print(resp.json()['compliment']) # Prints the compliment
```