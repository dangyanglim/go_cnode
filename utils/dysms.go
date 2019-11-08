package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// SendSmsReply 发送短信返回
type SendSmsReply struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

func replace(in string) string {
	rep := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	return rep.Replace(url.QueryEscape(in))
}

// SendSms 发送短信
func SendSms(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode string) error {
	paras := map[string]string{
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureNonce":   fmt.Sprintf("%d", rand.Int63()),
		"AccessKeyId":      accessKeyID,
		"SignatureVersion": "1.0",
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           "JSON",

		"Action":        "SendSms",
		"Version":       "2017-05-25",
		"RegionId":      "cn-hangzhou",
		"PhoneNumbers":  phoneNumbers,
		"SignName":      signName,
		"TemplateParam": templateParam,
		"TemplateCode":  templateCode,
	}

	var keys []string

	for k := range paras {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var sortQueryString string

	for _, v := range keys {
		sortQueryString = fmt.Sprintf("%s&%s=%s", sortQueryString, replace(v), replace(paras[v]))
	}

	stringToSign := fmt.Sprintf("GET&%s&%s", replace("/"), replace(sortQueryString[1:]))

	mac := hmac.New(sha1.New, []byte(fmt.Sprintf("%s&", accessSecret)))
	mac.Write([]byte(stringToSign))
	sign := replace(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	str := fmt.Sprintf("http://dysmsapi.aliyuncs.com/?Signature=%s%s", sign, sortQueryString)

	resp, err := http.Get(str)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ssr := &SendSmsReply{}

	if err := json.Unmarshal(body, ssr); err != nil {
		return err
	}

	if ssr.Code == "SignatureNonceUsed" {
		return SendSms(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode)
	} else if ssr.Code != "OK" {
		return errors.New(ssr.Code)
	}

	return nil
}
