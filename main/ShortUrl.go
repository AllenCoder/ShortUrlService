package main



import (
	"net/http"
	"io/ioutil"
	"fmt"
	_ "github.com/skip2/go-qrcode"
	"github.com/skip2/go-qrcode"
	"os"
	"encoding/base64"
	"strings"
	"log"
)
func GetShortCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, getShortUrl("http://www.douban.com/note/249723561/")) //这个写入到w的是输出到客户端的
}
func main() {

	http.HandleFunc("/", GetShortCode) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func getShortUrl(parm string) string{
	url := "http://api.t.sina.com.cn/short_url/shorten.json?source=31641035&url_long="+parm
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err == nil {
		//fmt.Print(png)
		ioutil.WriteFile("test.png", png, os.FileMode(0644))
	}
	//imgData,_,_:=image.Decode(bytes.NewBuffer(png))
	dist := make([]byte, 50000)
	base64.StdEncoding.Encode(dist, png)
	fmt.Println(string(dist))
	return "data:image/png;base64,"+string(dist)
}
