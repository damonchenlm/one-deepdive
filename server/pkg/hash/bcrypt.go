package hash

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	Cost int
}

//Make 加密方法
func (b *Bcrypt) Make(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, b.Cost)
}

//Check 检查方法
func (b *Bcrypt) Check(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
