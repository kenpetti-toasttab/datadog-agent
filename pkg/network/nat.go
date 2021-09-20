package network

import (
	"inet.af/netaddr"
)

// GetNATLocalAddress returns the translated (local ip, local port) pair
func GetNATLocalAddress(c ConnectionStats) (netaddr.IP, uint16) {
	localIP := c.Source
	localPort := c.SPort

	if c.IPTranslation != nil && !c.IPTranslation.ReplDstIP.IsZero() {
		// Fields are flipped
		localIP = c.IPTranslation.ReplDstIP
		localPort = c.IPTranslation.ReplDstPort
	}
	return localIP, localPort
}

// GetNATRemoteAddress returns the translated (remote ip, remote port) pair
func GetNATRemoteAddress(c ConnectionStats) (netaddr.IP, uint16) {
	remoteIP := c.Dest
	remotePort := c.DPort

	if c.IPTranslation != nil && !c.IPTranslation.ReplSrcIP.IsZero() {
		// Fields are flipped
		remoteIP = c.IPTranslation.ReplSrcIP
		remotePort = c.IPTranslation.ReplSrcPort
	}
	return remoteIP, remotePort
}
