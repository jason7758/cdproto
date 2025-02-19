// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package fedcm

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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm(in *jlexer.Lexer, out *SelectAccountParams) {
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
		case "dialogId":
			out.DialogID = string(in.String())
		case "accountIndex":
			out.AccountIndex = int64(in.Int64())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm(out *jwriter.Writer, in SelectAccountParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dialogId\":"
		out.RawString(prefix[1:])
		out.String(string(in.DialogID))
	}
	{
		const prefix string = ",\"accountIndex\":"
		out.RawString(prefix)
		out.Int64(int64(in.AccountIndex))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SelectAccountParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SelectAccountParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SelectAccountParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SelectAccountParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm1(in *jlexer.Lexer, out *EventDialogShown) {
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
		case "dialogId":
			out.DialogID = string(in.String())
		case "accounts":
			if in.IsNull() {
				in.Skip()
				out.Accounts = nil
			} else {
				in.Delim('[')
				if out.Accounts == nil {
					if !in.IsDelim(']') {
						out.Accounts = make([]*Account, 0, 8)
					} else {
						out.Accounts = []*Account{}
					}
				} else {
					out.Accounts = (out.Accounts)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Account
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Account)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Accounts = append(out.Accounts, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm1(out *jwriter.Writer, in EventDialogShown) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dialogId\":"
		out.RawString(prefix[1:])
		out.String(string(in.DialogID))
	}
	{
		const prefix string = ",\"accounts\":"
		out.RawString(prefix)
		if in.Accounts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Accounts {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventDialogShown) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventDialogShown) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventDialogShown) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventDialogShown) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm2(in *jlexer.Lexer, out *EnableParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm2(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm3(in *jlexer.Lexer, out *DismissDialogParams) {
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
		case "dialogId":
			out.DialogID = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm3(out *jwriter.Writer, in DismissDialogParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dialogId\":"
		out.RawString(prefix[1:])
		out.String(string(in.DialogID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DismissDialogParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DismissDialogParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DismissDialogParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DismissDialogParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm4(in *jlexer.Lexer, out *DisableParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm4(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm5(in *jlexer.Lexer, out *Account) {
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
		case "accountId":
			out.AccountID = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "givenName":
			out.GivenName = string(in.String())
		case "pictureUrl":
			out.PictureURL = string(in.String())
		case "idpConfigUrl":
			out.IdpConfigURL = string(in.String())
		case "idpSigninUrl":
			out.IdpSigninURL = string(in.String())
		case "loginState":
			(out.LoginState).UnmarshalEasyJSON(in)
		case "termsOfServiceUrl":
			out.TermsOfServiceURL = string(in.String())
		case "privacyPolicyUrl":
			out.PrivacyPolicyURL = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm5(out *jwriter.Writer, in Account) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"accountId\":"
		out.RawString(prefix[1:])
		out.String(string(in.AccountID))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"givenName\":"
		out.RawString(prefix)
		out.String(string(in.GivenName))
	}
	{
		const prefix string = ",\"pictureUrl\":"
		out.RawString(prefix)
		out.String(string(in.PictureURL))
	}
	{
		const prefix string = ",\"idpConfigUrl\":"
		out.RawString(prefix)
		out.String(string(in.IdpConfigURL))
	}
	{
		const prefix string = ",\"idpSigninUrl\":"
		out.RawString(prefix)
		out.String(string(in.IdpSigninURL))
	}
	{
		const prefix string = ",\"loginState\":"
		out.RawString(prefix)
		(in.LoginState).MarshalEasyJSON(out)
	}
	if in.TermsOfServiceURL != "" {
		const prefix string = ",\"termsOfServiceUrl\":"
		out.RawString(prefix)
		out.String(string(in.TermsOfServiceURL))
	}
	if in.PrivacyPolicyURL != "" {
		const prefix string = ",\"privacyPolicyUrl\":"
		out.RawString(prefix)
		out.String(string(in.PrivacyPolicyURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Account) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Account) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoFedcm5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Account) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Account) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoFedcm5(l, v)
}
