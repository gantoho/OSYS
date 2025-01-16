// 加盐

package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func Encrypt(pwd string) string {
	nawPwd := pwd + "ggbond"
	hash := md5.New()
	hash.Write([]byte(nawPwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
