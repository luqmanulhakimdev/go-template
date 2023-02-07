package str

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowString(t *testing.T) {
	res := ShowString(true, "showString")
	assert.Equal(t, res, "showString")

	res = ShowString(false, "hideString")
	assert.Equal(t, res, "")
}

func TestEmptyString(t *testing.T) {
	res := EmptyString("")
	var pointerVar *string
	pointerVar = nil
	assert.Equal(t, res, pointerVar)

	res = EmptyString("string")
	stringVar := "string"
	pointerVar = &stringVar
	assert.Equal(t, res, pointerVar)
}

func TestEmptyInt(t *testing.T) {
	res := EmptyInt(0)
	var pointerVar *int
	pointerVar = nil
	assert.Equal(t, res, pointerVar)

	res = EmptyInt(123)
	intVar := 123
	pointerVar = &intVar
	assert.Equal(t, res, pointerVar)
}

func TestEmptyFloat(t *testing.T) {
	res := EmptyFloat(0)
	var pointerVar *float64
	pointerVar = nil
	assert.Equal(t, res, pointerVar)

	res = EmptyFloat(123)
	floatVar := float64(123)
	pointerVar = &floatVar
	assert.Equal(t, res, pointerVar)
}

func TestStringToInt(t *testing.T) {
	res := StringToInt("123")
	assert.Equal(t, res, 123)

	res = StringToInt("abc")
	assert.Equal(t, res, 0)
}

func TestStringToBool(t *testing.T) {
	res := StringToBool("true")
	assert.Equal(t, res, true)
}

func TestStringToBoolString(t *testing.T) {
	res := StringToBoolString("true")
	assert.Equal(t, res, "true")

	res = StringToBoolString("truee")
	assert.Equal(t, res, "false")
}

func TestRandomString(t *testing.T) {
	res := RandomString(4)
	assert.IsType(t, res, "abcd")
}

func TestIsActive(t *testing.T) {
	res := IsActive("true")
	var pointerVar *string
	stringVar := "and is_active = 'true'"
	pointerVar = &stringVar
	assert.Equal(t, res, pointerVar)

	res = IsActive("false")
	stringVar = "and is_active = 'false'"
	pointerVar = &stringVar
	assert.Equal(t, res, pointerVar)

	res = IsActive("falsee")
	stringVar = ""
	pointerVar = &stringVar
	assert.Equal(t, res, pointerVar)
}

func TestUnique(t *testing.T) {
	res := Unique([]string{"abc", "abc", "def"})
	assert.Equal(t, res, []string{"abc", "def"})
}

func TestCheckEmail(t *testing.T) {
	res := CheckEmail("neobank@mail.com")
	assert.Equal(t, res, true)
}

func TestIsValidUUID(t *testing.T) {
	res := IsValidUUID("e5ef9c0b-87c7-4849-a5f6-e7a25d73a4f8")
	assert.Equal(t, res, true)
}

func TestIntNilChecker(t *testing.T) {
	var pointerVar *int
	intVar := 123
	pointerVar = &intVar
	res := IntNilChecker(pointerVar)
	assert.Equal(t, res, intVar)

	pointerVar = nil
	res = IntNilChecker(pointerVar)
	assert.Equal(t, res, 0)
}

func TestFloatNilChecker(t *testing.T) {
	var pointerVar *float64
	floatVar := float64(123)
	pointerVar = &floatVar
	res := FloatNilChecker(pointerVar)
	assert.Equal(t, res, floatVar)

	pointerVar = nil
	res = FloatNilChecker(pointerVar)
	assert.Equal(t, res, float64(0))
}

func TestStringNilChecker(t *testing.T) {
	var pointerVar *string
	stringVar := "abcdefgh"
	pointerVar = &stringVar
	res := StringNilChecker(pointerVar)
	assert.Equal(t, res, stringVar)

	pointerVar = nil
	res = StringNilChecker(pointerVar)
	assert.Equal(t, res, "")
}

func TestJoinWithQuotes(t *testing.T) {
	res := JoinWithQuotes([]string{"abc", "def"}, "'", "_")
	assert.Equal(t, res, "'abc'_'def'")
}

func TestCheckNumber(t *testing.T) {
	res := CheckNumber("123")
	assert.Equal(t, res, true)
}

func TestCheckBulkInt(t *testing.T) {
	res := CheckBulkInt([]string{"123"})
	assert.Equal(t, res, true)

	res = CheckBulkInt([]string{"abc"})
	assert.Equal(t, res, false)
}

func TestConvertStringPointer(t *testing.T) {
	var pointerVar *string
	stringVar := "abcdefgh"
	pointerVar = &stringVar
	res := ConvertStringPointer(pointerVar)
	assert.Equal(t, res, stringVar)

	pointerVar = nil
	res = ConvertStringPointer(pointerVar)
	assert.Equal(t, res, "")
}

func TestNormalizePhone(t *testing.T) {
	res := NormalizePhone("0812345")
	assert.Equal(t, res, "")

	res = NormalizePhone("628123456789")
	assert.Equal(t, res, "08123456789")

	res = NormalizePhone("08123456789")
	assert.Equal(t, res, "08123456789")
}

func TestMultipleValueParameter(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com?key=", nil)
	res := MultipleValueParameter(req, "key")
	assert.Equal(t, res, []string{})

	req, _ = http.NewRequest("GET", "http://example.com?key=abc,def", nil)
	res = MultipleValueParameter(req, "key")
	assert.Equal(t, res, []string{"abc", "def"})
}

func TestArrayStringToInt(t *testing.T) {
	res := ArrayStringToInt([]string{"123"})
	assert.Equal(t, res, []int{123})

	res = ArrayStringToInt([]string{"abc"})
	assert.Equal(t, res, []int{})
}

func TestGenerateTransactionCode(t *testing.T) {
	res := GenerateTransactionCode("trx", 1)
	assert.IsType(t, res, "trxcode1")
}

func TestArrayStringToString(t *testing.T) {
	res := ArrayStringToString([]string{"type1", "type2"}, ",")
	assert.Equal(t, res, "type1,type2")
}

func TestArrayIntToString(t *testing.T) {
	res := ArrayIntToString([]int{1, 2}, ",")
	assert.Equal(t, res, "1,2")
}
