package parsers

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"

	"golang.org/x/text/encoding/charmap"

	"github.com/valyala/fasthttp"
)

type ValCurs struct {
	Valute []Valute
}

type Valute struct {
	// NumCode  uint16
	CharCode string
	Nominal  uint16
	// Name     string
	Value float32
}

var valCurs ValCurs

func LoadCbrData() (ValCurs, error) {
	_, raw, err := fasthttp.Get(nil, "http://www.cbr.ru/scripts/XML_daily.asp")

	if nil != err {
		return valCurs, err
	}

	return ParseCbrData(raw)
}

func ParseCbrData(data []byte) (ValCurs, error) {
	decoder := xml.NewDecoder(bytes.NewReader([]byte(strings.Replace(string(data), ",", ".", -1))))
	decoder.CharsetReader = charsetReader

	return valCurs, decoder.Decode(&valCurs)
}

func charsetReader(_ string, input io.Reader) (io.Reader, error) {
	return charmap.Windows1251.NewDecoder().Reader(input), nil
}
