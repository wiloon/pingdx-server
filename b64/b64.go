package b64

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	str := "client=wechat"
	//str = "client=corp-wechat"
	fmt.Println("source :", str)
	out := base64.RawURLEncoding.EncodeToString([]byte(str))
	fmt.Println(out)
	bout, _ := base64.RawURLEncoding.DecodeString(out)
	fmt.Println(string(bout))
}

func Decode(str string) string {
	bout, _ := base64.RawURLEncoding.DecodeString(str)
	return string(bout)
}

func DecodeWrapper(c *gin.Context) {
	str := c.Query("str")
	decoded := Decode(str)
	fmt.Printf("decoded: %s", decoded)
	c.JSON(200, gin.H{
		"decoded": decoded,
	})
	return
}
