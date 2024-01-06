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

	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/novychok/fondyapi/types"
)

var password = "test"

type APIRequest struct {
	Request types.Request `json:"request"`
}

type APIResponse struct {
	Response any `json:"response"`
}

func NewAPIRequest(orderId, orderDesc, currency, amount,
	merchantId string) *APIRequest {
	return &APIRequest{
		Request: types.Request{
			OrderID:           orderId,
			OrderDesc:         orderDesc,
			Currency:          currency,
			Amount:            amount,
			Signature:         "",
			MerchantID:        merchantId,
			ServerCallbackURL: "https://11ee-37-252-93-141.ngrok-free.app",
		},
	}
}

func (a *APIRequest) SetSignature(password string) {
	fieldsMap := structs.Map(&a.Request)
	var fieldsSl []string
	var builder strings.Builder

	for val := range fieldsMap {
		fieldsSl = append(fieldsSl, val)
	}
	sort.Strings(fieldsSl)

	for _, val := range fieldsSl {
		if v, ok := fieldsMap[val]; ok {
			if v == "" {
				continue
			}
			builder.WriteString("|" + v.(string))
		}
	}
	formatingStr := password + builder.String()

	a.Request.Signature = a.calculateSignature(formatingStr)
}

func (a APIRequest) calculateSignature(s string) string {
	hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func main() {
	orderId := uuid.New().String()
	apiRequest := NewAPIRequest(orderId, "111 order", "USD",
		"1250", "1396424")

	apiRequest.SetSignature(password)
	// apiRequest.Request.ServerCallbackURL = "https://7143-37-252-93-141.ngrok-free.app"

	request := APIRequest{Request: apiRequest.Request}
	body, _ := json.Marshal(request)
	res, err := http.Post("https://pay.fondy.eu/api/checkout/url", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("error to make a request: %+v\n", err)
		return
	}
	defer res.Body.Close()

	var response APIResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		fmt.Printf("error to decode the response 111: %+v\n", err.Error())
		return
	}
	fmt.Printf("%+v\n", response)
}
