<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# db

```go
import "Lescatit/db"
```

## Index

- [type Config](<#type-config>)
  - [func NewConfig() Config](<#func-newconfig>)
- [type Connection](<#type-connection>)
  - [func NewConnection(cfg Config) (Connection, error)](<#func-newconnection>)


## type [Config](<https://github.com/mtnmunuklu/Lescatit/blob/main/db/config.go#L10-L13>)

```go
type Config interface {
    Dsn() string
    DbName() string
}
```

### func [NewConfig](<https://github.com/mtnmunuklu/Lescatit/blob/main/db/config.go#L24>)

```go
func NewConfig() Config
```

## type [Connection](<https://github.com/mtnmunuklu/Lescatit/blob/main/db/db.go#L9-L12>)

```go
type Connection interface {
    Close()
    DB() *mgo.Database
}
```

### func [NewConnection](<https://github.com/mtnmunuklu/Lescatit/blob/main/db/db.go#L19>)

```go
func NewConnection(cfg Config) (Connection, error)
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
