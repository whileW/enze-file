package local

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/whileW/enze-file/model"
	"github.com/whileW/enze-file/service"
	"github.com/whileW/enze-global/utils"
	"github.com/whileW/enze-global/utils/resp"
	"io"
	"os"
	"strings"
	"time"
)

type LocalPuter struct {}
const local_name  = "local"

func init()  {
	service.RegisterPutFileInter(local_name,&LocalPuter{})
}

func (l *LocalPuter)Put(at io.ReaderAt,name string,size int64) (string,string,error){
	file_buf := make([]byte,size)
	if _,err := at.ReadAt(file_buf,0);err != nil && err != io.EOF{
		return "","",errors.New(fmt.Sprintf("read file err:%v",err))
	}

	path := "upload/"+time.Now().Format("20060102")
	file_name := uuid.New().String()+"."+strings.Split(name,".")[1]
	//创建文件夹
	if err := utils.CreateDir(path); err != nil {
		return "","",errors.New(fmt.Sprintf("create dir err:%v",err))
	}
	//创建文件
	file, err := os.Create(path+"/"+file_name)
	defer file.Close()
	if err != nil{
		return "","",errors.New(fmt.Sprintf("create file err:%v",err))
	}
	//写入文件
	if _,err := file.Write(file_buf);err != nil{
		return "","",errors.New(fmt.Sprintf("write file err:%v",err))
	}
	return file_name,path+"/"+file_name,nil
}
func (l *LocalPuter)Get(c *gin.Context,file *model.File){
	f,err := os.Open(file.Path)
	defer f.Close()
	if err != nil {
		resp.NoFindResult(c)
		return
	}
	c.File(file.Path)
	return
}