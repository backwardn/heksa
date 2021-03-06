package bit

import (
	"github.com/raspi/heksa/pkg/color"
	"github.com/raspi/heksa/pkg/reader/byteFormatters/base"
)

type printer struct {
}

func New() base.ByteFormatter {
	return printer{}
}

func (p printer) Print(b byte) (o string) {
	if base.ChangePalette {
		base.ChangePalette = false
		o += base.Palette[b]
	}

	for idx, ru := range bitByteToString[b] {
		if idx == 0 {
			o += color.SetUnderlineOn
		}

		o += string(ru)

		if idx == 3 {
			o += color.SetUnderlineOff
		}
	}

	return o
}

func (p printer) GetPrintSize() int {
	return 8
}
