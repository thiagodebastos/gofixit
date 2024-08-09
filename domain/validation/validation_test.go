package validation

import "testing"

func TestValidatePassword_Length(t *testing.T) {
	passTooShortErr := ValidatePassword("abc1")
	if passTooShortErr == nil {
		t.Error("expected error, got valid password")
	}

	passTooLongErr := ValidatePassword("abc1XXXXXXXXXXXXXXXX")
	if passTooLongErr == nil {
		t.Error("expected error, got valid password")
	}

	passNoNumberErr := ValidatePassword("xxxxxxxx")
	if passNoNumberErr == nil {
		t.Errorf("expected error, got valid password")
	}
}
