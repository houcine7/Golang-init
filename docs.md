# Important Note

In Go, we can't declare a variable without using it; the compiler will throw an error.

## Data Types for Integers

- `var num int` or `var num = 5`
- `int` defines integer values coded in 32 bits or 64 bits depending on the system.
- `int8`, `int16`, `int32`, `int64` are different types of `int`.
  - `int8` = 8 bits, `int16` = 16 bits, `int32` = 32 bits, `int64` = 64 bits.
- `int16` can store values from -32768 to 32767 (2^16 numbers).

### Unsigned Integer Types

- Store only positive numbers: `uint`, `uint8`, `uint16`, `uint32`, `uint64`.
  - For example, `uint8` can store values from 0 to 255 (2^8 numbers).

## Float Types

There are two float types: `float32` and `float64`.

## Arithmetic Operations

We cannot add two different data types, e.g., int and int8 or int and float32, but we can use casting to convert one type to another.

## Strings

- Strings are stored as a slice of bytes and are immutable.
- We can't change a string once it is created.
- We can create a string using double quotes or backticks (` `).
  1. For example:
     - `var name string = "John"` (Can't be formatted in multiple lines but can use escape characters like `\n`).
     - `` var name string = `John`  `` (Can be formatted in multiple lines).
  2. `var name = "John"` (Type inference. dynamically assigns the type to the variable).
  3. `name := "John"` (Short-hand declaration).
- The length of the string: `len(name)` (returns the number of bytes in the string).
- To get the number of characters in the string, use: `utf8.RuneCountInString(name)`.
- utf8 from the package unicode/utf8
  > Strings under the hood are slices of bytes represented by the type []byte each character(Rune) is represented by one or more byte

## Runes

- Runes are used to represent a character in Go.
- A rune is an alias for int32 and is equivalent to int32 in all ways.
- The number of bits a rune uses to represent a character depends on the type of the character (ASCII uses 8 bits, Unicode uses 32 bits).
- To get the rune of a character: `rune('a')` (returns 97).
- To get the character of a rune: `string(97)` (returns 'a').

- Declaring a rune:
  - `var r rune = 'a'` (cannot use double quotes).
  - `var r rune = 97` (using ASCII code, range 0 to 127).

## Booleans

- Nothing extraordinary:
  - `var myBoolean bool = true` or `false`.
  - `var myBoolean = true` (type inference).
  - `myBoolean := true` (short-hand declaration).

## Variables Declaration

- Multiple variables in one line:
  - `var v1, v2 = 5, 6` or `var v1, v2 int = 5, 6` or `v1, v2 := 5, 6`.

## Constants

- What applies to variables applies to constants, but the value of a const can't be changed once it is declared.

## Default Values of Variables

- `int*`, `uint*`, `float*` have 0 as the default value.
- `string` has ""(empty string) as the default value.
- `bool` has false as the default value.

## Functions

- Defined using the `func` keyword.
- Return values:
  - `func functionName() int` (one return value)
  - `func functionName() (int, int)` (multiple return values).
- Parameters: `func functionName(param1 int, param2 int)`.
- Multiple return values and parameter examples provided.
  ```
  func functionName(param1 int, param2 string) (int,int){
    // function code
    }
  ```

## Arrays

- Hold a fixed number of values of the same type.
- Various ways to declare and initialize arrays shown:

```
    var arr [5]int // this array can hold 5 integers
    var arr = [5]int{1,2,3,4,5} // this array can hold 5 integers and it is initialized with the values 1,2,3,4,5
    var arr = [...]int{1,2,3,4,5} // this array can hold 5 integers and it is initialized with the values 1,2,3,4,5
    var arr = [5]int{1,2,3} // this array can hold 5 integers and it is initialized with the values 1,2,3,0,0
    var arr = [5]int{1:2,3:4} // this array can hold 5 integers and it is initialized with the values 0,2,0,4,0 (2 at index 1 and 4 at index 3)
    var arr = [...]int{1:2,3:4} // this array can hold 5 integers and it is initialized with the values 0,2,0,4,0
```

- Random access to elements, e.g., `arr[0]`
- To get the length of the array `len(arr)`.

## Slices

- Dynamic arrays, using arrays under the hood.
- `var slice []int` declared like an array but without a size.

## Maps

- Store key-value pairs.
- Declaring a Map:

  - `var map map[string]int32 = map[string]int32{"hello":12,"world": 13}`
    > Here we declared a map and instantiate it at the same time
  - using make: var map = `make(map[string]int32)`
  - Accessing element:

    - map["hello"] -> if it exist it returns the value if not it returns the default value of the type

      ```
      //but we can check if the value exists by
      var val,exist = map["hello"] // exist is boolean

      ```

  - To delete element from map: `delete(map,"hello")`

## Structs

- A collection of fields (properties) to group related data.
- Example of struct declaration and initialization.
  ```
    type user struct{
    name string
    age int
    }
  ```
- used to group related data together and create ad define a new type
- instantiation (like any other type): `var u user = user{name:"my-name",age:"my-age"}`
- access fields: `u.name`

## Pointers

- Used to store the address of a variable.
- Declaration and usage of pointers:
  - var ptr \*int32
    > A pointer to an int32 variable but for now it doesn't point to a variable (will allocate a memory location in the stack memory to hold the address but it will not point to anything)

## Concurrency

- Concurrency vs. parallelism explained.

  Concurrency is the ability to run multiple programs at the same time on the same core of the CPU

  Concurrency is not parallelism: parallelism is the ability to run multiple programs at the same time on different cores of the CPU (like multi threading in java)

  > so a parallelism is a concurrency but not all concurrency is parallelism

- Introduction to goroutines for concurrent execution.
  Goroutines are functions that run concurrently with other functions
  To create a goroutine we use the keyword go this will run the function in a new goroutine

  ```
      go myFunction() // this will run the function myFunction in a new goroutine
      time.Sleep()
  ```

  > the function runs in a new goroutine and the main function continues to run if the main function ends the program will end even if the goroutine is still running

  #### Here comes wait groups to handel this

- Use of wait groups to synchronize goroutines.

  - wait groups are used to wait for the goroutines to finish before the main function ends
  - for this we use the package sync

  ```
      var wg sync.WaitGroup // create a new wait group
      wg.Add(1) // add 1 to the wait group
      wg.Done() // remove 1 from the wait group
      wg.Wait() // wait for the wait group to be 0
  ```

  > this concurrent execution can introduce an other concept
  > data corruption: when multiple goroutines try to access the same variable at the same time
  > this gives inconsistent results because the goroutines are not synchronized

#### Addressing data corruption issues with mutexes (sync.Mutex and sync.RWMutex).

- mutexes are general concept in programming not only in go it is used to synchronize access to a shard resource (a variable in our case)
- only one goroutine can access the variable at a same time
- the other goroutines will wait until the goroutine that is using the variable finishes and frees the variable

```
    var mutex sync.Mutex // create a new mutex
    mutex.Lock() // lock the mutex
    resource++ // do something with the resource
    mutex.Unlock() // unlock the mutex
```

> Like if we wrap the shared resource with the mutex the other goroutines will wait until the mutex is unlocked and then they will be able to lock the mutex and use the resource

- there also the possibility to lock the Read and Write operations on the resource

  ```
      var rwMutex sync.RWMutex // create a new read write mutex
      rwMutex.RLock() // lock the write only
      resource++ // do something with the resource
      rwMutex.RUnlock() // unlock the write only

  ```

  > this will allow multiple goroutines to read the resource at the same time

```
   var rwMutex sync.RWMutex // create a new read write mutex
   rwMutex.Lock() // lock the write mutex
   resource++ // do something with the resource
   rwMutex.Unlock() // unlock the write mutex
```

> this will allow only one goroutine to write to the resource at a time but it will not allow any goroutine to read the resource until the write mutex is unlocked

## Channels

- channels are a way to communicate and synchronize between goroutine
- Channels provide a way for goroutines to send and receive data safely allowing for the coordination of concurrent task.
- Channels:
  - holds data
  - thread safe
  - listens for data
