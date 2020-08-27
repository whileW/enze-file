package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-file/model"
	"github.com/whileW/enze-file/service/qiniu"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/utils/resp"
	"io"
	"net/http"
)

type PutFile interface {
	GetName()string
	Put(at io.ReaderAt,name string,size int64) (string,string,error)
	Get(c *gin.Context,file *model.File)
} 
var put_file_inter = map[string]PutFile{}
func init()  {
	qn := qiniu.QnPuter{}
	put_file_inter[qn.GetName()] = &qn
}

func Put(file io.ReaderAt,name string,size int64) (string,error) {
	upload_type := global.GVA_CONFIG.Setting.GetString("upload_type")
	put_file_m := put_file_inter[upload_type]
	new_name,path,err := put_file_m.Put(file,name,size)
	if err != nil {
		return "",errors.New(fmt.Sprintf("上传失败：%v",err))
	}
	filem,err := AddFileS(name,new_name,path,size,put_file_m.GetName())
	if err != nil {
		return "",errors.New(fmt.Sprintf("入库失败：%v",err))
	}
	return filem.Code,nil
}
func Get(c *gin.Context,code string)  {
	file,err := GetFileByCode(code)
	if err != nil {
		global.GVA_LOG.Errorw("根据code获取文件","err",err)
		resp.FailWithMessage(c,"系统异常请稍后重试")
		return
	}
	if file == nil {
		c.Status(http.StatusNotFound)
		c.Abort()
		return
	}
	put_file_inter[file.SaveType].Get(c,file)
}