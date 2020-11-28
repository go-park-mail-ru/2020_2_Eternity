// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

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

func easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain(in *jlexer.Lexer, out *Notification) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "to_id":
			out.ToUserId = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		case "msg":
			if in.IsNull() {
				in.Skip()
				out.EncodedData = nil
			} else {
				out.EncodedData = in.Bytes()
			}
		case "time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreationTime).UnmarshalJSON(data))
			}
		case "is_read":
			out.IsRead = bool(in.Bool())
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
func easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain(out *jwriter.Writer, in Notification) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"to_id\":"
		out.RawString(prefix)
		out.Int(int(in.ToUserId))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.Int(int(in.Type))
	}
	{
		const prefix string = ",\"msg\":"
		out.RawString(prefix)
		out.Base64Bytes(in.EncodedData)
	}
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix)
		out.Raw((in.CreationTime).MarshalJSON())
	}
	{
		const prefix string = ",\"is_read\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsRead))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Notification) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Notification) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Notification) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Notification) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain(l, v)
}
func easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain1(in *jlexer.Lexer, out *NoteResp) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.EncodedData = nil
			} else {
				out.EncodedData = in.Bytes()
			}
		case "creation_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreationTime).UnmarshalJSON(data))
			}
		case "is_read":
			out.IsRead = bool(in.Bool())
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
func easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain1(out *jwriter.Writer, in NoteResp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.Int(int(in.Type))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		out.Base64Bytes(in.EncodedData)
	}
	{
		const prefix string = ",\"creation_time\":"
		out.RawString(prefix)
		out.Raw((in.CreationTime).MarshalJSON())
	}
	{
		const prefix string = ",\"is_read\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsRead))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NoteResp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NoteResp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NoteResp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NoteResp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain1(l, v)
}
func easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain2(in *jlexer.Lexer, out *NotePin) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Id":
			out.Id = int(in.Int())
		case "Title":
			out.Title = string(in.String())
		case "ImgLink":
			out.ImgLink = string(in.String())
		case "UserId":
			out.UserId = int(in.Int())
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
func easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain2(out *jwriter.Writer, in NotePin) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"Title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"ImgLink\":"
		out.RawString(prefix)
		out.String(string(in.ImgLink))
	}
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix)
		out.Int(int(in.UserId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NotePin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NotePin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NotePin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NotePin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain2(l, v)
}
func easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain3(in *jlexer.Lexer, out *NoteFollow) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "FollowerId":
			out.FollowerId = int(in.Int())
		case "UserId":
			out.UserId = int(in.Int())
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
func easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain3(out *jwriter.Writer, in NoteFollow) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"FollowerId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.FollowerId))
	}
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix)
		out.Int(int(in.UserId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NoteFollow) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NoteFollow) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NoteFollow) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NoteFollow) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain3(l, v)
}
func easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain4(in *jlexer.Lexer, out *NoteComment) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Id":
			out.Id = int(in.Int())
		case "Path":
			if in.IsNull() {
				in.Skip()
				out.Path = nil
			} else {
				in.Delim('[')
				if out.Path == nil {
					if !in.IsDelim(']') {
						out.Path = make([]int32, 0, 16)
					} else {
						out.Path = []int32{}
					}
				} else {
					out.Path = (out.Path)[:0]
				}
				for !in.IsDelim(']') {
					var v7 int32
					v7 = int32(in.Int32())
					out.Path = append(out.Path, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Content":
			out.Content = string(in.String())
		case "PinId":
			out.PinId = int(in.Int())
		case "UserId":
			out.UserId = int(in.Int())
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
func easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain4(out *jwriter.Writer, in NoteComment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"Path\":"
		out.RawString(prefix)
		if in.Path == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Path {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"PinId\":"
		out.RawString(prefix)
		out.Int(int(in.PinId))
	}
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix)
		out.Int(int(in.UserId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NoteComment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NoteComment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9806e1EncodeGithubComGoParkMailRu20202EternityPkgDomain4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NoteComment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NoteComment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9806e1DecodeGithubComGoParkMailRu20202EternityPkgDomain4(l, v)
}