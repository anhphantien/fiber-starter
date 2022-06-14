package config

type __ContentType struct {
	JPEG string
	PNG  string
}

var _ContentType = __ContentType{
	JPEG: "image/jpeg",
	PNG:  "image/png",
}

type _File struct {
	MaxSize     int64
	ContentType __ContentType
}

var File = _File{
	MaxSize:     10 * 1000 * 1000, // 10 MB
	ContentType: _ContentType,
}
