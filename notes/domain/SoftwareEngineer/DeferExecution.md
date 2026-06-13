# Defer Execution

Taks `ProperptyChain` proejct as example

A property file need to open file and parse each line from it as a property

## straightforward

Do open file and parse in property constructor

## decouple source loader

Open file in source loader and pass to type of property

ex: get source file from http and pass io.reader to property contructor

Cons:

- The source loader can't close io.reader when source loader finish the function scope, because it has to pass to property contructor. Potential memory leak.
- If read file as []byte, memory is required ahead

## decouple and defer source loader

Called `Visitor Pattern` or `Double Dispatch`

Create a source loader function and pass it to property constructor

When constructor read from the io.reader(inside function), in the end of function, it also close the file.
It also defers the file opening when property constructor needs it

Pros:

- No memory ahead
- No potential file opening leak

Patterns:

- factory factory to generate property handler
- strategy patten in each handler
- visitor pattern to load file

## Enhance defer feature

Let file close do in async backgroup

Pros: Performance tuning
Cons: Too complex
