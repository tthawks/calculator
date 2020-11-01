# calculator

Calculator is simple web application which uses get request on predefined path with query parameters as input data to do simple math operations. Application implements caching of results for one minute.

Implemented mathematical operations are add, subtract, multiply, and divide.

Results are returned as JSON object with keys:
action - math operation used
x, y - input values
answer - result of used operation 
cached - is result from cache (or is it calculated)

## Implemented math operations
### add

Sumarizes input values

Example

Request:
```http://localhost:3000/add?x=20&y=7```

Response: 
```{"action":"add","x":20,"y":7,"answer":27,"cached":false}```

### subtract

Subtract y value from x value

Example

Request:
```http://localhost:3000/subtract?x=20&y=7```

Response: 
```{"action":"subtract","x":20,"y":7,"answer":13,"cached":false}```

### multiply

Multiplies input values

Example

Request:
```http://localhost:3000/multiply?x=20&y=7```

Response: 
```{"action":"multiply","x":20,"y":7,"answer":140,"cached":false}```


### divide

Divides x value with y value

Example

Request:
```http://localhost:3000/divide?x=20&y=7```

Response: 
```{"action":"divide","x":20,"y":7,"answer":2.857142857142857,"cached":false}```



## Installation

`go get github.com/tthawks/calculator`