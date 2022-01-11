package http2

import (
	"strings"

	mdata "github.com/go-gost/gost/pkg/metadata"
)

type metadata struct {
	proxyAgent      string
	probeResistance *probeResistance
	sni             bool
	enableUDP       bool
}

func (h *http2Handler) parseMetadata(md mdata.Metadata) error {
	const (
		proxyAgent     = "proxyAgent"
		probeResistKey = "probeResistance"
		knock          = "knock"
		sni            = "sni"
		enableUDP      = "udp"
	)

	h.md.proxyAgent = mdata.GetString(md, proxyAgent)

	if v := mdata.GetString(md, probeResistKey); v != "" {
		if ss := strings.SplitN(v, ":", 2); len(ss) == 2 {
			h.md.probeResistance = &probeResistance{
				Type:  ss[0],
				Value: ss[1],
				Knock: mdata.GetString(md, knock),
			}
		}
	}
	h.md.sni = mdata.GetBool(md, sni)
	h.md.enableUDP = mdata.GetBool(md, enableUDP)

	return nil
}

type probeResistance struct {
	Type  string
	Value string
	Knock string
}
