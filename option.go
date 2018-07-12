package gojsonq

import "errors"

// option describes type for providing configuration options to JSONQ
type option struct {
	decoder Decoder
}

// OptionFunc represents a contract for option func, it basically set options to jsonq instance options
type OptionFunc func(*JSONQ) error

// SetDecoder take a custom decoder to decode JSON
func SetDecoder(u Decoder) OptionFunc {
	return func(j *JSONQ) error {
		if u == nil {
			return errors.New("decoder can not be nil")
		}
		j.option.decoder = u
		return nil
	}
}
