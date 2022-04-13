package gmail

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type MimeMessage struct {
	Sender     string
	To         []string
	Cc         []string
	Bcc        []string
	Subject    string
	Body       string
	IsHtmlBody bool
}

func (mm *MimeMessage) Raw(encode bool) string {

	msg := ""
	if mm.IsHtmlBody {
		msg = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	}

	msg += fmt.Sprintf("From: %s\r\n", mm.Sender)

	if len(mm.To) > 0 {
		msg += fmt.Sprintf("To: %s\r\n", mm.To[0])
	}

	if len(mm.Cc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mm.Cc, ";"))
	}

	msg += fmt.Sprintf("Subject: %s\r\n", mm.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mm.Body)

	if encode {
		data := []byte(msg)
		dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(dst, data)
		msg = string(dst)
	}

	return msg
}
