package designPatterns

import "fmt"

type HashAlgorithm interface {
	Hash(passwordProtector *PasswordProtector)
}

type PasswordProtector struct {
	User         string
	PasswordName string
	Algorithm    HashAlgorithm
}

func (passwordProtector *PasswordProtector) SetAlgorithm(newAlgorithm HashAlgorithm) {
	passwordProtector.Algorithm = newAlgorithm
}

func (passwordProtector *PasswordProtector) Hash() {
	passwordProtector.Algorithm.Hash(passwordProtector)
}

func NewPasswordProtector(user string, passwordName string, algorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		User:         user,
		PasswordName: passwordName,
		Algorithm:    algorithm,
	}
}

type SHA struct{}

func (sha *SHA) Hash(passwordProtector *PasswordProtector) {
	fmt.Printf("using sha for passwordProtector %s\n", passwordProtector.PasswordName)
}

type MD5 struct{}

func (md5 *MD5) Hash(passwordProtector *PasswordProtector) {
	fmt.Printf("using md5 for passwordProtector %s\n", passwordProtector.PasswordName)
}
