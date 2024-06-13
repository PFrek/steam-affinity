package urls

import (
	"fmt"
	"strings"
)

type URLChain struct {
	URL string
}

func (chain URLChain) lastChar() string {
	if len(chain.URL) == 0 {
		return ""
	}

	return fmt.Sprintf("%c", chain.URL[len(chain.URL)-1])
}

func (chain URLChain) chain(path string) URLChain {
	base := chain.URL

	if chain.lastChar() != "/" {
		base += "/"
	}

	path, _ = strings.CutPrefix(path, "/")

	return URLChain{
		URL: base + path,
	}
}
