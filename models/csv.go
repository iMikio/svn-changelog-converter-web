package models

import (
	"encoding/csv"
	"io"
)

func WriteCsv(w io.Writer, s *[][]string, comma rune) error {
	cw := csv.NewWriter(w)
	cw.Comma = rune(comma)
	defer cw.Flush()
	return cw.WriteAll(*s)
}
