package idgen

import (
	"fmt"
	"strings"

	"github.com/onkyou/idgen/pkg/idprefix"

	"github.com/segmentio/ksuid"
)

// NewID generates a new ID using the provided prefix and a KSUID.
func NewID(prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, ksuid.New().String())
}

// ExtractPrefix returns the prefix from a prefixed ID like "usr_xxx".
func ExtractPrefix(id string) (string, error) {
	parts := strings.SplitN(id, "_", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid ID format: %s", id)
	}
	return parts[0], nil
}

// LookupEntityType maps the prefix (e.g., "usr") to its entity name (e.g., "User").
func LookupEntityType(id string) (string, error) {
	prefix, err := ExtractPrefix(id)
	if err != nil {
		return "", err
	}
	entity, ok := idprefix.PrefixToEntity[prefix]
	if !ok {
		return "", fmt.Errorf("unknown prefix: %s", prefix)
	}
	return entity, nil
}
