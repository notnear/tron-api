package client

var (
	MAIN   = "https://api.trongrid.io"
	Shasta = "https://api.shasta.trongrid.io"
	NILE   = "https://nile.trongrid.io"
)

type Client struct {
	Host    string
	Headers map[string]string
	Debug   bool
}
