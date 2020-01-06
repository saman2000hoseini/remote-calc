# Remote Calculator
trying to implement simple calculator using client and server
# Run
````
go run ./cmd/server/main.go
go run ./cmd/client/main.go
````
# Usage
first run server and then client

you can send operations with 2 operands and 1 operator in following pattern:

operand1

operator

operand2

**example:**
````
1.5
+
2
````
supported operators : {+,-,*,/}
