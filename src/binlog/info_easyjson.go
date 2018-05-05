// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package binlog

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonDdc53814DecodeBinlog(in *jlexer.Lexer, out *info) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "binlog":
			out.Binlog = string(in.String())
		case "gtid":
			out.Timestamp = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonDdc53814EncodeBinlog(out *jwriter.Writer, in info) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"binlog\":")
	out.String(string(in.Binlog))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"gtid\":")
	out.Int64(int64(in.Timestamp))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v info) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDdc53814EncodeBinlog(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *info) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDdc53814DecodeBinlog(&r, v)
	return r.Error()
}