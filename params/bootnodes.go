// Copyright 2015 The go-popcateum Authors
// This file is part of the go-popcateum library.
//
// The go-popcateum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-popcateum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-popcateum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/popcateum/go-popcateum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Popcateum network.
var MainnetBootnodes = []string{
	// Popcateum Foundation Go Bootnodes
	"enode://903d9c6369a23e991027a5ca9ca7dc25c0f9ed2980dca43b0849e13d7cdb3a1d947df91afb9cc4880141e748800722cecf6ad98820cd53a4664254bec173241b@34.64.121.103:60606",
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	//
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	//
}

// HappycatBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var HappycatBootnodes = []string{
	//
}

var V5Bootnodes = []string{
	// Popcateum Foundation Go Bootnodes
	"enr:-KC4QGy0-EGaYy100TynH7pIeq3811WXlAsy5l3IGlUs3VlYYmvMIvBgowNbHWtJbja9DPmx1eM2qUtQ6KrlM36d0hYCg2V0aMrJhKNYcKeDPQkAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQOQPZxjaaI-mRAnpcqcp9wlwPntKYDcpDsISeE9fNs6HYRzbmFwwIN0Y3CC7L6DdWRwguy-",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/popcateum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case HappycatGenesisHash:
		net = "happycat"
	case SepoliaGenesisHash:
		net = "sepolia"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
