package commons

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

/**
md5加密
sha1加密
*/

func EncodeMd5(value string) string {

	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func Sha1(data string) string {
	sha := sha1.New()
	sha.Write([]byte(data))
	return hex.EncodeToString(sha.Sum([]byte("")))
}

/**
字符串转数字
*/

func StrToInt(s string) (int, bool) {
	i, err := strconv.Atoi(s)
	if CheckErr(err, "") {
		return i, true
	}
	return 0, false

}

/**
随机字符串
*/
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

/**
字符串拼接
*/
func StringJoin(a ...string) string {
	var buf bytes.Buffer
	for _, k := range a {
		buf.WriteString(k)
	}
	return buf.String()
}

/**
简单的waf
*/
func StringFilter(s string, length int) bool {
	v := strings.TrimSpace(s)
	match, _ := regexp.MatchString(".*?<\\?|\\?>|>|\\(", v)

	if match {
		return false
	}
	if v != "" && len(v) >= length {
		return true
	}
	return false
}

/**
判断字符串是否为空
*/
func StringIsEmpty(s string, length int) bool {
	if strings.TrimSpace(s) != "" {
		if StringFilter(s, length) {
			return true
		}
	}
	return false
}

func BinarySearch(s []string, m string) int {
	//fmt.Printf("%v , %v", s, m)
	//针对对字符串
	if !StringFilter(m, 1) {
		return -1
	}

	for k, v := range s {
		if strings.EqualFold(v, m) {
			return k
		}
	}

	////_k := m[1 : len(m)-1]
	////k, _ := strconv.Atoi(m)
	//ints := sort.SearchStrings(s, m)
	//if ints == len(s) {
	//	return -1
	//} else {
	//	return ints
	//}
	////lo, hi := 0, len(s)-1
	////for lo <= hi {
	////	m := (lo + hi) >> 1
	////	if s[m] < k {
	////		lo = m + 1
	////	} else if s[m] > k {
	////		hi = m - 1
	////	} else {
	////		return m
	////	}
	////}
	return -1
}

func Abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func GetRawBody(r *http.Request) ([]byte, error) {
	body := r.Body
	defer body.Close()
	rawBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return rawBody, nil
}

func Signature(appId, method string, body []byte, RequestURL string, timestamp string) (result string) {
	stringToSign := fmt.Sprintf("%v\n%v\n%v\n%v\n", method, string(body), RequestURL, timestamp)
	// fmt.Println(1111, stringToSign)
	s := sha256.New
	hash := hmac.New(s, []byte(appId))
	hash.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

/**

 */
func HttpRequest(url, params, data, method string, interval int, header map[string]string) {
	//url := "https://pipenv.com/admin/api/v1/login"
	time.Sleep(time.Duration(interval))
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Duration(interval),
			}).DialContext,
		},
	}
	payload := strings.NewReader(params)
	req, _ := http.NewRequest(method, url, payload)

	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)

		}
	}

	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func GetAfterDay(day int) string {
	a := time.Now()
	location, _ := time.LoadLocation("Local")

	return a.In(location).AddDate(0, 0, day).Format("2006-01-02 15:04:05")
}

func CompareTime(date string) bool {
	a := time.Now()
	location, _ := time.LoadLocation("Local")

	p, _ := time.ParseInLocation("2006-01-02 15:04:05", date, location)

	if a.In(location).Unix()-p.In(location).Unix() >= 0 {
		return true
	}
	return false
}

func SToM(obj interface{}) (string, []interface{}) {
	slice := make([]interface{}, 0)
	sql := make([]interface{}, 0)
	query := ""

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		switch t.Field(i).Type.Kind() {
		case reflect.String:
			if v.Field(i).Interface() != "" {
				sql = append(sql, t.Field(i).Name)
				slice = append(slice, v.Field(i).Interface())
			}
		case reflect.Int:
			if v.Field(i).Interface() != 0 {
				sql = append(sql, t.Field(i).Name)
				slice = append(slice, v.Field(i).Interface())
			}
		}

	}

	for k, v := range sql {
		if k == len(sql)-1 {
			query += fmt.Sprintf(" u.%s = ? ", v)
		} else {
			query += fmt.Sprintf(" u.%s = ? and ", v)
		}
	}
	return query, slice
}
