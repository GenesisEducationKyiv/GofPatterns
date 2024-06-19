package service

type BlacklistChecker struct {
}

func NewBlacklistChecker() *BlacklistChecker {
	return &BlacklistChecker{}
}

func (b *BlacklistChecker) IsInBlackList(email string) bool {
	// Check if email is in popular blacklist like spamhaus.org

	return false
}
