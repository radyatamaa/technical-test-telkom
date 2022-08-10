package helper

import (
	"database/sql/driver"
	"encoding/json"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"technical-test-telkom/models"
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(24*time.Hour, 24*time.Hour)

type Emp []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SetCache(key string, emp interface{}) bool {
	Cache.Set(key, emp, cache.NoExpiration)
	return true
}

func GetCache(key string) (string, bool) {
	var emp string
	var found bool
	data, found := Cache.Get(key)
	if found {
		emp = data.(string)
	}

	return emp, found
}

func Pagination(qpage, qperPage string) (limit, page, offset int) {
	limit = 20
	page = 1
	offset = 0

	page, _ = strconv.Atoi(qpage)
	limit, _ = strconv.Atoi(qperPage)
	if page == 0 && limit == 0 {
		page = 1
		limit = 10
	}
	offset = (page - 1) * limit

	return
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrUnAuthorize:
		return http.StatusUnauthorized
	case models.ErrConflict:
		return http.StatusConflict
	case models.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func GetKeyJsonStruct(value interface{}) []string {

	j, _ := json.Marshal(value)
	// a map container to decode the JSON structure into
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		i++
	}

	return k
}

func GetValueStruct(value interface{}) []driver.Value {
	var result []driver.Value
	rv := reflect.ValueOf(value)
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)

		dv := driver.Value(fv.Interface())
		result = append(result, dv)
	}
	return result
}

func GetValueAndColumnStructToDriverValue(value interface{}) ([]driver.Value, []string) {
	var result []driver.Value

	//column
	j, _ := json.Marshal(value)
	// a map container to decode the JSON structure into
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, e := range c {
		k[i] = s
		v, _ := e.MarshalJSON()
		var val driver.Value
		err := json.Unmarshal(v, &val)
		if err != nil {
			panic(err)
		}

		//dv := driver.Value(fv.Interface())
		result = append(result, val)
		i++
	}

	return result, k

}

func NowYmd() string {
	t := time.Now()
	timeFormated := t.Format("2006-01-02 15:04:05")
	return timeFormated
}
func FloatToString(input_num float64) string {
	// to convert a float number to a string
	if input_num != 0 {
		return strconv.FormatFloat(input_num, 'f', 0, 64)
	} else {
		return "0"
	}
}

func StringToFloat(value *string) float64 {
	if value != nil {
		res, _ := strconv.ParseFloat(*value, 64)
		return res
	}
	return 0
}
func FloatNUllableToString(input_num *float64) string {
	// to convert a float number to a string
	if input_num != nil {
		return strconv.FormatFloat(*input_num, 'f', 0, 64)
	} else {
		return ""
	}
}
func FloatNUllableToFloat(value *float64) float64 {
	if value != nil {
		return *value
	}
	return 0
}
func DateTimeToDateTimeNUllable(value time.Time) *time.Time {
	return &value
}
func IntToIntNullable(value int) *int {
	return &value
}
func IntNullableToInt(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}
func StringToStringNullable(value string) *string {
	return &value
}
func StringNullableToString(value *string) string {
	if value != nil {
		return *value
	}
	return ""
}
func IntNullableToStringNullable(value *int) *string {

	if value != nil {
		result := strconv.Itoa(*value)
		return &result
	}
	return nil
}
func IntNullableToString(value *int) string {

	if value != nil {
		result := strconv.Itoa(*value)
		return result
	}
	return "0"
}
func StringToIntNullable(value string) *int {

	if value != "" {
		result, _ := strconv.Atoi(value)
		return &result
	}
	return nil
}
func StringNullableToInt(value *string) int {

	if value != nil {
		result, _ := strconv.Atoi(*value)
		return result
	}
	return 0
}
func StringNullableToDateTimeNullable(value *string) *time.Time {
	if value != nil {
		var layoutFormat string
		var date time.Time

		layoutFormat = "2006-01-02 15:04:05"
		date, _ = time.Parse(layoutFormat, *value)
		return &date
	}

	return nil
}

func DateTimeNullableToStringNullable(value *time.Time) *string {
	if value != nil {
		layoutFormat := "2006-01-02 15:04:05"
		date := value.Format(layoutFormat)
		return &date
	}

	return nil
}

func StringNullableToStringDefaultFormatDate(value *string) *string {
	if value != nil {
		var layoutFormat string
		var date time.Time

		layoutFormat = "2006-01-02T15:04:05Z"
		date, _ = time.Parse(layoutFormat, *value)
		dateString := date.Format("2006-01-02 15:04:05")
		return &dateString
	}

	return nil
}
func StringToDateTimeNullable(value string) *time.Time {
	if value != "" {
		var layoutFormat string
		var date time.Time

		layoutFormat = "2006-01-02 15:04:05"
		date, _ = time.Parse(layoutFormat, value)
		return &date
	}

	return nil
}
func StringToDate(value string) time.Time {
	if value != "" {
		var layoutFormat string
		var date time.Time

		layoutFormat = "2006-01-02"
		date, _ = time.Parse(layoutFormat, value)
		return date
	}

	return time.Time{}
}
func StringNullableToDateNullable(value *string) *string {
	if value != nil {
		var layoutFormat string
		var date time.Time

		layoutFormat = "20060102"
		date, _ = time.Parse(layoutFormat, *value)
		dateString := date.Format("20060102")
		return &dateString
	}

	return nil
}
func ConvertIntBool(value *int) bool {
	if value != nil {
		if *value == 1 {
			return true
		}
	}
	return false
}
func NowAddDay() string {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = NowYmd()
	date, _ = time.Parse(layoutFormat, value)

	return date.AddDate(0, 0, 1).Format(layoutFormat)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}
