package validator

import (
	"testing"

	"pdf-archive/internal/config"
	"pdf-archive/models"
)

func TestValidate_ThresholdCountsAsSuccess(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.Pipeline.LowConfThreshold = 0.6
	v := New(cfg)
	report := v.Validate(map[string]models.ExtractedField{
		"amount": {Name: "amount", Value: 12.0, Confidence: 0.6},
	}, &config.DocTypeConfig{
		ExtractRules: []config.ExtractRule{{FieldName: "amount", Required: true}},
	})
	if report.SuccessCount != 1 {
		t.Fatalf("confidence==threshold SuccessCount=%d, want 1", report.SuccessCount)
	}
}
