# Zero Width Encode and Decode
Encode Content to Zero Width Characters or Decode Zero Width Characters to the Raw Content. Use Go(Golang).
## Use
Download [ZeroWidth.go](ZeroWidth.go), then use the go command:
```bash
$ go run ./Zerowidth.go -d [encoded content] -e [raw content] -b [content before] -a [content after]
```

Or, you can use it online: ***[I'm the link!](https://play.golang.org/p/x80M6r70kef)***
## Explanation:
```
  -d [encoded content](string)
        what you want to decode
  -e [raw content](string)
        what you want to encode
  -b [content before](string)
        what you want to put before encoded content (default "encoded content ->")
  -a [content after](string)
        what you want to put before encoded content (default "<- encoded content")
```
## Example
### Encode
```bash
$ go run ./ZeroWidth.go -e "Hello, World. 你好。"
Encoded:
encoded content ->‌​‍‍​‍‍‍‌​​‍‍​‍​‌​​‍​​‍‍‌​​‍​​‍‍‌​​‍​​​​‌​‍​​‍‍‌​‍‍‍‍‍‌​‍​‍​​​‌​​‍​​​​‌​​​‍‍​‍‌​​‍​​‍‍‌​​‍‍​‍‍‌​‍​​​‍‌​‍‍‍‍‍‌​‍‍​​​​‍​​‍‍‍‍‍‌​‍​​‍‍​‍​​​​​‍​‌​​‍‍‍‍‍‍‍‍‍‍​‍‌<- encoded content

$ go run ./ZeroWidth.go -e "Hello, World. 你好。" -b "hel" -a "lo"
Encoded:
hel‌​‍‍​‍‍‍‌​​‍‍​‍​‌​​‍​​‍‍‌​​‍​​‍‍‌​​‍​​​​‌​‍​​‍‍‌​‍‍‍‍‍‌​‍​‍​​​‌​​‍​​​​‌​​​‍‍​‍‌​​‍​​‍‍‌​​‍‍​‍‍‌​‍​​​‍‌​‍‍‍‍‍‌​‍‍​​​​‍​​‍‍‍‍‍‌​‍​​‍‍​‍​​​​​‍​‌​​‍‍‍‍‍‍‍‍‍‍​‍‌lo
```
### Decode
```bash
$ go run ./ZeroWidth.go -d "‌​‍‍​‍‍‍‌​​‍‍​‍​‌​​‍​​‍‍‌​​‍​​‍‍‌​​‍​​​​‌​‍​​‍‍‌​‍‍‍‍‍‌​‍​‍​​​‌​​‍​​​​‌​​​‍‍​‍‌​​‍​​‍‍‌​​‍‍​‍‍‌​‍​​​‍‌​‍‍‍‍‍‌​‍‍​​​​‍​​‍‍‍‍‍‌​‍​​‍‍​‍​​​​​‍​‌​​‍‍‍‍‍‍‍‍‍‍​‍‌"
Decoded:
Hello, World. 你好。

$ go run ./ZeroWidth.go -d "encoded content ->‌​‍‍​‍‍‍‌​​‍‍​‍​‌​​‍​​‍‍‌​​‍​​‍‍‌​​‍​​​​‌​‍​​‍‍‌​‍‍‍‍‍‌​‍​‍​​​‌​​‍​​​​‌​​​‍‍​‍‌​​‍​​‍‍‌​​‍‍​‍‍‌​‍​​​‍‌​‍‍‍‍‍‌​‍‍​​​​‍​​‍‍‍‍‍‌​‍​​‍‍​‍​​​​​‍​‌​​‍‍‍‍‍‍‍‍‍‍​‍‌<- encoded content"
Decoded:
Hello, World. 你好。

$ go run ./ZeroWidth.go -d "hel‌​‍‍​‍‍‍‌​​‍‍​‍​‌​​‍​​‍‍‌​​‍​​‍‍‌​​‍​​​​‌​‍​​‍‍‌​‍‍‍‍‍‌​‍​‍​​​‌​​‍​​​​‌​​​‍‍​‍‌​​‍​​‍‍‌​​‍‍​‍‍‌​‍​​​‍‌​‍‍‍‍‍‌​‍‍​​​​‍​​‍‍‍‍‍‌​‍​​‍‍​‍​​​​​‍​‌​​‍‍‍‍‍‍‍‍‍‍​‍‌lo"
Decoded:
Hello, World. 你好。
```
## Notice
UTF-8 Support! Maybe :)

## License
[MIT License](LICENSE)
