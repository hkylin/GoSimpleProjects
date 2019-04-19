# UUID-renamer
Rename files into UUID. Use Go(Golang).
## Use
Install the **UUID** package by using the go command:
```bash
$ go get github.com/satori/go.uuid
```
Download the file `UUID-renamer.go`, and go to the path where it puts.  
Use the go command:
```bash
$ go run ./UUID-renamer.go -p [path]
```
## Explanation
```
-p [path](string) 
    path: The path of your folder which you want to rename all the files under it into UUID file name.
```

## Example
```bash
$ go run ./UUID-renamer.go -p ~/folder/
```
## Notice
Please add `/`(Linux or others) or `\`(Windows or others) to the end of the path, if it doesn't exist.
## License
[MIT License](LICENSE)
