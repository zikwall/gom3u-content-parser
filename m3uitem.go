package gom3u_content_parser

import (
	"reflect"
	"strings"
)

var availableAttributes = []string{
	"id", "tvg-id", "group_id", "group-title", "tvg-shift",
	"tvg-name", "tvg-logo", "audio-track", "audio-track-num",
	"censored", "tvg-country", "tvg-language", "tvg-url",
}

type M3UItem struct {
	Id              string            `json:"id"`
	TvgId           string            `json:"tvg_id"`
	TvgName         string            `json:"tvg_name"`
	TvgUrl          string            `json:"tvg_url"`
	TvgLogo         string            `json:"tvg_logo"`
	TvgCountry      string            `json:"tvg_country"`
	TvgLanguage     string            `json:"tvg_language"`
	AudioTrack      string            `json:"audio_track"`
	AudioTrackNum   string            `json:"audio_track_num"`
	TvgShift        string            `json:"tvg_shift"`
	Censored        string            `json:"censored"`
	GroupId         string            `json:"group_id"`
	GroupTitle      string            `json:"group_title"`
	ExtGrp          string            `json:"ext_grp"`
	ExtraAttributes map[string]string `json:"extra_attributes"`
}

func setM3UItemField(m3uitem *M3UItem, field string, value string) {
	v := reflect.ValueOf(m3uitem).Elem().FieldByName(field)
	if v.IsValid() {
		v.SetString(value)
	}
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

		for _, attr := range availableAttributes {
			if _, ok := m3uitem.ExtraAttributes[attr]; ok {
				structFiledName := ucFirst(Camelize(attr))
				setM3UItemField(m3uitem, structFiledName, m3uitem.ExtraAttributes[attr])
			}
		}
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
