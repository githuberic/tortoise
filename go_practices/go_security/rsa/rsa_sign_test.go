package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"testing"
)

func Sign(content, prvKey []byte) (sign string, err error) {
	block, _ := pem.Decode(prvKey)
	if block == nil {
		fmt.Println("pem.Decode err")
		return
	}
	var private interface{}
	private, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	privateKey := private.(*rsa.PrivateKey)
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(content))
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey,
		crypto.SHA1, hashed)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(signature)
	return
}

func RSAVerify(origdata, ciphertext string, publicKey []byte) (bool, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return false, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(origdata))
	digest := h.Sum(nil)
	body, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, digest, body)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
Not verfied
*/
func TestSign(t *testing.T) {
	pubKey := "-----BEGIN PUBLIC KEY-----\n{MIIDuTCCAqGgAwIBAgIQDvoyveXz+P2SFaDiRUEkWzANBgkqhkiG9w0BAQsFADBnMQswCQYDVQQGEwJjbjERMA8GA1UECBMIU2hhbmdIYWkxETAPBgNVBAcTCFNoYW5nSGFpMQ4wDAYDVQQKEwVzdW5taTEMMAoGA1UECxMDSW9UMRQwEgYDVQQDDAsqLnN1bm1pLmNvbTAeFw0yMjA0MTkwNjMzMDRaFw0zMjA0MTkwNjMzMDRaMGcxCzAJBgNVBAYTAmNuMREwDwYDVQQIEwhTaGFuZ0hhaTERMA8GA1UEBxMIU2hhbmdIYWkxDjAMBgNVBAoTBXN1bm1pMQwwCgYDVQQLEwNJb1QxFDASBgNVBAMMCyouc3VubWkuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqejr5SCJWh54kgeq+7WuSuoosNv4lHAMLXFAxDBvBwc5BPtA4R+Jt1Iet4FFBkOSENsbehyEIKgekrEt3GSPDQlc+grSyufQoxWT0TgAp7AezuNwF3ooFZv9z1E6LMjESg9jGhMotQu2YaNB3mTUt7VZBLqmTCFaTX67Ex3a6VsNxiPNAMwJf9KT9vTaQS5p98z/XuDXEYb5rYrWLUMpHx536NBZQ7HI8EHFlH8eTFE8g+7YH8NW4tvWZq14wrQS8ksLk0BhQVkA6fGjbKkqPmN+oRdBPp4qqu1AsV9lhGSrNdRWGpCDfBaiNap5tLWkUOMrTAbtyI5vdkTvSW0KfwIDAQABo2EwXzAOBgNVHQ8BAf8EBAMCAoQwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFAtOBs9VE2O1h3dMl73nb6vnq9WiMA0GCSqGSIb3DQEBCwUAA4IBAQCUuGkatR2hoGMzbXpQiyZwLM3ddXNYYweXT8EhUW8M12g+d9dXUoC/fOVN4P85kAR4dU2RNWSW0z4Lbfz+4vbQkAMsAhRnTfGKH2BlQVYq2oIT6pzxMItXsxROOoWMcJpXbwLA05lLWvvD2pNcwHSJ2V909fJSlhmWL8TyeC4yVXhtTGlOPg/44MYue53fE5Uy/Ft8CqzmRQgZzkbAwRDcYG73WrfjDRPzuBsjNt9i3FKjPLMOsvwuQ3JTCoW44eb/T22O7wAnEBDiEJ42Z8x2Ag2/fEQUoeaTwpM6E8lis0f9Vzz5bELCoLQ/ArjYmMJhNgxOlZbp/i97L3r0xn/d}\n-----END PUBLIC KEY-----"
	prvKey := "-----BEGIN RSA PRIVATE KEY-----\n{MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCp6OvlIIlaHniSB6r7ta5K6iiw2/iUcAwtcUDEMG8HBzkE+0DhH4m3Uh63gUUGQ5IQ2xt6HIQgqB6SsS3cZI8NCVz6CtLK59CjFZPROACnsB7O43AXeigVm/3PUTosyMRKD2MaEyi1C7Zho0HeZNS3tVkEuqZMIVpNfrsTHdrpWw3GI80AzAl/0pP29NpBLmn3zP9e4NcRhvmtitYtQykfHnfo0FlDscjwQcWUfx5MUTyD7tgfw1bi29ZmrXjCtBLySwuTQGFBWQDp8aNsqSo+Y36hF0E+niqq7UCxX2WEZKs11FYakIN8FqI1qnm0taRQ4ytMBu3Ijm92RO9JbQp/AgMBAAECggEABijzpHLxh9ri9ChjAo0rrHugzgirknYdeuhIONSosdCJOK0DRSYnkAXiJzUAk6GPghzvehXelHCD6z2RNBRRO1/Dhy0tRz3wSnlVypgFLg8sjsOgJLnBCkeFhCalvUnqE+VK6n8XUrkB+7g3gyCUtbRVbd52ePwckgKuNZs3fI5JXcf1AEpqwzTu4fl0WHgS7a+WyB1bFxIbX56gOF4TfZOc+SdJtLpZ8tuio32lweqsVWfk820t0wH1w7n7DjqUZ//+LYqkihvY+BCiVkRujTFHi25uoNj6v8Xow441l8jwNTUe6D19m2OLiaxp++BepdrwLhf6Z4+ljzUCGvJBAQKBgQDh7YTU05ibV7gCXVEKRsvwdhM5D/1hqAq0fnZs2KyR4Yq2utbatXm+npqb0BWRvHiHT4M66XsOh4YaQXE0m3MW0IydVLFsExhmCWrkfPQXp3kJqKA72qTMI+uyCC6mcWO33LBybr25TOsOrQgoIsXJEzUGNYGRKHncSqSRYPj+8QKBgQDAhpfGKBeE0s2hO6aUS8tETplmqc9J1mS78ocRhQqJ6u0fexdayGUBI14quS0T73oU335Sue6BOb3ysp+Fo6E9XL6ggNcjNtSHfhdLjZPMmKhwJzp2YkreA+sWsQ8BqSRKmu+6lPBUAIAp0C5sQmI5YpLGKXeEUyVDId5JOd6AbwKBgFf8e3WzLkSZReVN2EjLR5NPOK9JZeGJYxnmAlLjk1mA39fILfKhKFevRwEVacEHCF2R/saLtUv0Raql/+OAArNI5upC2JIBYcgjNMjlTN6Fb2Luk/CbdwC1oI4GF7zkffRV8zWtoR+wL/7YLkwLNsHXRBg18E7K16KOQ0cG4BHRAoGBAKKlkjcfy660gG/D+xixZ4nPdR0cPDi+3N9DJwSQBrU5ORYtYS+auGnbvUQOqJx85gAYR5oP5gRaCbStXsMB18gFQC5ManQfcDD/PyLe+owQM8x795HnvCQP312OO4VFpZk5h2lzRg1RVvT8IRvogR7FvgWXJpctvv8V77qhHu4VAoGBAJbcUdyhfP7tZj2cjpDOGD9f/AauwD0u3eJQFo8m9X+lOZEqLSTZWhxhEwIxNdRr//b2JyFN9VJbYhNDG07PK7E5GcWO5beT+BBghxAQoZ2IANR8znNaHcMjeQf62x85meuAbSzLp9rKwi9ArWEXWn5begJqdCC0fz6wypZgNBra}\n-----END RSA PRIVATE KEY-----"

	content := "模型训练需要花多长时间"

	sign, err := Sign([]byte(content), []byte(prvKey))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("sign签名结果：", sign)

	res, err := RSAVerify(content, sign, []byte(pubKey))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("验签结果：", res)
}
