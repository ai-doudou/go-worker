/**
 * @Author : jinchunguang
 * @Date : 19-11-7 下午5:49
 * @Project : go-worker
 */
package utils

import (
    "crypto/md5"
    "encoding/hex"
)

// generate md5 checksum of URL in hex format
func GenerateMd5(str string) string {
    m := md5.New()
    m.Write([]byte(str))
    c := m.Sum(nil)
    return hex.EncodeToString(c)
}
