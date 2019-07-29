package utilstring

import (
	"crypto/md5"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/satori/go.uuid"
	mrand "math/rand"
	"wolfkill/wolfkill/common/auth"
	"wolfkill/wolfkill/result/errs"
	"strconv"
	"time"
)

// string to md5string
func Md5String(str string)string{
	m:=md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// create uuid
func SimpleUUID()string{
	return uuid.NewV4().String()
}

func SimpleTokenCreate(str string)(string,error){
	if str==""{
		return "", errs.ErrFunc(errs.Str_NullString)
	}
	rsaStr,err:= RsaEncrypt(str)
	if err != nil{
		return "",err
	}
	return base64.StdEncoding.EncodeToString(rsaStr),nil
}

func SimpleTokenParase(token string)(string,error){
	if token==""{
		return "", errs.ErrFunc(errs.Str_NullString)
	}
	tokenDecode,err:=base64.StdEncoding.DecodeString(token)
	if err != nil{
		return "", errs.ErrFunc(errs.Auth_TokenErr)
	}
	return RsaDecrypt(string(tokenDecode))
}

// encryption
func RsaEncrypt(str string) ([]byte, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode([]byte(auth.PublicKey))
	if block == nil {
		return nil, errs.ErrFunc(errs.Auth_PublicKeyErr)
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	// 加密
	rsaByte,err:=rsa.EncryptPKCS1v15(crand.Reader, pub, []byte(str))
	return rsaByte,err
}

// Decrypt
func RsaDecrypt(CiphertextStr string) (string, error) {
	// 解密
	block, _ := pem.Decode([]byte(auth.PrivateKey))
	if block == nil {
		return "", errs.ErrFunc(errs.Auth_PrivateKeyErr)
	}
	// 解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 解密
	clearStr,err:= rsa.DecryptPKCS1v15(crand.Reader, priv, []byte(CiphertextStr))
	return string(clearStr),err
}

func RandNumber(lenght int) string {
	fmtLenght := "%0"+strconv.Itoa(lenght)+"v"
	minNumberStr:="1"
	for i:=0;i<lenght;i++{
		minNumberStr = minNumberStr+"0"
	}
	minNumber,_:= strconv.Atoi(minNumberStr)
	return fmt.Sprintf(fmtLenght,
		mrand.New(mrand.NewSource(time.Now().UnixNano())).Int31n(int32(minNumber)))
}

