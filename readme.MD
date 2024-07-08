# OOP in GO lang

OOP is a programming paradigm based on the concept of object , which can contain data and code : data in the form of fields (often known as attribute or properties) and code in the form of procedures (often known as methods)

Go doesn't have a class like in object-oriented language , it supports some oo programming concepts through its struct types and struct

```go
//define    a struct
type Triangle struct {
    Base , Height float64;
}
func (t Triangle) Area()float64{
    return 0.5*t.Base*t.Height
}
```

## Encapsulation :

It is a practice to hide implementation details of one object form another object. It is a process which encourages decoupling

for example :

```go
type Student struct{
    name string
    rollNo unit64
}

```

name and rollNo are private because they start with lowercase letters.
To provide public access you can define getter and setter methods for the respective fields

```go
func(s *Student) GetName()string{
    return s.name
}

func (s *Student) SetName(name string){
    s.name = name
}

func (s *Student) GetRollNo() unit64{
    return s.rollNo
}

func (s *Student) SetRollNo(roll unit64){
    s.rollNo =  roll;
}

```

## Abstraction

Data abstraction is a design pattern in which are visble only to semantically related functions , so as to prevent misuse. The success of data abstraction leads to frequent incorporation of data hiding as a desing principles in object oriented and pure funcitional programming

In go it can be achieved using interfaces. An interface defines a set of method that that type must implement in order to satisfy the interface

```go
type iStrudent interface{
    calcPercentage() float64
}

```

## Inheritance

Inheritance is not supported in go , but go uses composition to achieve similar funcitonality. A struct can embade another struct , allows it to inherit fields of methods

```go

type Person struct {
    name string
    age unit64
}
type Student  struct {
    Person
    studentId unit64
}

```

the struct embades the person stuct, whcih means it inherits its name and age proerties and also Student have its own studentId field on its own => Student is an Person with some additional fields.

## Polimorphism :

It is when calling code can be independent of whcih class in the supported hierarchy it is operating on the prretn class or one of its decendets

# GO Types

Go is type safe , statically typed , compiled programming language . The type system with types denoted by type names and type declaration is designed to prevent occurences of unchecked reuntime tyep errors.

In go There are several built in tyeps for identifiers known as predeclared types , they includes Boolean , stirngs , numberic (floats32, float64, int , int8, int16, int32, complex ) and many more other types
Addidionaly there are compsite types which are composed from predeclared tyeps

## Compsite types

Are mainly constructed using type literals. They include arrays , slices , interfaces , structs ,functions , maps and many more

### Pointers in go

A pointer type is reference type that demotes the set of all pointers to variables of a given type. Generally pointers type holds a memory adress of another variables . The zero value of pointer is nil.

## Structs

Structs contains a fields of the same or d/t types. Structs are basically collection of named fields that have alogical meaning or contruct , where in each field has a specific type.

Struct types are combintion of user-defined types . They are specialized types because they allow us to define cusotom data types in such case where the built in types are not sufficient

```go
type  blogPost struct {
    author string
    title string
    postId int
}

```

### Nested or embeded struct fields

We can create a nested or embaded type incide other structs .

```go
type blogPost struct {
    title string
    postId int
    published bool
}

type Author struct {
    firstName , lastName , Biography string
    photId int
}

```

#### things to concider

- Filed name must be a unique incide a struct type
- A Field or a mthod of an embeding type can be promoted
- Promoted fields can not be used as a field names in a struct
- A field declaration can be followed by an optional stirng literal tag
- An exported struct field must began with capital letter
- Apart from basic types , we can also have a function types and interface types as struct fields

# SOLID Prenciples In go

## Single Responsiblity Prenciple(SRP)

States that a class shlould have only one reason to change. If we violate , the class will have multiple responsiblity , making it harder to to maintain , test , and extend. This can lead to code that is tightly coupled , dificult to reuse , and prone to error

## OpenClose Prenciple (OCP)

This prenciple states that a classes should be open for exetentiosion but closed for modeification . If we violate this prenciple we may have to modify existing code to add new functionality , which can introduce bugs and make it difficult to maintain the code. This can be achieved by defining an interface for processing trades and implementing it for each trade type.

## Liskov's Substitution Prenciple (LSP)

States that subtypes should ebe substitutable for their base types.If we violate this prenciple we may introduce beheviour that is unexpected and inconsitent which can lead to errors that are difficult to track down.
This can aloso make it difficult to write code that works with a viriety of d/t types

## Interface segregation Prenciples (ISP)

States that clients should not be forced to depend on interfaces they do not use. IF we violate this prenciple we may have interface that are too large and contains methods that are not relevant to some some client , which can lead to code that is difficult to understand and maintain .

This can alos result in code that is not reusable , and that can cause unnecessary coupling b/n modules.

## Dependency Inversion Prenciples (DIP)

States that high-level modules should not depend on low-level modules. Instead , both should depend on abstractions. If we violate prenciples we may have a code that is difficult to maintain and extend.
