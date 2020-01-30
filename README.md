<div align="center">
    <h1>Golang m3u Content Parser</h1>
    <h5>Minimalistic, functional and easy to use playlist parser</h5>
</div>

### Example usage

```go
package main

import (
	"encoding/json"
	"fmt"
	gom3uparser "github.com/zikwall/gom3u-content-parser"
)

func main() {
	parser := gom3uparser.NewM3UContentParser().
		LoadSource("https://iptv-org.github.io/iptv/countries/ru.m3u", false).
		Parse()

	jsonOutput, _ := json.Marshal(parser.Offset(2).Limit(3).All())

	fmt.Println(string(jsonOutput))
}
```

### More example

```go

func main() {
	parser := gom3uparser.NewM3UContentParser().
		LoadSource("https://iptv-org.github.io/iptv/countries/ru.m3u", false).
		Parse()

	for _, item := range parser.Limit(10).All() {
		fmt.Println(fmt.Sprintf("Language is: %s, Group is: %s", item.TvgLanguage, item.GroupTitle))
	}
}

```

### Get tvg url

```go
fmt.Println(parser.GetTvgUrl())
```

### Installation

```go
go get github.com/zikwall/gom3u-content-parser
```

### Available m3u item attributes, every all string type

#### Formats
Origin `tvg-id`, `tvg-logo` transform to `TvgId` and `TvgLogo` objects

- [x] `Id`
- [x] `TvgId`
- [x] `TvgName`
- [x] `TvgUrl`
- [x] `TvgLogo`
- [x] `TvgCountry`
- [x] `TvgLanguage`
- [x] `TvgShift`
- [x] `AudioTrack`
- [x] `AudioTrackNum`
- [x] `Censored`
- [x] `GroupId`
- [x] `GroupTitle`
- [x] `ExtGrp`
- [x] `ExtraAttributes` (all original attributes in m3u item after parsing)

### Available methods & props in parser

- [x] `GetTvgUrl()` : string
- [x] `GetM3UContent()` : string
- [x] `GetDirtyItems()` : []string - Return m3u item lines before parsing
- [x] `GetItems()` : []M3UItem - Return all items after parsing
- [x] `Offset(int offset)` : M3UContentParser - Set offset for returned results
- [x] `Limit(int limit)` : M3UContentParser - Set limits for returned results
- [x] `Cache`
- [x] `Refresh`
- [x] `CountItems`
- [x] `TvgUrl`

#### Questions?

For all questions and suggestions - welcome to Issues
