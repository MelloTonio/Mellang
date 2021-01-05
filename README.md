# :melon: Mellang :melon:
> Mellang, an interpreted programming language

We can do almos anything that a basic programming language does:
## Mellang has:
* Integers
* Strings
* Booleans
* Float
* Variables
* Functions
* Hashes
* Lists
* Built-in Functions (map, reduce, sum, replace...)


## Installation

TODO

## Usage example
### Declarating variable
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

## Development setup

TODO

## Meta

Antonio Mello Babo – [@MelloTonio](https://github.com/MelloTonio/)

*It was inspired by the monkey language of the book "writing an interpreter in go - thorsten ball"*



