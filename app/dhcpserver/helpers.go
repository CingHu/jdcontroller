package dhcpserver

import (
	"encoding/binary"
	"net"
	"time"
	"strings"
	"strconv"
	"os/exec"
	"io/ioutil"
	"errors"
	"fmt"
)

// SelectOrderOrAll has same functionality as SelectOrder, except if the order
// param is nil, whereby all options are added (in arbitary order).
func (o Options) SelectOrderOrAll(order []byte) []Option {
	if order == nil {
		opts := make([]Option, 0, len(o))
		for i, v := range o {
			opts = append(opts, Option{Code: i, Value: v})
		}
		return opts
	}
	return o.SelectOrder(order)
}

// SelectOrder returns a slice of options ordered and selected by a byte array
// usually defined by OptionParameterRequestList.  This result is expected to be
// used in ReplyPacket()'s []Option parameter.
func (o Options) SelectOrder(order []byte) []Option {
	opts := make([]Option, 0, len(order))
	for _, v := range order {
		if data, ok := o[OptionCode(v)]; ok {
			opts = append(opts, Option{Code: OptionCode(v), Value: data})
		}
	}
	return opts
}

// IPRange returns how many ips in the ip range from start to stop (inclusive)
func IPRange(start, stop net.IP) int {
	//return int(Uint([]byte(stop))-Uint([]byte(start))) + 1
	return int(binary.BigEndian.Uint32(stop.To4())) - int(binary.BigEndian.Uint32(start.To4())) + 1
}

// IPAdd returns a copy of start + add.
// IPAdd(net.IP{192,168,1,1},30) returns net.IP{192.168.1.31}
func IPAdd(start net.IP, add int) net.IP { // IPv4 only
	start = start.To4()
	//v := Uvarint([]byte(start))
	result := make(net.IP, 4)
	binary.BigEndian.PutUint32(result, binary.BigEndian.Uint32(start)+uint32(add))
	//PutUint([]byte(result), v+uint64(add))
	return result
}

// IPLess returns where IP a is less than IP b.
func IPLess(a, b net.IP) bool {
	b = b.To4()
	for i, ai := range a.To4() {
		if ai != b[i] {
			return ai < b[i]
		}
	}
	return false
}

// IPInRange returns true if ip is between (inclusive) start and stop.
func IPInRange(start, stop, ip net.IP) bool {
	return !(IPLess(ip, start) || IPLess(stop, ip))
}

// OptionsLeaseTime - converts a time.Duration to a 4 byte slice, compatible
// with OptionIPAddressLeaseTime.
func OptionsLeaseTime(d time.Duration) []byte {
	leaseBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(leaseBytes, uint32(d/time.Second))
	//PutUvarint(leaseBytes, uint64(d/time.Second))
	return leaseBytes
}

// JoinIPs returns a byte slice of IP addresses, one immediately after the other
// This may be useful for creating multiple IP options such as OptionRouter.
func JoinIPs(ips []net.IP) (b []byte) {
	for _, v := range ips {
		b = append(b, v.To4()...)
	}
	return
}

func execShell(command string) ([]byte, error) {
	cmd := exec.Command("/bin/sh", "-c", command);

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return nil, err
	}

	if len(bytesErr) != 0 {
		return nil, errors.New(string(bytesErr))
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	if len(bytes) == 0 {
		return nil, errors.New("We get nothing")
	}
	return bytes, nil
}

func stringToIp(bytes string) []byte {
	fmt.Println("string", bytes)
	bits := strings.Split(bytes, ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	ipaddr := make([]byte, 4)
	ipaddr[0] = byte(b0)
	ipaddr[1] = byte(b1)
	ipaddr[2] = byte(b2)
	ipaddr[3] = byte(b3)

	return ipaddr
}

func GetIpByName(name string) ([]byte, error) {
	//data, err := execShell("ifconfig " + name + " | grep -w \"inet\"")
	//if err != nil {
	//	return nil, err
	//}

	//ipstring := strings.Split(data, " ")
	//deviceIp := stringToIp(ipstring[1])
	deviceIp := stringToIp("192.168.0.23")

	return deviceIp, nil
}