# toml-cli

## Build
```
go build
```

## Usage
```
Usage: toml CMD FILE [QUERY] [VALUE]

Commands:
lint output linted file
get  get query result
set  set value in toml file and save
```

### Examples
All examples assume a toml file as follow
```
[some]
[some.where]
key="value"
[some.time]
key="value"
```

#### Get a single value
```
$ toml-cli get file.toml some.where.key
value
```

#### Get whole file
```
$ toml-cli get file.toml
[some]

  [some.where]
    key = "value"

  [some.time]
    "key.subkey" = "value"
```
> NOTE: this also lint the file

#### Setting value
```
$ toml-cli set file.toml some.where.key newvalue
$ toml-cli set file.toml "some.where.'key.subkey'" newvalue
```
> NOTE: This modify the file in place

#### Lint
```
$ toml-cli lint file.toml
```
> NOTE: This modify the file in place
