package libs

import (
	"crypto/dsa"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/wangxiaojie1005/hdlclient-sdk/units"
	"strconv"
)

type HSPubkey struct {
	Name  string
	Value interface{}
}

func (k *HSPubkey) GetPubKey(d []byte) {
	offset := 0
	keyTypeLen := units.BytesToUint32(d[offset : offset+4])
	offset += 4
	keyType := units.Byte2string(d[offset : offset+int(keyTypeLen)])
	offset += int(keyTypeLen)
	offset += 2
	switch keyType {
	case "DSA_PUB_KEY":
		dk := HSDSAPubkey{}

		q_len := units.BytesToUint32(d[offset : offset+4])
		offset += 4
		//fmt.Printf("<<<<<<<><><><> %v\n", hex.EncodeToString(d[offset:offset+int(q_len)]))
		dk.Q = hex.EncodeToString(d[offset : offset+int(q_len)])
		offset += int(q_len)

		p_len := units.BytesToUint32(d[offset : offset+4])
		offset += 4
		dk.P = hex.EncodeToString(d[offset : offset+int(p_len)])
		offset += int(p_len)

		g_len := units.BytesToUint32(d[offset : offset+4])
		offset += 4
		dk.G = hex.EncodeToString(d[offset : offset+int(g_len)])
		offset += int(g_len)

		y_len := units.BytesToUint32(d[offset : offset+4])
		offset += 4
		dk.Y = hex.EncodeToString(d[offset : offset+int(y_len)])
		offset += int(y_len)

		k.Name = "DSA_PUB_KEY"
		k.Value = dk

	case "DH_PUB_KEY":
		println("")
	case "RSA_PUB_KEY":
		rk := HSRSAPubkey{}
		n_len := units.BytesToUint32(d[offset : offset+4])
		offset += 4
		rk.N = hex.EncodeToString(d[offset : offset+int(n_len)])
		offset += int(n_len)

		e_len := units.BytesToUint32(d[offset : offset+4])
		offset += 4
		rk.E = hex.EncodeToString(d[offset : offset+int(e_len)])
		offset += int(e_len)
		k.Name = "RSA_PUB_KEY"
		k.Value = rk
	}
}

func (k *HSPubkey) ToString() {
	data, _ := json.MarshalIndent(k, "", "\t")
	fmt.Printf("Recv MessageEnvelope:%s\n", data)
}

type HSDSAPubkey struct {
	P string
	Q string
	G string
	Y string
}

func (d *HSDSAPubkey) New(p string, q string, g string, y string) *HSDSAPubkey {
	dk := HSDSAPubkey{
		P: p,
		Q: q,
		G: g,
		Y: y,
	}
	return &dk
}

func (d *HSDSAPubkey) GetDSAPubKey() *dsa.PublicKey {
	dk := dsa.PublicKey{}
	dk.P = units.HexStringToBigInt(d.P)
	dk.Q = units.HexStringToBigInt(d.Q)
	dk.G = units.HexStringToBigInt(d.G)
	dk.Y = units.HexStringToBigInt(d.Y)
	return &dk
}

type HSRSAPubkey struct {
	N string
	E string
}

func (r *HSRSAPubkey) New(n string, e string) *HSRSAPubkey {
	rk := HSRSAPubkey{
		N: n,
		E: e,
	}
	return &rk
}

func (r *HSRSAPubkey) GetRSAPubKey() *rsa.PublicKey {
	rk := rsa.PublicKey{}
	rk.N = units.HexStringToBigInt(r.N)
	a, _ := strconv.Atoi(r.E)
	rk.E = a
	return &rk
}
