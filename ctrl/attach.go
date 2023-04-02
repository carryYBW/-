package ctrl

import (
	"helloweb/util"
	"net/http"
	"os"
)

func init() {
	os.MkdirAll("./mnt", os.ModePerm)
}

func UpLoadLocal(w http.ResponseWriter, r http.Request) {
	//获取上传的文件
	file, header, err := r.FormFile("file")
	if err != nil {
		util.RespFail(w, err.Error())
	}

}
