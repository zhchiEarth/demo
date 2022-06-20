package data

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"strconv"
	"testing"
)

func TestClient(t *testing.T) {
	v := 3.1415926535
	fmt.Println(strconv.FormatFloat(v, 'e', -1, 32))
	fmt.Println(strconv.FormatFloat(v, 'E', -1, 64))
	fmt.Println(strconv.FormatFloat(v, 'f', -1, 64))
}

func TestLog(t *testing.T) {
	_, err := decimal.NewFromString("")
	fmt.Printf("%+v\n", errors.Wrap(err, "sss"))
}
