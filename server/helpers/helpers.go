package helpers

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

var (
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func LogFatal(i interface{}) {
	var err error

	switch i.(type) {
	case error:
	case fmt.Stringer:
		err = errors.New(i.(fmt.Stringer).String())
	case string:
		err = errors.New(i.(string))
	default:
		msg := fmt.Sprintf("log fatal: %v")
		err = errors.New(msg)
	}

	LogError(err)
	os.Exit(1)
}

func DebugInterface(i interface{}) (v reflect.Value, t reflect.Type, ts string) {
	v = reflect.ValueOf(i)
	t = reflect.TypeOf(i)
	ts = t.String()
	return
}

func DebugError(err error) (msg string) {
	msg = err.Error()
	return
}

func LogError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v\n", fn, line, err)
}

func Dump(i interface{}) {
	fmt.Printf("%+v\n", i)
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// need struct, not ptr - otherwise func panics
func StructToMap(input interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	s := reflect.ValueOf(input)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		r[typeOfT.Field(i).Name] = f.Interface()
	}
	return r
}

func ParseTime(d string) (t time.Time, err error) {
	t, err = time.Parse("2006-01-02", d)
	if err != nil {
		LogError(err)
	}
	return
}

func GetDefaultJWTExpiresAt() time.Time {
	return time.Now().Add(time.Duration(365) * 24 * time.Hour)
}

func GenerateS3FilePath(prefix, fileExtension string) string {

	currentTime := time.Now()
	year := fmt.Sprintf("%04d", currentTime.Year())
	month := fmt.Sprintf("%02d", currentTime.Month())
	day := fmt.Sprintf("%02d", currentTime.Day())

	// Step 2: Generate a unique hash value.
	hash := GenerateRandomString(18)

	// Step 3: Combine the date and hash to create the file path.
	filePath := filepath.Join(prefix, year, month, day, hash+fileExtension)

	return filePath
}
