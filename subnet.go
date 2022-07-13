package main

import (
	"errors"
	"net/netip"
	"strconv"
	"strings"
)

var (
	ErrInvalidMask = errors.New("invalid mask")
	ErrInvalidIP   = errors.New("invalid ip")
)

const big = 0xFFFFFF

// Parse returns the ip and available net address range.
//  For example, ParseCIDR("192.0.2.1/24") returns the IP address
// 192.0.2.1 and the available network address range 192.0.2.1 - 192.0.2.254
func Parse(sn string) (string, string, error) {
	i := strings.Index(sn, "/")
	addr, err := netip.ParseAddr(sn[:i])
	if err != nil {
		return "", "", err
	}
	n, err := strconv.Atoi(sn[i+1:])
	if err != nil {
		return "", "", err
	}

	if n > big {
		n = big
	}

	var ar string
	if addr.Is4() {
		in, err := v4ToI(addr.String())
		if err != nil {
			return "", "", err
		}
		ar = iToV4(in)
	} else if addr.Is6() {
		// TODO: IPv6

	}
	return addr.String(), ar, nil
}

func iToV4(i int64) string {
	b0 := strconv.FormatInt((i>>24)&0xff, 10)
	b1 := strconv.FormatInt((i>>16)&0xff, 10)
	b2 := strconv.FormatInt((i>>8)&0xff, 10)
	b3 := strconv.FormatInt(i&0xff, 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func v4ToI(s string) (int64, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 4 {
		return 0, ErrInvalidIP
	}
	var i int64
	for _, part := range parts {
		n, err := strconv.ParseInt(part, 10, 8)
		if err != nil {
			return 0, err
		}
		i = (i << 8) | n
	}

	return i, nil
}
