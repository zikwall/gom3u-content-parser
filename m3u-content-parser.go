package gom3u_content_parser

import (
	"regexp"
)

type M3UContentParser struct {
	m3uFileContent string
	dirtyItems     []string
	items          []M3UItem
	countItems     int

	TvgUrl  string
	cache   int
	refresh int

	offsets int
	limits  int
}

func NewM3UContentParser() *M3UContentParser {
	return &M3UContentParser{}
}

func (parser *M3UContentParser) LoadSource(source string, fromFile bool) *M3UContentParser {
	if fromFile {
		parser.m3uFileContent = ReadStringContentFromFile(source)
	} else {
		parser.m3uFileContent = ReadStringContentFromRemote(source)
	}

	return parser
}

func (parser *M3UContentParser) Parse() *M3UContentParser {

	a := regexp.MustCompile(`(#EXTINF:0|#EXTINF:-1|#EXTINF:-1,)`)
	parser.dirtyItems = a.Split(parser.m3uFileContent, -1)

	// first is #EXTM3U tag
	parser.parseAndSetTvgUrl(parser.dirtyItems[0])
	parser.dirtyItems = parser.dirtyItems[1:]

	for _, item := range parser.dirtyItems {
		parser.items = append(parser.items, *NewM3UItem(item))
		parser.countItems++
	}

	return parser
}

func (parser *M3UContentParser) parseAndSetTvgUrl(url string) *M3UContentParser {
	re := regexp.MustCompile(`"([^"]+)"`)
	matches := re.FindAllString(url, -1)

	// if tv program exist
	if len(matches) > 0 {
		parser.TvgUrl = matches[0]
	}

	return parser
}

func (parser *M3UContentParser) GetTvgUrl() string {
	return parser.TvgUrl
}

func (parser *M3UContentParser) GetM3UContent() string {
	return parser.m3uFileContent
}

func (parser *M3UContentParser) GetDirtyItems() []string {
	return parser.dirtyItems
}

func (parser *M3UContentParser) GetItems() []M3UItem {
	return parser.items
}

func (parser *M3UContentParser) Offset(offset int) *M3UContentParser {
	parser.offsets = offset

	return parser
}

func (parser *M3UContentParser) Limit(limit int) *M3UContentParser {
	parser.limits = limit

	return parser
}

func (parser *M3UContentParser) All() []M3UItem {
	if parser.limits <= 0 {
		parser.limits = parser.countItems
	}

	return parser.items[parser.offsets:parser.limits]
}
