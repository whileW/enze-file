package qiniu

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/qiniu/api.v7/storage"
	"github.com/whileW/enze-file/model"
	"github.com/whileW/enze-file/service"
	"io"
	"net/http"
)

type QnPuter struct {}
const qn_name  = "qiniu"

func init()  {
	service.RegisterPutFileInter(qn_name,&QnPuter{})
}
//获取上传对象
func get_uploader() (string,*storage.ResumeUploader) {
	qns := GetSetting()

	//获取上传凭证
	mac := GetMac()
	putPolicy := storage.PutPolicy{
		Scope: qns.GetString("bucket"),
	}
	upToken := putPolicy.UploadToken(mac)
	//配置
	cfg := storage.Config{}
	cfg.Zone = GetZone()	// 空间对应的机房
	cfg.UseHTTPS = qns.GetBool("is_use_https") // 是否使用https域名
	cfg.UseCdnDomains = qns.GetBool("is_use_cdn_up") // 上传是否使用CDN上传加速
	resumeUploader := storage.NewResumeUploader(&cfg)
	return upToken,resumeUploader
}
//获取文件完整地址
func get_file_complete_path(key string) string {
	return GetSetting().GetString("img_root_path")+key
}

//上传
func (q *QnPuter) Put(at io.ReaderAt,name string,size int64) (string,string,error) {
	upToken,resumeUploader := get_uploader()
	//上传
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{
	}
	new_file_name := uuid.New().String()
	err := resumeUploader.Put(context.Background(), &ret, upToken,new_file_name, at,size, &putExtra)
	if err != nil {
		return new_file_name,"",err
	}
	return new_file_name,get_file_complete_path(ret.Key),nil
}
func (q *QnPuter) PutFile(local_path string) (string,error) {
	upToken,resumeUploader := get_uploader()
	//上传
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{
	}
	err := resumeUploader.PutFile(context.Background(), &ret, upToken,uuid.New().String(), local_path, &putExtra)
	if err != nil {
		return "",err
	}
	return get_file_complete_path(ret.Key),nil
}
func (q *QnPuter) Get(c *gin.Context,file *model.File)  {
	c.Redirect(http.StatusMovedPermanently, file.Path)
}