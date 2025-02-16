package plugins

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	common "github.com/collect-ui/collect/src/collect/common"
	config "github.com/collect-ui/collect/src/collect/config"
	templateService "github.com/collect-ui/collect/src/collect/service_imp"
	utils "github.com/collect-ui/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"strings"
)

type HandlerPassword struct {
	templateService.BaseHandler
}

func (si *HandlerPassword) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	params := template.GetParams()
	field := handlerParam.Field
	op := handlerParam.Operation

	aseKey := utils.GetAppKey("ase_key")
	aseIv := utils.GetAppKey("ase_iv")
	aseTag := utils.GetAppKey("ase_tag")
	ase_way := utils.GetAppKey("ase_way")
	company_key := utils.GetAppKey("company_key")

	if !utils.IsValueEmpty(handlerParam.Foreach) { // 如果是数组

		list, errMsg := utils.RenderVarToArrMap(handlerParam.Foreach, params)
		if !utils.IsValueEmpty(errMsg) {
			return common.NotOk(errMsg)
		}
		fieldName := utils.GetRenderVarName(field)
		for i, item := range list {
			target := utils.RenderVar(field, item)
			if target == nil {
				continue
			}
			value := target.(string)
			if op == "encrypt" { // 处理加密
				if strings.HasSuffix(value, aseTag) && strings.HasPrefix(value, aseTag) {
					continue
				}
				result, ok := encrypt([]byte(aseKey), []byte(aseIv), []byte(value))
				tmp := hex.EncodeToString([]byte(result))
				if ok == nil {
					list[i][fieldName] = aseTag + tmp + aseTag
				}
			} else {
				if !(strings.HasSuffix(value, aseTag) && strings.HasPrefix(value, aseTag)) {
					continue
				}
				value = strings.ReplaceAll(value, aseTag, "")
				// 公司解密
				if ase_way == "company" {
					decryptedText := AESDecryptECBStr(value, company_key)
					list[i][fieldName] = gocast.ToString(decryptedText)
				} else { // 个人项目解密
					// 去掉加密标签
					//hex 处理回来
					b, _ := hex.DecodeString(value)
					bs := gocast.ToString(b)
					// 解密
					decryptedText, err := decrypt([]byte(aseKey), []byte(aseIv), bs)
					if err == nil {
						list[i][fieldName] = gocast.ToString(decryptedText)
					}
				}

			}
		}
		r := common.Ok(list, "处理参数成功")
		return r

	} else {
		value := utils.RenderVar(field, params).(string)
		var r *common.Result
		if op == "encrypt" {
			if strings.HasSuffix(value, aseTag) && strings.HasPrefix(value, aseTag) {
				r = common.Ok(value, "无需处理")
				return r
			} else {
				// 加密处理
				result, ok := encrypt([]byte(aseKey), []byte(aseIv), []byte(value))
				//hex 处理一下
				tmp := hex.EncodeToString([]byte(result))
				if ok == nil { // 加密成功
					// 加上加密标签
					value = aseTag + tmp + aseTag
					r = common.Ok(value, "处理参数成功")
				} else { // 加密失败
					r = common.NotOk("加密失败")
				}
			}

		} else {
			// 判断是否与加密标签，无加密标签直接返回
			if !(strings.HasSuffix(value, aseTag) && strings.HasPrefix(value, aseTag)) {
				r = common.Ok(value, "无需处理")
				return r
			}
			value = strings.ReplaceAll(value, aseTag, "")
			// 去掉加密标签
			//hex 处理回来
			b, _ := hex.DecodeString(value)
			bs := gocast.ToString(b)
			// 解密
			decryptedText, err := decrypt([]byte(aseKey), []byte(aseIv), bs)
			if err == nil {
				r = common.Ok(gocast.ToString(decryptedText), "解密成功")
			} else {
				r = common.NotOk("解密失败")
			}

		}
		return r
	}

}

func encrypt(key, iv, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	paddedPlaintext := pkcs7Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
func decrypt(key, iv []byte, ciphertext string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	decryptedData := make([]byte, len(decodedCiphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decryptedData, decodedCiphertext)
	return pkcs7Unpadding(decryptedData), nil
}

func generateKeys(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
func AESDecryptECBStr(encrypteds string, keys string) string {
	// 字符串转换成切片
	//encrypted := []byte(encrypteds)
	encrypted, _ := hex.DecodeString(encrypteds)
	key := []byte(keys)

	cipher, _ := aes.NewCipher(generateKeys(key))
	decrypted := make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return string(decrypted[:trim])
}

// 使用PKCS7填充方式对数据进行填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 对使用PKCS7填充方式的数据进行去填充
func pkcs7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
