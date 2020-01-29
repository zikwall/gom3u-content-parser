<div align="center">
    <h1>Golang m3u Content Parser</h1>
    <h5>Minimalistic, functional and easy to use playlist parser</h5>
</div>

### Example usage

```go
func main() {
	parser := p.NewM3UContentParser().
		LoadSource("https://iptv-org.github.io/iptv/countries/ru.m3u", false).
		Parse()

	jsonOutput, _ := json.Marshal(parser.Offset(2).Limit(3).All())

	fmt.Println(string(jsonOutput))
}
```

### Installation

```go
go get github.com/zikwall/gom3u-content-parser
```

### API (todo)

Get attributes