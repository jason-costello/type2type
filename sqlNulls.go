package type2type

import (
	"database/sql"
	"errors"
)

var (
	ErrInvalidNullInt32ToInt32OrError = errors.New("failed to convert from NullInt32 to Int32 or Error")

	ErrInvalidNullInt16ToInt16OrError = errors.New("failed to convert from NullInt16 to Int16 or Error")

	ErrInvalidNullInt64ToInt64OrError = errors.New("failed to convert from NullInt64 to Int64 or Error")

	ErrInvalidNullFloat64ToFloat64OrError = errors.New("failed to convert from NullFloat64 to Float64 or Error")

	ErrInvalidNullStringToStringOrError = errors.New("failed to convert from NullString to String or Error")

	ErrInvalidNullBoolToBoolOrError = errors.New("failed to convert from NullBool to Bool or Error")

	ErrInvalidNullByteToByteOrError = errors.New("failed to convert from NullByte to Byte or Error")
)

func NullInt32ToInt32OrDefault(nt sql.NullInt32) int32 {

	var v int32
	if nt.Valid {
		v = nt.Int32
	}

	return v

}

func NullInt32ToInt32OrPanic(nt sql.NullInt32) int32 {

	var v int32
	if nt.Valid {
		v = nt.Int32
	}

	if !nt.Valid {
		panic(ErrInvalidNullInt32ToInt32OrError)
	}

	return v

}

func NullInt32ToInt32OrError(nt sql.NullInt32) (int32, error) {

	var v int32
	if nt.Valid {
		v = nt.Int32
	}

	if !nt.Valid {
		return v, ErrInvalidNullInt32ToInt32OrError
	}
	return v, nil

}

func NullInt16ToInt16OrDefault(nt sql.NullInt16) int16 {

	var v int16
	if nt.Valid {
		v = nt.Int16
	}

	return v

}

func NullInt16ToInt16OrPanic(nt sql.NullInt16) int16 {

	var v int16
	if nt.Valid {
		v = nt.Int16
	}

	if !nt.Valid {
		panic(ErrInvalidNullInt16ToInt16OrError)
	}

	return v

}

func NullInt16ToInt16OrError(nt sql.NullInt16) (int16, error) {

	var v int16
	if nt.Valid {
		v = nt.Int16
	}

	if !nt.Valid {
		return v, ErrInvalidNullInt16ToInt16OrError
	}
	return v, nil

}

func NullInt64ToInt64OrDefault(nt sql.NullInt64) int64 {

	var v int64
	if nt.Valid {
		v = nt.Int64
	}

	return v

}

func NullInt64ToInt64OrPanic(nt sql.NullInt64) int64 {

	var v int64
	if nt.Valid {
		v = nt.Int64
	}

	if !nt.Valid {
		panic(ErrInvalidNullInt64ToInt64OrError)
	}

	return v

}

func NullInt64ToInt64OrError(nt sql.NullInt64) (int64, error) {

	var v int64
	if nt.Valid {
		v = nt.Int64
	}

	if !nt.Valid {
		return v, ErrInvalidNullInt64ToInt64OrError
	}
	return v, nil

}

func NullFloat64ToFloat64OrDefault(nt sql.NullFloat64) float64 {

	var v float64
	if nt.Valid {
		v = nt.Float64
	}

	return v

}

func NullFloat64ToFloat64OrPanic(nt sql.NullFloat64) float64 {

	var v float64
	if nt.Valid {
		v = nt.Float64
	}

	if !nt.Valid {
		panic(ErrInvalidNullFloat64ToFloat64OrError)
	}

	return v

}

func NullFloat64ToFloat64OrError(nt sql.NullFloat64) (float64, error) {

	var v float64
	if nt.Valid {
		v = nt.Float64
	}

	if !nt.Valid {
		return v, ErrInvalidNullFloat64ToFloat64OrError
	}
	return v, nil

}

func NullStringToStringOrDefault(nt sql.NullString) string {

	var v string
	if nt.Valid {
		v = nt.String
	}

	return v

}

func NullStringToStringOrPanic(nt sql.NullString) string {

	var v string
	if nt.Valid {
		v = nt.String
	}

	if !nt.Valid {
		panic(ErrInvalidNullStringToStringOrError)
	}

	return v

}

func NullStringToStringOrError(nt sql.NullString) (string, error) {

	var v string
	if nt.Valid {
		v = nt.String
	}

	if !nt.Valid {
		return v, ErrInvalidNullStringToStringOrError
	}
	return v, nil

}

func NullBoolToBoolOrDefault(nt sql.NullBool) bool {

	var v bool
	if nt.Valid {
		v = nt.Bool
	}

	return v

}

func NullBoolToBoolOrPanic(nt sql.NullBool) bool {

	var v bool
	if nt.Valid {
		v = nt.Bool
	}

	if !nt.Valid {
		panic(ErrInvalidNullBoolToBoolOrError)
	}

	return v

}

func NullBoolToBoolOrError(nt sql.NullBool) (bool, error) {

	var v bool
	if nt.Valid {
		v = nt.Bool
	}

	if !nt.Valid {
		return v, ErrInvalidNullBoolToBoolOrError
	}
	return v, nil

}

func NullByteToByteOrDefault(nt sql.NullByte) byte {

	var v byte
	if nt.Valid {
		v = nt.Byte
	}

	return v

}

func NullByteToByteOrPanic(nt sql.NullByte) byte {

	var v byte
	if nt.Valid {
		v = nt.Byte
	}

	if !nt.Valid {
		panic(ErrInvalidNullByteToByteOrError)
	}

	return v

}

func NullByteToByteOrError(nt sql.NullByte) (byte, error) {

	var v byte
	if nt.Valid {
		v = nt.Byte
	}

	if !nt.Valid {
		return v, ErrInvalidNullByteToByteOrError
	}
	return v, nil

}
