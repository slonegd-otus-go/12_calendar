// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: ./web/models/event.go

package models

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Event) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *Event) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	if j.Date != nil {
		buf.WriteString(`{ "date":`)
		fflib.WriteJsonString(buf, string(*j.Date))
	} else {
		buf.WriteString(`{ "date":null`)
	}
	if j.Description != nil {
		buf.WriteString(`,"description":`)
		fflib.WriteJsonString(buf, string(*j.Description))
	} else {
		buf.WriteString(`,"description":null`)
	}
	if j.Duration != nil {
		buf.WriteString(`,"duration":`)
		fflib.FormatBits2(buf, uint64(*j.Duration), 10, *j.Duration < 0)
	} else {
		buf.WriteString(`,"duration":null`)
	}
	buf.WriteByte(',')
	if j.ID != 0 {
		buf.WriteString(`"id":`)
		fflib.FormatBits2(buf, uint64(j.ID), 10, j.ID < 0)
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtEventbase = iota
	ffjtEventnosuchkey

	ffjtEventDate

	ffjtEventDescription

	ffjtEventDuration

	ffjtEventID
)

var ffjKeyEventDate = []byte("date")

var ffjKeyEventDescription = []byte("description")

var ffjKeyEventDuration = []byte("duration")

var ffjKeyEventID = []byte("id")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Event) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Event) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtEventbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtEventnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffjKeyEventDate, kn) {
						currentKey = ffjtEventDate
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyEventDescription, kn) {
						currentKey = ffjtEventDescription
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyEventDuration, kn) {
						currentKey = ffjtEventDuration
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyEventID, kn) {
						currentKey = ffjtEventID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyEventID, kn) {
					currentKey = ffjtEventID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyEventDuration, kn) {
					currentKey = ffjtEventDuration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyEventDescription, kn) {
					currentKey = ffjtEventDescription
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyEventDate, kn) {
					currentKey = ffjtEventDate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtEventnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtEventDate:
					goto handle_Date

				case ffjtEventDescription:
					goto handle_Description

				case ffjtEventDuration:
					goto handle_Duration

				case ffjtEventID:
					goto handle_ID

				case ffjtEventnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Date:

	/* handler: j.Date type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			j.Date = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			j.Date = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Description:

	/* handler: j.Description type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			j.Description = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			j.Description = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Duration:

	/* handler: j.Duration type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

			j.Duration = nil

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			ttypval := int64(tval)
			j.Duration = &ttypval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ID:

	/* handler: j.ID type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.ID = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
