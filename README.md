# Programming Wisdom Reference Service
An example go microService providing programming wisdom

## Usage

To dispense wisdom from command cli

```bash
> go run cmd/wisdom.go dispense

```

To Dispense wisdom from an API Server
```bash
> go run /cmd/wisdom.go serve & 

  curl https://localhost:3000/quote

  {
      "quote":"Some quote",
      "author":"author name"
  }
```

## Development
```bash
$ git clone https://github.com/pytyagi/wisdom.git
$ cd wisdom

$ make run # run the service
$ make test # run the unit tests
4 make cover # show code coverage report (browse to localhost:3000)
```