package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-file/service"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/utils"
	"github.com/whileW/enze-global/utils/resp"
	"strings"
)

//上传
func Upload(c *gin.Context)  {
	mfile, err := c.FormFile("file")
	if err != nil {
		global.GVA_LOG.Errorw("获取文件异常","err",err)
		resp.FailWithDetailed(c,resp.ParamterError,"没有获取到文件",nil)
		return
	}
	filei,err := mfile.Open()
	if err != nil {
		global.GVA_LOG.Errorw("获取文件异常","err",err)
		resp.FailWithDetailed(c,resp.ParamterError,"没有获取到文件",nil)
		return
	}
	//文件后缀
	file_suffix := ""
	if f := strings.Split(mfile.Filename,"."); len(f)>=2 {
		file_suffix = strings.Split(mfile.Filename,".")[1]
	}
	file_allow_suffix := strings.Split(global.GVA_CONFIG.Setting.GetString("file_allow_suffix"),";")
	if suffix_allow := utils.SliceStringContains(file_allow_suffix,file_suffix); !suffix_allow {
		resp.FailWithDetailed(c,resp.ParamterError,"不允许的文件后缀",nil)
		return
	}
	//上传
	Code,err := service.Put(filei,mfile.Filename,mfile.Size)
	if err != nil {
		global.GVA_LOG.Errorw("上传异常","err",err)
		resp.FailWithMessage(c,"上传异常，请稍后重试")
		return
	}
	resp.OkWithData(c,Code)
}

func Get(c *gin.Context)  {
	code := c.Param("code")
	if code == "" {
		resp.NoFindResult(c)
		return
	}
	service.Get(c,code)
}