package drain

import (
	"github.com/bmizerany/lpx"
)

// ToString converts a parsed logplex log record back into its original string
func ToString(header *lpx.Header, data []byte) (string, error) {
	line := fmt.Sprintf("%s %s %s %s %s %s %s\n",
		string(header.PrivalVersion),
		string(header.Time),
		string(header.Hostname),
		string(header.Name),
		string(header.Procid),
		string(header.Msgid),
		string(data))
	return line
}
