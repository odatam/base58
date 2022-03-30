package base58

import (
	"bytes"
	"errors"	
	"math/big"
)

// base58 encode/decode inspired & simplified from https://github.com/m0t0k1ch1/base58
func Base58EncodeToString(b []byte) (string, error) {
	var chars = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	
	n := &big.Int{}
	n.SetBytes(b)

	buf := &bytes.Buffer{}
	for _, c := range b {
		if c == 0x00 {
			if err := buf.WriteByte(chars[0]); err != nil {
				return "", err
			}
		} else {
			break
		}
	}

	zero := big.NewInt(0)
	div := big.NewInt(58)
	mod := &big.Int{}

	tmpBuf := &bytes.Buffer{}
	for {
		if n.Cmp(zero) == 0 {
			tmpBytes := tmpBuf.Bytes()
			for i := len(tmpBytes) - 1; i >= 0; i-- {
				buf.WriteByte(tmpBytes[i])
			}
			return buf.String(), nil
		}

		n.DivMod(n, div, mod)
		if err := tmpBuf.WriteByte(chars[mod.Int64()]); err != nil {
			return "", err
		}
	}
}

func Base58DecodeToString(s string) ([]byte, error) {
	var charIndex = map[byte]int64{49:0, 50:1, 51:2, 52:3, 53:4, 54:5, 55:6, 56:7, 57:8, 65:9, 66:10, 67:11, 68:12, 69:13, 70:14, 71:15, 72:16, 74:17, 75:18, 76:19, 77:20, 78:21, 80:22, 81:23, 82:24, 83:25, 84:26, 85:27, 86:28, 87:29, 88:30, 89:31, 90:32, 97:33, 98:34, 99:35, 100:36, 101:37, 102:38, 103:39, 104:40, 105:41, 106:42, 107:43, 109:44, 110:45, 111:46, 112:47, 113:48, 114:49, 115:50, 116:51, 117:52, 118:53, 119:54, 120:55, 121:56, 122:57}
	
	b := []byte(s)
	startIdx := 0
	buf := &bytes.Buffer{}
	for i, c := range b {
		if c == byte(49) {
			if err := buf.WriteByte(0x00); err != nil {
				return nil, err
			}
		} else {
			startIdx = i
			break
		}
	}

	n := big.NewInt(0)
	div := big.NewInt(58)

	for _, c := range b[startIdx:] {
		currentCharIndex, ok := charIndex[c]
		if !ok {
			return nil, errors.New("invalid char")
		}

		n.Add(n.Mul(n, div), big.NewInt(currentCharIndex))
	}

	buf.Write(n.Bytes())

	return buf.Bytes(), nil
}
