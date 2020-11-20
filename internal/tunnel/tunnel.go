package tunnel

import (
	"fmt"
	"net"

	"github.com/selftechio/sandhog/internal/links"
	"github.com/vishvananda/netlink"
)

// Tunnel contains the data sandhog needs to know about tunnels.
type Tunnel struct {
	Address net.IPNet
	Name    string
}

// NewTunnel creates a new tunnel.
func NewTunnel(address, name string) (*Tunnel, error) {
	// parse the link address
	linkAddr, err := netlink.ParseAddr(address)
	if err != nil {
		return nil, err
	}

	// find existing addresses and check whether the proposed address already exists
	// if the address already exists, return with an error
	existingLinks, err := netlink.LinkList()
	if err != nil {
		return nil, err
	}
	// fixme 20/11/2020: a double for loop is probabil not the best idea
	for _, existingLink := range existingLinks {
		existingAddrList, err := netlink.AddrList(existingLink, netlink.FAMILY_ALL)
		if err != nil {
			return nil, err
		}
		for _, existingAddr := range existingAddrList {
			if existingAddr.Equal(*linkAddr) {
				return nil, fmt.Errorf("proposed address is already in use: %s", existingAddr.String())
			}
		}
	}

	// create the wireguard link
	link := links.NewWireguard(name)
	err = netlink.LinkAdd(link)
	if err != nil {
		return nil, err
	}

	// add the address to the wireguard link
	// remove the link if unsuccessful
	err = netlink.AddrAdd(link, linkAddr)
	if err != nil {
		err2 := netlink.LinkDel(link)
		if err2 != nil {
			return nil, fmt.Errorf("failed to delete link %s after failing to add address %s to it", name, linkAddr.String())
		}
		return nil, err
	}

	// bring the link up
	// remove the link if unsuccessful
	err = netlink.LinkSetUp(link)
	if err != nil {
		err2 := netlink.LinkDel(link)
		if err2 != nil {
			return nil, fmt.Errorf("failed to delete link %s after failing to bring it up", name)
		}
	}

	tunnel := new(Tunnel)
	tunnel.Address = *linkAddr.IPNet
	tunnel.Name = name

	return tunnel, nil
}

// Close closes a tunnel.
func (t *Tunnel) Close() error {
	link := links.NewWireguard(t.Name)
	return netlink.LinkDel(link)
}
