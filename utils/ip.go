/*
ip 地址处理工具包

*/
package utils

import (
	"errors"
	"net"
)

func Uint32ToIP(ip uint32) net.IP {
	res := make([]byte, 0, 4)
	res = append(res, uint8(ip>>24))
	res = append(res, uint8(ip>>16))
	res = append(res, uint8(ip>>8))
	res = append(res, uint8(ip))
	return net.IP(res)
}

func IP2Uint32(ip net.IP) uint32 {
	ip = ip.To4()
	return uint32(ip[0])<<24 + uint32(ip[1])<<16 + uint32(ip[2])<<8 + uint32(ip[3])
}

func Uint32TO4Byte(ip uint32) [4]byte {
	var b [4]byte
	b[0] = uint8(ip >> 24)
	b[1] = uint8(ip >> 16)
	b[2] = uint8(ip >> 8)
	b[3] = uint8(ip)
	return b
}

func String2IPV4(ip string) (net.IP, error) {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return nil, errors.New("ip is not valid")
	}
	return ipv4, nil
}
