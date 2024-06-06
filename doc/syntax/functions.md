## functions

if just writing commmands in .mln file like you would in mcf

```go
time set day
```

these commands will go to the load function in your namespace (`test:load`). to create your own function use the `func` keyword. call it tick to make it run every game-tick

```go
// test:make_day
func make_day() {
  time set day
}

// this is the tick function.
func tick() {
  execute as @a run effect give @s speed 1 100 true
}
```

to create function under a specific directory, like `test:uwu/owo` we could use one of the following methods

1. dots in filename, affects the single defined function

```go
func uwu.owo() {
  time set day
}
```

2. `dir` block, affects code within the block

```go
dir uwu {
  func owo() {
    time set day
  }
}
```

3. `dir` keyword, affects code below the keyword

```go
dir uwu

func owo() {
  time set day
}
```
