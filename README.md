# lambda-request-parser

- Convert the lambda request (QueryStringParameters or RequestBody) to a go struct.
- Provides easy validation with tags. (use [go-playground/validator](https://github.com/go-playground/validator))

## Installation

Use go get.

```
go get -u github.com/r-konishi/lambda-request-parser/parser
```

Then import the parser package into your own code.

```
import "github.com/r-konishi/lambda-request-parser/parser"
```

## Usage

[test code](./parser/parser_test.go)

exapmles

[simple lambda example (with api gateway)](./examples/simple/)

## License

Distributed under MIT License, please see license file within the code for more details.
