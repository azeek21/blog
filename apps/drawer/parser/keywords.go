package drawer_parser

type keywords struct {
	END          []byte
	COLORS_START []byte
	IMAGE_START  []byte
	VERSION      []byte
	COMMENT      []byte
}

var KEYWORDS = keywords{
	END:          []byte("end"),
	COLORS_START: []byte("colors"),
	IMAGE_START:  []byte("image"),
	VERSION:      []byte("version"),
	COMMENT:      []byte("comment"),
}
