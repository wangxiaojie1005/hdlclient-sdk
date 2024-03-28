package units

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// 随机16位的UUID
func GetUuid() []byte {
	b := make([]byte, 16)
	io.ReadFull(rand.Reader, b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return b
}

// b2s converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func Byte2string(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

// s2b converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func String2byte(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

//整形转换成字节
func UintToBytes(n int, m int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if m == 8 {
		binary.Write(bytesBuffer, binary.BigEndian, uint8(n))
	} else if m == 16 {
		binary.Write(bytesBuffer, binary.BigEndian, uint16(n))
	} else if m == 32 {
		binary.Write(bytesBuffer, binary.BigEndian, uint32(n))
	} else if m == 64 {
		binary.Write(bytesBuffer, binary.BigEndian, uint64(n))
	} else {
		binary.Write(bytesBuffer, binary.BigEndian, uint32(n))
	}
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToUint(b []byte, m int) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

func BytesToUint8(b []byte) uint8 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint8
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return uint8(x)
}

func BytesToUint16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint16
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return uint16(x)
}

func BytesToUint32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return uint32(x)
}

func BytesToUint64(b []byte) uint64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return uint64(x)
}

// DecimalToBinary() function that will take Decimal number as int, and return it's Binary equivalent as string.
func DecimalToBinary(num int) (string, error) {
	if num < 0 {
		return "", errors.New("integer must have +ve value")
	}
	if num == 0 {
		return "0", nil
	}
	var result string = ""
	for num > 0 {
		result += strconv.Itoa(num & 1)
		num >>= 1
	}
	return Reverse(result), nil
}

// Reverse() function that will take string, and returns the reverse of that string.
func Reverse(str string) string {
	rStr := []rune(str)
	for i, j := 0, len(rStr)-1; i < len(rStr)/2; i, j = i+1, j-1 {
		rStr[i], rStr[j] = rStr[j], rStr[i]
	}
	return string(rStr)
}

//二进制字符串转16位非符合证书
func BinaryToUint16(binary string) (uint16, error) {
	l := len(binary)
	if l != 16 {
		return uint16(0), fmt.Errorf("input param is not uint16 type.")
	}
	var result, base uint16 = 0, 1
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			result += base
		}
		base *= 2
	}
	return result, nil
}

func TimeStampFormate(t uint32) string {
	timeTemplate := "2006-01-02 15:04:05"
	tm := time.Unix(int64(t), 0)
	timeStr := tm.Format(timeTemplate)
	return timeStr
}

func U2S(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}

	return fmt.Sprintf("%c", form), nil
}

func U16To8(u16 []byte) string {
	if len(u16)%2 != 0 {
		return ""
	}

	var body bytes.Buffer

	for i := 0; i < len(u16)/2; i++ {
		v := int(u16[2*i]) + int(u16[2*i+1])<<8
		if v <= 127 {

			body.WriteByte(byte(v))
		} else if v <= 2047 {
			a1 := byte(v&63) + 128

			v = v >> 6
			a2 := byte(v&31) + 192
			body.WriteByte(a2)
			body.WriteByte(a1)

		} else if v <= 65535 {
			a1 := byte(v&63) + 128

			v = v >> 6
			a2 := byte(v&63) + 128

			v = v >> 6
			a3 := byte(v&15) + 224
			body.WriteByte(a3)
			body.WriteByte(a2)
			body.WriteByte(a1)
		}
	}
	return string(body.Bytes())
}

//
//func intToByte(d uint32, size int) []byte {
//	b := make([]byte, 32)
//	binary.BigEndian.PutUint32(b, d)
//	return b
//}

func ByteToBigInt(d []byte) *big.Int {
	b := big.Int{}
	b.SetBytes(d)
	return &b
}

func ByteToIPDaddress(d []byte) string {
	ip := []string{}
	for _, v := range d {
		if uint8(v) != 0 {
			ip = append(ip, strconv.Itoa(int(v)))
		}
	}
	ipa := ""
	if len(ip) <= 4 {
		ipa = strings.Join(ip, ".")
	} else {

		ipa = strings.Join(ip, ":")
	}
	return ipa
}

//16进制字符串转Big.Int
func HexStringToBigInt(s string) *big.Int {
	result, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic(s)
	}
	return result
}

func HexstringToByteArray(hexstr string) []byte {
	data, err := hex.DecodeString(strings.ToUpper(hexstr))
	if err != nil {
		return nil
	}
	return data
}
