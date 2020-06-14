## Usage
### Dependencies
-  [Install go](https://golang.org/doc/install)

### Run JSON Flattener
1. Clone the repo
```sh
git clone https://github.com/ndhanushkodi/jsonflattener.git
cd jsonflattener
```

2. Compile `jsonflattener`
```sh
go build .
```

3. `jsonflattener` accepts JSON input on `stdin`. To run it with the provided example file, run:
```sh
cat examplefile | ./jsonflattener 
```

## Tests
```sh
cd jsonflattener
go test ./flatten
```