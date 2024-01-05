package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
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
	hash := sha1.Sum([]byte(result))

	return hex.EncodeToString(hash[:])
}

func main() {

	// password := "test"
	sign := sha1.Sum([]byte("test|125|USD|1396424|test order|abbb"))
	tostr := hex.EncodeToString(sign[:])
	request := types.Request{
		OrderID:     "abbb",
		OrderDesc:   "test order",
		Currency:    "USD",
		Amount:      "125",
		Signature:   tostr,
		MerchantID:  "1396424",
		CallbackURL: "https://3bc8-37-252-93-141.ngrok-free.app",
	}

	sl := make([]string, 6)
	sl[1] = request.OrderID
	sl[2] = request.OrderDesc
	sl[3] = request.Currency
	sl[4] = request.Amount
	sl[5] = request.MerchantID

	// request.Signature = createSignature(password, sl)

	req := types.APIRequest{Request: request}
	requestBody, _ := json.Marshal(req)
	res, err := http.Post("https://pay.fondy.eu/api/checkout/url/", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("error to make a request: %+v\n", err)
		return
	}
	defer res.Body.Close()

	var response types.APIResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		fmt.Printf("error to decode the response: %+v\n", err)
		return
	}
	fmt.Printf("%+v\n", response)
}
