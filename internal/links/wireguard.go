package links

import "github.com/vishvananda/netlink"

// Wireguard links.
type Wireguard struct {
	attrs netlink.LinkAttrs
}

// Attrs returns the link attributes.
func (w *Wireguard) Attrs() *netlink.LinkAttrs {
	return &w.attrs
}

// Type returns the link type.
func (w *Wireguard) Type() string {
	return "wireguard"
}

// NewWireguard creates a new wireguard link.
func NewWireguard(name string) *Wireguard {
	wg := new(Wireguard)
	wg.attrs = netlink.NewLinkAttrs()
	wg.attrs.Name = name
	return wg
}
