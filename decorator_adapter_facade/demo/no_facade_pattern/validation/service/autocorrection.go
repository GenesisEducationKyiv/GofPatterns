package validation

type AutoCorrection struct {
}

func NewAutoCorrection() *AutoCorrection {
	return &AutoCorrection{}
}

func (a *AutoCorrection) CorrectEmail(email string) string {
	// Correct email process like replace gmail.con with gmail.com

	return email
}
