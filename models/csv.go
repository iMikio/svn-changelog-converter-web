package models

import (
	"encoding/csv"
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func WriteCsv(w io.Writer, s *[][]string, comma rune, code *string) error {
	if *code == "shift-jis" {
		tw := transform.NewWriter(w, japanese.ShiftJIS.NewEncoder())
		defer tw.Close()
		return writeCsv(tw, s, comma, true)
	} else {
		return writeCsv(w, s, comma, false)
	}
}

func writeCsv(w io.Writer, s *[][]string, comma rune, useCRLF bool) error {
	cw := csv.NewWriter(w)
	cw.Comma = rune(comma)
	cw.UseCRLF = useCRLF
	defer cw.Flush()
	return cw.WriteAll(*s)
}
