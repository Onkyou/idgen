package idgen_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/onkyou/idgen/pkg/idgen"
	"github.com/onkyou/idgen/pkg/idprefix"
)

func ExampleNewID() {
	id := idgen.NewID(idprefix.User)
	fmt.Println("Prefix:", id[:3])
	// Output:
	// Prefix: usr
}

func TestNewID_GeneratesWithPrefix(t *testing.T) {
	id := idgen.NewID(idprefix.User)
	if !strings.HasPrefix(id, idprefix.User+"_") {
		t.Errorf("expected prefix %q in id %q", idprefix.User+"_", id)
	}
}

func TestLookupEntityType_KnownPrefix(t *testing.T) {
	id := idgen.NewID(idprefix.Achievement)
	entity, err := idgen.LookupEntityType(id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if entity != "Achievement" {
		t.Errorf("expected 'Achievement', got %q", entity)
	}
}

func TestLookupEntityType_UnknownPrefix(t *testing.T) {
	id := "xxx_abc123"
	_, err := idgen.LookupEntityType(id)
	if err == nil {
		t.Fatal("expected error for unknown prefix, got nil")
	}
}

func TestExtractPrefix(t *testing.T) {
	id := "evt_abc123"
	prefix, err := idgen.ExtractPrefix(id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if prefix != "evt" {
		t.Errorf("expected prefix 'evt', got %q", prefix)
	}
}

func TestExtractPrefix_InvalidFormat(t *testing.T) {
	id := "missingdelimiter"
	_, err := idgen.ExtractPrefix(id)
	if err == nil {
		t.Fatal("expected error for invalid ID format")
	}
}
