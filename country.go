package caddy_maxmind_geolocation

type Names struct {
	De   string `maxminddb:"de"`
	En   string `maxminddb:"en"`
	Es   string `maxminddb:"es"`
	Fr   string `maxminddb:"fr"`
	PtBr string `maxminddb:"pt-BR"`
	Ru   string `maxminddb:"ru"`
	ZhCn string `maxminddb:"zh-CN"`
}

type Country struct {
	GeonameId         int    `maxminddb:"geoname_id"`
	IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	ISOCode           string `maxminddb:"iso_code"`
	Names             Names  `maxminddb:"names"`
}

type Subdivision struct {
	ISOCode string `maxminddb:"iso_code"`
}

type Record struct {
	Country      Country     `maxminddb:"country"`
	Subdivision1 Subdivision `maxminddb:"subdivision_1"`
	Subdivision2 Subdivision `maxminddb:"subdivision_2"`
	MetroCode    string      `json:"metro_code"`
}
