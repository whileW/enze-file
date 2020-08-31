package service

import (
	"github.com/jinzhu/gorm"
	"github.com/whileW/enze-file/model"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/utils/cache"
)

var file_db_cache = cache.NewFIFO()

//增加File
func AddFileS(name,new_name,path string,size int64,save_type string) (*model.File,error) {
	return AddFile(name,new_name,path,size,save_type,1,global.GVA_DB.Get(model.FileDb))
}
//增加临时File
func AddFileTemp(name,new_name,path string,size int64,save_type string) (*model.File,error)  {
	return AddFile(name,new_name,path,size,save_type,0,global.GVA_DB.Get(model.FileDb))
}
func AddFile(name,new_name,path string,size int64,save_type string,state int,db *gorm.DB) (*model.File,error) {
	file := &model.File{
		Name:name,
		NewName:new_name,
		FileSize:size,
		Path:path,
		SaveType:save_type,
		State:state,
	}
	if err := db.Create(file).Error;err != nil{
		return nil,err
	}
	file_db_cache.Push(file.Code,file)
	return file,nil
}

//获取文件
func GetFileByCode(code string) (*model.File,error) {
	if f := get_file_by_code_cache(code);f != nil {
		return f,nil
	}
	return get_file_by_code_db(code)
}
func get_file_by_code_db(code string) (*model.File,error) {
	f := &model.File{}
	if err := global.GVA_DB.Get(model.FileDb).Model(f).First(&f,"code = ?",code).Error;err != nil{
		if gorm.IsRecordNotFoundError(err) {
			return nil,nil
		}
		return nil,err
	}
	file_db_cache.Push(f.Code,f)
	return f,nil
}
func get_file_by_code_cache(code string) (*model.File) {
	if d := file_db_cache.Get(code);d != nil{
		return d.(*model.File)
	}
	return nil
}