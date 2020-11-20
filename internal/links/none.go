package links

import "github.com/vishvananda/netlink"

// None is an empty link, used for getting all addresses on the system.
type None struct {
	attrs netlink.LinkAttrs
}

// Attrs returns the link attributes.
func (n *None) Attrs() *netlink.LinkAttrs {
	return &n.attrs
}

// Type returns the link type.
func (n *None) Type() string {
	return ""
}

// NewNone creates a new none link.
func NewNone() *None {
	none := new(None)
	none.attrs = netlink.NewLinkAttrs()
	return none
}
