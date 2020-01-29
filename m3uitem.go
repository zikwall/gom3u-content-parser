package gom3u_content_parser

import "strings"

var availableAttributes = []string{
	"id", "tvg-id", "group_id", "group-title", "tvg-shift",
	"tvg-name", "tvg-logo", "audio-track", "audio-track-num",
	"censored", "tvg-country", "tvg-language", "tvg-url",
}

type M3UItem struct {
	Id              string `json:"id"`
	TvgId           string
	TvgName         string
	TvgUrl          string `json:"tvg_url"`
	TvgLogo         string
	TvgCountry      string
	TvgLanguage     string
	AudioTrack      string
	AudioTrackNum   int
	TvgShift        int
	Censored        int
	GroupId         int
	GroupTitle      string
	ExtGrp          string
	ExtraAttributes map[string]string `json:"extra_attributes"`
}

func NewM3UItem(item string) *M3UItem {
	m3uitem := new(M3UItem)
	result := strings.Split(strings.Replace(item, "\r\n", "\n", -1), "\n")

	m3uitem.TvgName = result[0]
	m3uitem.TvgUrl = result[1]

	if strings.Index(m3uitem.TvgName, ",") != -1 {
		ex := strings.Split(m3uitem.TvgName, ",")
		m3uitem.TvgName = ex[1]

		m3uitem.ExtraAttributes = ParseAttributes(ex[0])
	}

	if len(result) > 2 && result[2] != "" {
		if strings.Index(result[1], "#EXTGRP") != -1 {
			groupName := strings.Split(result[1], ":")
			m3uitem.ExtGrp = groupName[1]
		}

		m3uitem.TvgUrl = result[2]
	}

	return m3uitem
}

func (m3uitem *M3UItem) GetExtraAttributes() map[string]string {
	return m3uitem.ExtraAttributes
}
