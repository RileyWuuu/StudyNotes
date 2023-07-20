package Encryption

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParamsEncrypt(t *testing.T) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	// set params
	currentTime := time.Now().Format("20060102150405000")
	orderId := currentTime + "test123"
	params := url.Values{}
	params.Set("s", "0")
	params.Set("account", "test123")
	params.Set("money", "1000")
	params.Set("orderid", orderId)
	params.Set("ip", "127.0.0.1")
	params.Set("lineCode", "ABC")
	params.Set("kindID", "0")
	fmt.Println("params:", params.Encode())
	fmt.Println("timestamp:", timestamp)
	// testMsg := "hello world"
	encrypted, err := EncryptMessage([]byte(params.Encode()), []byte(desKey))
	// encrypted, err := EncryptMessage([]byte(testMsg), []byte(desKey))
	if err != nil {
		fmt.Println("AESEncrypt err:", err.Error())
		return
	}
	fmt.Println("decrypt:", encrypted)

	decrypted, err := DecryptMessage([]byte(desKey), encrypted)
	if err != nil {
		fmt.Println("DecryptMessage err:", err.Error())
		return
	}
	require.Nil(t, err)
	require.Equal(t, params.Encode(), decrypted)
}
