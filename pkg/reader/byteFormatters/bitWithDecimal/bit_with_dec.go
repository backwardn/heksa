package bitWithDecimal

import (
	"github.com/raspi/heksa/pkg/reader/byteFormatters/base"
	"github.com/raspi/heksa/pkg/reader/byteFormatters/bit"
	"github.com/raspi/heksa/pkg/reader/byteFormatters/decimal"
)

type printer struct {
	p base.ByteFormatter
}

func New() base.ByteFormatter {
	return printer{
		p: bit.New(),
	}
}

func (p printer) Print(b byte) (o string) {
	base.ChangePalette = true

	o += p.p.Print(b)
	o += ` ` + base.SpecialBreak + `[` + base.HilightBreak
	o += decimal.DecByteToString[b]
	o += base.SpecialBreak + `]`
	return o
}

func (p printer) GetPrintSize() int {
	return 14
}
