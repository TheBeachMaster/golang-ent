package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"com.thebeachmaster/entexample/common"
)

func GenerateEntityHash(input *common.HashInput) (string, error) {
	now := time.Now().Unix()
	str_now := strconv.FormatInt(now, 10)
	hash := hmac.New(md5.New, []byte(str_now))
	metadataBuilder := strings.Builder{}
	metadataBuilder.WriteString(fmt.Sprintf("%s#%s", input.Metadata, input.Name))
	metadata := metadataBuilder.String()
	_, err := io.WriteString(hash, metadata)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash.Sum(nil)), nil
}
