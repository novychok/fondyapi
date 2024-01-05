package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/novychok/fondyapi/types"
)

func createSignature(password string, credentials []string) string {
	var builder strings.Builder
	sort.Strings(credentials)
	for _, v := range credentials {
		if v == "" {
			continue
		}
		builder.WriteString("|" + v)
	}
	result := password + builder.String()

	fmt.Printf("res: %s\n", result)
	hash := sha1.Sum([]byte(result))

	return hex.EncodeToString(hash[:])
}

func main() {

	password := "test"
	request := types.APIRequest{
		OrderID:    "test123456",
		OrderDesc:  "test order",
		Currency:   "USD",
		Amount:     "125",
		Signature:  "",
		MerchantID: "1396424",
	}

	sl := make([]string, 6)
	sl[1] = request.OrderID
	sl[2] = request.OrderDesc
	sl[3] = request.Currency
	sl[4] = request.Amount
	sl[5] = request.MerchantID

	request.Signature = createSignature(password, sl)
	fmt.Println(request.Signature)

	if request.Signature != "f0ee6288b9295d3b808bcd8d720211c7201245e1" {
		fmt.Println("not equal")
	}
	a := sha1.Sum([]byte("test|125|USD|1396424|test order|test123456"))
	fmt.Println(hex.EncodeToString(a[:]))
}
