package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type File struct{
	gorm.Model
	//UUID        		string    		`json:"uuid" gorm:"type:varchar(128)"`		//文件id
	//Version				int				`json:"version" gorm:"type:int"`			//文件版本
	Code 				string			`json:"code" gorm:"type:varchar(128)"`		//code
	Name				string			`json:"name" gorm:"type:varchar(128)"`		//文件原名称
	NewName				string			`json:"new_name" gorm:"type:varchar(128)"`	//新文件名称
	Path 				string			`json:"path" gorm:"type:varchar(128)"`		//文件地址
	FileSize 			int64			`json:"file_size" gorm:"type:int"`			//文件大小
	//FileType 			int				`json:"file_type" gorm:"type:int"`			//文件类型
	//保存方式 local-本地 qiniu-七牛
	SaveType 			string			`json:"save_type" gorm:"type:varchar(128)"`
	//状态 -2已删除 -1 等待删除 0临时文件 1正常  2低频存储  3归档存储
	State 				int				`json:"state" gorm:"type:int"`
}

func (f *File) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("code", uuid.New().String())
	return nil
}