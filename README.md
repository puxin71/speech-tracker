## Synopsis

This application is a simple REST API server that exposes endpoints to allow accessing and manipulating the "talk" database records. The operations that the endpoint will allow include:
* Fetching all talks and their attendants
* Fetching one talk and its attendees

## API Reference
The APIs should:
* Fetch all talks and their attendants in response to a valid GET request at /talks
* Fetch all attendees that are registered to a talk in response to a valid GET request at /talks/[0-9]+/attendees

## Code Example

Show what the library does as concisely as possible, developers should be able to figure out **how** your project solves their problem by looking at the code example. Make sure the API you are showing off is obvious, and that your code is short and concise.

## Motivation

A short description of the motivation behind the creation and maintenance of the project. This should explain **why** the project exists.

## Installation

For the short term, you have to fetch the sour code from the [Github repo] (https://github.com/puxin71/speech-tracker) and run it as:
```
RESOURCES="./resources" go run server.go
```

Eventually we can launch the server using Docker
## Tests

Describe and show how to run the tests with code examples.

## Contributors

Let people know how they can dive into the project, include important links to things like issue trackers, irc, twitter accounts if applicable.

## License

A short snippet describing the license (MIT, Apache, etc.)
