# MyFizzBuzz HTTP Service

### Prerequisite

To compile this projetc, you need:

```sh
$ go install myfizzbuzz
```

### Usage

To run:

```sh
$ ./myfizzbuzz
```

or you can use (without install):

```sh
$ go run myfizzbuzz
```

### Tests

To execute fizzbuzz by HTTP Request, go to :
- http://localhost:8080/?string1=fizz&string2=buzz&int1=3&int2=5&limit=100

### required fields

- string1
- string2
- int1 (>0)
- int2 (>0)
- limit (>0)
