# :melon: Mellang :melon:
> Mellang, an interpreted programming language

### Mellang VSCode Extension 
> You can download it on https://marketplace.visualstudio.com/items?itemName=Mello.mellang

</br>


<h3> Mellang can do almost anything that a basic programming language does: </h3>

### Types

Mellang has the following data types: `null`, `bool`, `int`, `str`, `array`,
`hash`,`fn` and a "strange" `float`

Type      | Syntax                                    | 
--------- | ----------------------------------------- | 
null      | `null`                                    |
bool      | `true false`                              |
int       | `0 24 1654 -10`                           | 
str       | `"blablabla string here"`                 | 
array     | `[] [1, 2] [1, 2, 3]`                     |
hash      | `{} {"a": 1} {"a": 1, "b": 2}`            |
float     | `1.6 1.6324 215.16`                       |


### Variable Bindings
> You can use 'moonvar (var) = (expression)' or just '(var) <- (expression)'
```
moonvar myVar = 5
anotherVar <- 5 
```

### Arithmetic Expressions
```
>> myVar <- 10
>> moonvar anotherVar <- myVar * 2
>> (myVar + anotherVar) / 2 - 3
12
```

### Conditional Expressions

Mellang supports `if` and `else`:

```
>> a <- 10
>> b <- a * 2
>> c <- if (b > a) { 99 } else { 100 }
>> c
> 99

>> a <- 10
>> b <- 10
>> c <- if (b >= a) { 99 } else { 100 }
>> c
> 99
```
## Installation

TODO

## Usage example
### Variable Declaration
```
moonvar myVariable = 5
moonvar myFunction = fn(x){x + 1}
moonvar myHash = {"name": "myName", "otherName : "randomName"}
moonvar myList = [1,3,4,5, fn(x){x + 1}]

myList[4](5)
> 6

myHash["name"]
>myName

map([1,2,3,4],fn(x){x * 2})
>[2,4,6,8]

sum(map([1,2,3,4],fn(x){x * 2}))
>20
```

### While Loops

Mellang supports only one looping construct, the `while` loop:

```
x <- 3
myList <- [1,2,3,4,5]

while (x > 0) {
 myList <- map(myList,fn(x){x * 2}) 
 plsShow(myList) 
 x <- x - 1
}

// [2,4,6,8,10]
// [4,8,12,16,20]
// [8,16,24,32,40]
```
### Functions and Closures

You can define named or anonymous functions, including functions inside
functions that reference outer variables (*closures*).

```sh
>> moonvar multiply = fn(x, y) { x * y }
>> multiply(50 / 2, 1 * 2)
50
>> fn(x) { x + 10 }(10)
20
>> moonvar newAdder = fn(x) { fn(y) { x + y } }
>> moonvar addTwo = newAdder(2)
>> addTwo(3)
5
>> moonvar sub = fn(a, b) { a - b }
>> moonvar applyFunc = fn(a, b, func) { func(a, b) }
>> applyFunc(10, 2, sub)
8
```
### Recursive Functions

Monkey also supports recursive functions including recursive functions defined
in the scope of another function (*self-recursion*).

```
>> moonvar wrapper = fn() { moonvar inner = fn(x) { if (x == 0) { return 2 } else { return inner(x - 1) } } return inner(1) }
>> wrapper()
2


>> fib <- fn(n, a, b) { if (n == 0) { return a } if (n == 1) { return b } return fib(n - 1, b, a + b) }
>> fib(35, 0, 1)
9227465
```

### Strings

```
>> makeGreeter <- fn(greeting) { fn(name) { greeting + " " + name + "!" } }
>> hello <- makeGreeter("Hello")
>> hello("mellum")
Hello mellum!
```

### Arrays

```sh
>> myArray := ["Thorsten", "Ball", 28, fn(x) { x * x }]
>> myArray[0]
Thorsten
>> myArray[4 - 2]
28
>> myArray[3](2)
4
```
### Hashes

```sh
>> myHash := {"name": "Jimmy", "age": 72, true: "yes, a boolean", 99: "correct, an integer"}
>> myHash["name"]
Jimmy
>> myHash["age"]
72
>> myHash[true]
yes, a boolean
>> myHash[99]
correct, an integer
```
### Builtin functions
- `len(iterable)`
  Returns the length of the iterable (`str`, `array` or `hash`).
- `first(iterable)`
  Returns the first element of the array.
- `last(iterable)`
  Returns the last element of the array.
- `push(iterable)`
  Add an element to the array.
- `replace(string,element_to_replace,element_to_put)`
  Replace something inside a string.
- `plsShow(element)`
  Print something in the screen.
- `Strcomp(string,string)`
  Compare the pointers of two strings.


## Development setup

After dowloading the Mellang VSCode Extension, you have to create a ".mel" file on the root of the project and specify the filename after running the "main.go".


## Meta

Antonio Mello Babo – [@MelloTonio](https://github.com/MelloTonio/)

*It was inspired by the monkey language of the book "writing an interpreter in go - thorsten ball"*



