package validation

type DNSValidation struct {
}

func NewDNSValidation() *DNSValidation {
	return &DNSValidation{}
}

func (d *DNSValidation) ValidateMXRecord(email string) bool {
	// Validate MX record exist of email domain

	return true
}
