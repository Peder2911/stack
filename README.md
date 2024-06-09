
# Stack

Stack is a very minimal project that I use for keeping track of things that I want to remember and keep track of. Stack is a FIFO stack.


## Installation

```
go build -o $HOME/bin/stack cmd/server.go
```

## Running

```
stack
```

## Usage

```
# Push some data onto the stack
curl 0.0.0.0:8000/push --data "Some data"

# Show number of items in stack
curl 0.0.0.0:8000/

# Get an item from the top of the stack
curl 0.0.0.0/pop

# Clear the stack
curl -XDELETE 0.0.0.0/
```
