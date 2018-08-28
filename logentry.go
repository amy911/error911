package error911

import (
	"net/html"
	"strings"
	"time"
)

type LogEntry struct {
	TimeStamp time.Time
	Language  string
	SubTitle  string
	Payload   string
}

func (le LogEntry) HTML() string {
	return "<h6><code class=\"timestamp\">" + le.TimeStamp.Format("2006-01-02_15:04:05.000000") +
		"</code><span class=\"subtitle\">" + le.SubTitle + "</span></h6>\n" +
		"<pre><code class=\"payload\" language=\"" + le.Language + "\">" +
		html.EscapeString(le.Payload) +
		"</pre></code>"
}

func (le LogEntry) MarkDown() string {
	return "###### `" + le.TimeStamp.Format("2006-01-02_15:04:05.000000") +
		"` \"" + le.SubTitle + "\" ######\n" +
		"```" + le.Language + "\n" +
		strings.Replace(le.Payload, "\n", "\n    ", -1) +
		"\n```"
}

func (le LogEntry) Text() string {
	return "--- " + le.TimeStamp.Format("2006-01-02_15:04:05.000000") +
		" \"" + le.SubTitle + "\"\n" +
		strings.Replace(le.Payload, "\n", "\n    ", -1)
}
