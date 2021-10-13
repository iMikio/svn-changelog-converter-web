package models

import (
	"encoding/csv"
	"io"
)

func WriteCsv(w io.Writer, s *[][]string, comma rune, useCRLF bool) error {
	cw := csv.NewWriter(w)
	cw.Comma = rune(comma)
	cw.UseCRLF = useCRLF
	defer cw.Flush()
	return cw.WriteAll(*s)
}
