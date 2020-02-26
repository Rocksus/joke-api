# Joke-Api
An API for any jokes written in Go.

# Usage
## Endpoints
### Random Joke API
[https://horrible-jokes.appspot.com/jokes](https://horrible-jokes.appspot.com/jokes)
### Random Joke API by Category
`https://horrible-jokes.appspot.com/jokes/{category}`

Example:
[https://horrible-jokes.appspot.com/jokes/programming](https://horrible-jokes.appspot.com/jokes/programming)
### Joke API by ID
`https://horrible-jokes.appspot.com/joke/{id}`

Example:
[https://horrible-jokes.appspot.com/joke/1](https://horrible-jokes.appspot.com/joke/1)

# Contribute!
Make a pull request on data/jokes.json file in the following format:
```javascript
{
  "id": last_joke_id + 1,
  "category": "animals",
  "setup": "Why do the French eat snails?",
  "punchline": "They don't like fast food.",
  "lang": "en"
}
```

## Usage
To run locally, just do

`go mod download`

and then

`go run main.go`

you can open locally from (default port is 8000, change it in your .env)

`localhost:8000/jokes`