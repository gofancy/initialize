# `initialize` runs initializers of a Go application

Suppose your application has many dependencies. When it starts, it will customize
log output and format, set up connections to database and Redis, and connect to
Elasticsearch. How can you implement it?

Maybe you write an `init` function, which has a thousand lines of code.

```golang
func init() {
    // customize log
    ...
    ...
    ...

    // connect to database
    ...
    ...
    ...

    // connect to Redis
    ...
    ...
    ...
    ...
    ...
}
```

Or better, you divide it into a bundle of smaller functions.

```golang
func customizeLog(){
    ...
}

func connectToDB(){
    ...
}

func connectToRedis(){
    ...
}

...
...

func init() {
    customizeLog()
    connectToDB()
    connectToRedis()
    ...
    ...
}
```

With this library, the `init` function can be simplified to one line of code at
the cost of making each initializer function a little more complex.

```golang
type initializer struct{}

func (v initializer) Initialize01Log(){
    ...
}

func (v initializer) Initialize02DB(){
    ...
}

func (v initializer) Initialize03Redis(){
    ...
}

...
...

func init() {
	initialize.AllFrom(initializer{})
}
```

There are several advantages:

* No need to touch `init` function to add / remove initializers
* The order of initializers is defined by numbers
* The file structure can be clearer, like this:

```text
  ▾ initializers/
      00_log.go
      01_db.go
      02_redis.go
      03_elasticsearch.go
      init.go
```

## Get it

```
$ go get github.com/gofancy/initialize
```
