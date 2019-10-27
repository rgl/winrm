package winrm

import (
	"encoding/xml"

	"strings"
)

func xmlEscapeText(text string) string {
	var b strings.Builder
	xml.EscapeText(&b, []byte(text))
	return b.String()
}
