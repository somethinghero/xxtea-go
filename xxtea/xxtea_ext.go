package xxtea;

import(
	"math/rand"
	"time"
	"fmt"
	"errors"
)

func fixData(data []byte) []byte {
	dataLen := len(data)
	fixByteCount := 4 - ((dataLen + 1) % 4)
	fixedData := make([]byte, dataLen + fixByteCount + 1)

	fixedData[0] = (byte)(fixByteCount)
	for i := 0; i < fixByteCount; i++ {
		fixedData[i + 1] = (byte)(rand.Intn(255))
	}
	copy(fixedData[fixByteCount + 1:], data[:])
	return fixedData
}

func unfixData(data []byte) []byte {
	fixByteCount := (int)(data[0])
	realLen := len(data) - fixByteCount - 1
	realData := make([]byte, realLen)
	copy(realData, data[fixByteCount + 1:])
	return realData
}

func init(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println("xxtea_ext init")
}

func EncryptExt(data []byte, key []byte) []byte{
	fixed := fixData(data)
	return Encrypt(fixed, key)
}

func DecryptExt(data []byte, key []byte) ([]byte, error) {
	if data == nil || len(data) == 0 || key == nil || len(key) == 0 {
		return nil, errors.New("para error")
	}
	fixed := Decrypt(data, key)
	if nil == fixed {
		return nil, errors.New("decrypt error")
	}
	return unfixData(fixed), nil
}