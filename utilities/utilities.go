package utilities

import (
	"database/sql"
	"fmt"
	"strconv"
)

// SetUpMiddleware contains the middleware that applies to every request.
func DisplayPrettyFloat(f float32) string {
	return fmt.Sprintf("%.2f", f)
}

func DisplayPrettyFloat64(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

func DisplayPrettyNullFloat64(sqlf sql.NullFloat64) string {
	//return "Hello"
	return fmt.Sprintf("%.2f", sqlf.Float64)
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsPositiveInteger(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}
