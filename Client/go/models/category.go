package models
import (
    "errors"
)
type Category int

const (
    HELPDOCS_CATEGORY Category = iota
    PUBLICHUBS_CATEGORY
    TRACKDB_CATEGORY
)

func (i Category) String() string {
    return []string{"helpDocs", "publicHubs", "trackDb"}[i]
}
func ParseCategory(v string) (any, error) {
    result := HELPDOCS_CATEGORY
    switch v {
        case "helpDocs":
            result = HELPDOCS_CATEGORY
        case "publicHubs":
            result = PUBLICHUBS_CATEGORY
        case "trackDb":
            result = TRACKDB_CATEGORY
        default:
            return 0, errors.New("Unknown Category value: " + v)
    }
    return &result, nil
}
func SerializeCategory(values []Category) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Category) isMultiValue() bool {
    return false
}
