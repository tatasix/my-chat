package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"math/rand"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetExcelDataByPath(filename string) ([][]string, error) {
	fileHandler, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	rows := fileHandler.GetRows("Sheet1")
	return rows, err
}

func GetCSVDataByPath(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 将CSV文件的字节数据转换为io.Reader接口
	f := transform.NewReader(file, simplifiedchinese.GBK.NewDecoder())

	reader := csv.NewReader(f)
	result := make([][]string, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, record)
	}
	return result, nil
}

func GenerateSnowflakeInt64() int64 {
	snow, _ := GenerateSnowflake()

	return snow.Int64()
}

func GenerateSnowflakeString() string {
	snow, _ := GenerateSnowflake()

	return snow.String()
}

func MD5(str string) string {
	md5Data := md5.Sum([]byte(str))
	return hex.EncodeToString(md5Data[:])
}

func GenerateSnowflake() (snowflake.ID, error) {
	node, errNode := snowflake.NewNode(1)
	if errNode != nil {
		return 0, errNode
	}
	return node.Generate(), nil
}
func Unique[T comparable](s []T) []T {
	m := make(map[T]bool)
	uniq := make([]T, 0)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = true
			uniq = append(uniq, v)
		}
	}
	return uniq
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StrToMSS(s string) map[string]string {
	if s == "" {
		return nil
	}
	fields := strings.Split(s, "&")
	mss := make(map[string]string)
	for _, field := range fields {
		if strings.Contains(field, "=") {
			keyValue := strings.Split(field, "=")
			key, _ := url.PathUnescape(keyValue[0])
			value := ""
			if len(keyValue) == 2 {
				value, _ = url.PathUnescape(keyValue[1])
				value = strings.ReplaceAll(value, "+", " ")
			}
			mss[key] = value
		}
	}
	return mss
}

func JsonToMSS(s string) map[string]string {
	if s == "" {
		return nil
	}
	msi := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &msi)
	if err != nil {
		return nil
	}
	mss := make(map[string]string)
	for k, v := range msi {
		mss[k] = convertAnyToStr(v)
	}
	return mss
}

// 将任意类型转string
func convertAnyToStr(v interface{}) string {
	if v == nil {
		return ""
	}
	switch d := v.(type) {
	case string:
		return d
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10)
	case []byte:
		return string(d)
	case float32, float64:
		return strconv.FormatFloat(reflect.ValueOf(v).Float(), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(d)
	default:
		return fmt.Sprint(v)
	}
}

func CheckMobile(mobile string) bool {
	reg := `^1\d{10}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}

// GetRandomBoth 获取随机数  数字和文字
func GetRandomBoth(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomNum 获取随机数  纯数字
func GetRandomNum(n int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// Sha1En sha1加密
func Sha1En(data string) string {
	t := sha1.New() ///产生一个散列值得方式
	_, _ = io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func Desensitize(str string) string {
	strLen := len(str)
	if strLen < 8 {
		var asterisks string
		for i := 1; i < strLen; i++ {
			asterisks += "*"
		}
		return str[:1] + asterisks
	}

	maskedStr := str[:4]
	for i := 4; i < strLen-4; i++ {
		if len(maskedStr) > 8 {
			break
		}
		maskedStr += "*"
	}
	maskedStr += "****" + str[strLen-4:]

	return maskedStr
}

func GenerateRandomNumber(min, max int) uint32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rInt := r.Intn(max-min+1) + min
	return uint32(rInt)
}

func NowTimeFormat() string {
	return time.Now().Format("2006-01-02")
}

/** 加密方式 **/

func Md5ByString(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func ConvertToInt64(v interface{}) int64 {
	switch value := v.(type) {
	case string:
		a, _ := strconv.Atoi(value)
		return int64(a)

	}
	return 0
}

func GetGender(g int64) string {
	if g == 1 {
		return "男"
	} else if g == 2 {
		return "女"
	}
	return "未知"
}
