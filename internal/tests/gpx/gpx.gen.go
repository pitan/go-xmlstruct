// This file is automatically generated. DO NOT EDIT.

package gpx

import "time"

type GPX struct {
	ActivePoint *struct {
		Lat float64 `xml:"lat,attr"`
		Lon float64 `xml:"lon,attr"`
	} `xml:"active_point"`
	Author *string `xml:"author"`
	Bounds *struct {
		MaxLat float64 `xml:"maxlat,attr"`
		MaxLon float64 `xml:"maxlon,attr"`
		MinLat float64 `xml:"minlat,attr"`
		MinLon float64 `xml:"minlon,attr"`
	} `xml:"bounds"`
	Desc     *string `xml:"desc"`
	Email    *string `xml:"email"`
	Keywords *string `xml:"keywords"`
	Metadata struct {
		Author struct {
			Email struct {
				Domain string `xml:"domain,attr"`
				ID     string `xml:"id,attr"`
			} `xml:"email"`
			Link struct {
				Href string `xml:"href,attr"`
				Text string `xml:"text"`
			} `xml:"link"`
			Name string `xml:"name"`
		} `xml:"author"`
		Copyright struct {
			Author  string `xml:"author,attr"`
			License string `xml:"license"`
			Year    int    `xml:"year"`
		} `xml:"copyright"`
		Desc     string `xml:"desc"`
		Keywords string `xml:"keywords"`
		Link     struct {
			Href string `xml:"href,attr"`
			Text string `xml:"text"`
			Type string `xml:"type"`
		} `xml:"link"`
		Name string    `xml:"name"`
		Time time.Time `xml:"time"`
	} `xml:"metadata"`
	Name *string `xml:"name"`
	Rte  []struct {
		Desc string `xml:"desc"`
		Link *struct {
			Href string `xml:"href,attr"`
			Text string `xml:"text"`
		} `xml:"link"`
		Name   string `xml:"name"`
		Number string `xml:"number"`
		RtePt  []struct {
			Lat  float64  `xml:"lat,attr"`
			Lon  float64  `xml:"lon,attr"`
			Cmt  *string  `xml:"cmt"`
			Desc string   `xml:"desc"`
			Ele  *float64 `xml:"ele"`
			Link *struct {
				Href string `xml:"href,attr"`
				Text string `xml:"text"`
			} `xml:"link"`
			Name string     `xml:"name"`
			Sym  string     `xml:"sym"`
			Time *time.Time `xml:"time"`
			Type string     `xml:"type"`
		} `xml:"rtept"`
	} `xml:"rte"`
	Time *time.Time `xml:"time"`
	Trk  []struct {
		Color *string `xml:"color"`
		Desc  string  `xml:"desc"`
		Link  *struct {
			Href string `xml:"href,attr"`
			Text string `xml:"text"`
		} `xml:"link"`
		Name   string `xml:"name"`
		Number string `xml:"number"`
		TrkSeg struct {
			TrkPt []struct {
				Lat  float64    `xml:"lat,attr"`
				Lon  float64    `xml:"lon,attr"`
				Ele  *float64   `xml:"ele"`
				Sym  string     `xml:"sym"`
				Time *time.Time `xml:"time"`
			} `xml:"trkpt"`
		} `xml:"trkseg"`
	} `xml:"trk"`
	URL     *string `xml:"url"`
	URLName *string `xml:"urlname"`
	Wpt     []struct {
		Lat  float64  `xml:"lat,attr"`
		Lon  float64  `xml:"lon,attr"`
		Cmt  *string  `xml:"cmt"`
		Desc string   `xml:"desc"`
		Ele  *float64 `xml:"ele"`
		Link *struct {
			Href string `xml:"href,attr"`
			Text string `xml:"text"`
		} `xml:"link"`
		Name *string    `xml:"name"`
		Sym  string     `xml:"sym"`
		Time *time.Time `xml:"time"`
		Type *string    `xml:"type"`
	} `xml:"wpt"`
}
