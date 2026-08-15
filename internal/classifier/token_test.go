package classifier

import "testing"

func TestTokenize_ShortEnglishTokens(t *testing.T) {
	got := Tokenize("invoice no 123")
	for _, tok := range got {
		if tok == "no" {
			return
		}
	}
	t.Fatalf("Tokenize(%q)=%v, want token %q", "invoice no 123", got, "no")
}
