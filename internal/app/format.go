package app

import "errors"

type Format string

type adrListEntry struct {
	Number   int    `json:"number"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Date     string `json:"date"`
	Filepath string `json:"filepath"`
}

const (
	FormatMd   Format = "md"
	FormatRaw  Format = "raw"
	FormatJSON Format = "json"
)

var ErrInvalidFormat = errors.New("invalid format")

func ParseFormat(s string) (Format, error) {
	switch Format(s) {
	case FormatMd, FormatRaw, FormatJSON:
		return Format(s), nil
	default:
		return FormatMd, ErrInvalidFormat
	}
}
