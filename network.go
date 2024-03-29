package main

import (
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type NetworkConfig struct {
	Date     *time.Time  `yaml:",omitempty"`
	Networks NetworkList `yaml:",omitempty"`
}

type NetworkList []*Network

func (n *NetworkList) GetByName(search string) (*Network, error) {
	for _, network := range *n {
		if network.Name == search {
			return network, nil
		}
	}
	return &Network{}, errors.New("network not found: " + search)
}

type Network struct {
	Name    string      `yaml:""`
	Subnet  string      `yaml:""`
	Domains *[]string   `yaml:",omitempty"`
	Sub     NetworkList `yaml:",omitempty"`
	Devices DeviceList  `yaml:",omitempty"`

	parent *Network
}

func (n *Network) SetParents() {
	for _, sub := range n.Sub {
		sub.parent = n
		sub.SetParents()
	}
}

func (n *Network) GetDomains() *[]string {
	if n.Domains != nil {
		return n.Domains
	}

	// look up the tree
	p := n.parent
	if p == nil {
		return nil
	}
	return p.GetDomains()
}

func (n *Network) GetIPv4Addresses() ([]string, error) {
	return getHosts(n.Subnet)
}

func (n *Network) GetDeviceByIPv4(ipv4 string) *Device {
	return n.Devices.GetByIPv4(n, ipv4)
}

func (n *Network) GetIPv4First() (string, error) {
	hosts, err := getHosts(n.Subnet)
	if err != nil {
		return "", err
	}
	if len(hosts) < 1 {
		return "", errors.New("no address in network")
	}
	return hosts[0], nil
}

func (n *Network) GetIPv4Last() (string, error) {
	hosts, err := getHosts(n.Subnet)
	if err != nil {
		return "", err
	}
	if len(hosts) < 1 {
		return "", errors.New("no address in network")
	}
	return hosts[len(hosts)-1], nil
}

type DeviceList []*Device

func (d *DeviceList) GetByIPv4(parentNetwork *Network, address string) *Device {
	for _, device := range *d {
		if ip, err := device.GetIPv4(parentNetwork); err == nil && ip == address {
			return device
		}
	}
	return nil
}

type Device struct {
	Name        string   `yaml:""`
	Description string   `yaml:",omitempty"`
	MAC         string   `yaml:",omitempty"`
	Hostnames   []string `yaml:",omitempty"`
	Network     string   `yaml:",omitempty"`
	IPv4        string   `yaml:",omitempty"`
}

func (d *Device) GetHostnames() []string {
	if len(d.Hostnames) > 0 {
		return d.Hostnames
	}

	return []string{d.Name}
}

func (d *Device) GetHostnamesInNetwork(network *Network) []string {
	result := make([]string, 0)

	domains := network.GetDomains()
	if domains == nil || len(*domains) == 0 {
		domains = &[]string{""}
	}
	hostnames := d.GetHostnames()
	for _, domain := range *domains {
		if domain != "" && !strings.HasPrefix(domain, ".") {
			domain = "." + domain
		}
		for _, c := range hostnames {
			result = append(result, c+domain)
		}
	}
	return result
}

func (d *Device) GetName() string {
	return d.Name
}

func (d *Device) GetNetwork(nl NetworkList) (*Network, error) {
	return nl.GetByName(d.Name)
}

func (d *Device) GetNetworkName() string {
	return d.Network
}

func (d *Device) GetDescription() string {
	return d.Description
}

func (d *Device) GetIPv4(parentNetwork *Network) (string, error) {
	if !strings.Contains(d.IPv4, "nw+") {
		return d.IPv4, nil
	}

	hosts, err := getHosts(parentNetwork.Subnet)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	re := regexp.MustCompile(`nw\+(\d+)`)
	result := re.FindStringSubmatch(d.IPv4)
	if len(result) < 2 {
		return "", errors.New("invalid ip definition: " + d.IPv4)
	}
	index, err := strconv.Atoi(result[1])
	if err != nil {
		return "", err
	}
	index-- // 1-based vs. 0-based
	if index >= len(hosts) {
		return "", errors.New("ip not in network range: " + d.IPv4)
	}
	return hosts[index], nil
}

func LoadNetworkConfigFromFile(filename string) (*NetworkConfig, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	return LoadNetworkConfig(f)
}

func LoadNetworkConfig(from io.Reader) (*NetworkConfig, error) {
	data, err := io.ReadAll(from)
	if err != nil {
		return nil, err
	}

	var networkConfig NetworkConfig
	err = yaml.Unmarshal(data, &networkConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Unmarshalling failed")
	}

	// augment for tree structure
	for _, n := range networkConfig.Networks {
		n.SetParents()
	}

	return &networkConfig, nil
}
