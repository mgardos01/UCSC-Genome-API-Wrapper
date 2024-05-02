package models
import (
    "errors"
)
type Format int

const (
    TEXT_FORMAT Format = iota
)

func (i Format) String() string {
    return []string{"text"}[i]
}
func ParseFormat(v string) (any, error) {
    result := TEXT_FORMAT
    switch v {
        case "text":
            result = TEXT_FORMAT
        default:
            return 0, errors.New("Unknown Format value: " + v)
    }
    return &result, nil
}
func SerializeFormat(values []Format) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Format) isMultiValue() bool {
    return false
}
