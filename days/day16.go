package days

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
	"strconv"
)

func Day16(inputFile string, part int) {
	rs := util.RuneScanner(inputFile)
	s := &Stream{rs, "", make([]uint64, 0), 0, 0}
	p := decodeBITS(s)
	fmt.Println("Accumulate version sum: ", p.cumsum)
}

func decodeBITS(s *Stream) Packet {
	s.curr = 0
	version := s.feed(3, false)
	typeID := s.feed(3, false)
	p := Packet{version, typeID, -1, -1, -1, -1, version, make([]Packet, 0)}
	if p.typeID == 4 {
		for s.feed(1, false) == 1 {
			s.feed(4, true)
		}
		s.feed(4, true)
		p.bits = s.curr
		p.literal = s.accept()
	} else {
		p.ltypeID = s.feed(1, false) // 1 for 15-bit and stream length, 2 for 11-bit packet-number
		if p.ltypeID == 0 {
			p.max = s.feed(15, false)
			p.bits = s.curr
			for p.bits-22 < p.max {
				subpacket := decodeBITS(s)
				p.bits += subpacket.bits
				p.cumsum += subpacket.cumsum
				p.subpackets = append(p.subpackets, subpacket)
			}
		} else if p.ltypeID == 1 {
			p.max = s.feed(11, false)
			p.bits = s.curr
			for i := 0; i < p.max; i++ {
				subpacket := decodeBITS(s)
				p.cumsum += subpacket.cumsum
				p.bits += subpacket.bits
				p.subpackets = append(p.subpackets, subpacket)
			}
		}
	}

	return p
}

// ======
// PACKET
// ======
type Packet struct {
	version    int
	typeID     int
	// for literal packet values
	literal    int
	// for operator packet values
	ltypeID    int // 0: length is a 15-bit number of bits, 1: length is a 11-bit number of packets
	max        int
	bits 		int // bits used to represent this packet (including subpackets)
	cumsum 		int
	subpackets []Packet
}

func (p *Packet) print(tab string, vsum *int) {
	fmt.Printf("%s* Version: %d | Type: %d ", tab, p.version, p.typeID)
	if p.typeID == 4 {
		fmt.Printf("| Literal: %d\n", p.literal)
		return
	} else if p.ltypeID == 0 {
		fmt.Printf("| Operator | Bit-length %d\n", p.max)
	} else if p.ltypeID == 1 {
		fmt.Printf("| Operator | Packet-length %d\n", p.max)
	}

	fmt.Printf("%s  Subpackets: \n", tab)
	tab += "\t"
}

// ============
// INPUT STREAM
// ============
// Input stream allows us to read in hex as binary data, to either discard immediately
// and return the value, or keep in the buffer which can be evaluated later on.
type Stream struct {
	scanner 	*bufio.Scanner
	buffer 		string // binary string of 0s and 1s
	backlog 	[]uint64 // the next values to feed in
	max 		int // maximum number of bits to read before hard stop
	curr 		int // number of bits read so far
}

// Feeds the next n bytes in to the stream
// keep = false:
//		X where X is the int value of the stream
//		Stream is emptied.
//
// keep = true:
//		returns -1, because stream is not ready to be read.
//		Stream is not emptied
func (s *Stream) feed(n int, keep bool) int {
	// Count the number of bits we feed in
	s.curr += n
	for len(s.backlog) < n {
		r, _ := util.Read(s.scanner)
		s.backlog = append(s.backlog, util.Hex2Bits(r)...)
	}

	tempBuffer := ""
	toFeed := s.backlog[:n]
	for _, f := range toFeed {
		tempBuffer += strconv.FormatUint(f, 2)
	}
	s.backlog = s.backlog[n:]

	if !keep {
		literal := util.Bin2Dec(tempBuffer)
		return literal
	} else {
		s.buffer += tempBuffer
		return -1
	}
}

func (s *Stream) accept() int {
	val := util.Bin2Dec(s.buffer)
	s.buffer = ""
	return val
}
