package sessions

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestNewID(t *testing.T) {
	cases := []struct {
		name        string
		hint        string
		signingKey  string
		expectError bool
	}{
		{
			"Empty Signing Key",
			"Remember to return an error if `signingKey` is zero-length",
			"",
			true,
		},
		{
			"Valid Signing Key",
			"Remember to return a valid base64-url-encoded SessionID if the `signingKey` is non-zero-length",
			"test key",
			false,
		},
	}
	for _, c := range cases {
		sid, err := NewSessionID(c.signingKey)
		fmt.Println(sid)
		if err != nil && !c.expectError {
			t.Errorf("case %s: unexpected error generating new SessionID: %v\nHINT: %s", c.name, err, c.hint)
		}
		if err == nil {
			//ensure sid is non-zero-length
			if len(sid) == 0 {
				t.Errorf("case %s: new SessionID is zero-length\nHINT: %s", c.name, c.hint)
			}
			//ensure sid is base64-url-encoded
			_, err := base64.URLEncoding.DecodeString(string(sid))
			if err != nil {
				t.Errorf("case %s: new SessionID failed base64-url-decoding: %v\nHINT: %s", c.name, err, c.hint)
			}
		}
	}
}
