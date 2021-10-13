package models

import (
	"encoding/csv"
	"io"
)

func WriteCsv(w io.Writer, s *[][]string) error {
	cw := csv.NewWriter(w)
	cw.Comma = rune('\t')
	defer cw.Flush()
	return cw.WriteAll(*s)
}
