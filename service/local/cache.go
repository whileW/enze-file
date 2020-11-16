package local

import (
	"io/ioutil"
)
// 编写中
var local_file_cache = map[string]int{}

func init()  {
	get_dir_child_file("upload")
}

func get_dir_child_file(path string) {
	f,_ := ioutil.ReadDir(path)
	for _,t := range f {
		if t.IsDir() {
			get_dir_child_file(path+"/"+t.Name())
		}else {
			local_file_cache[path+"/"+t.Name()] = 0
		}
	}
}
func is_have_file(path string) bool {
	if _,ok := local_file_cache[path]; ok {
		return ok
	}
	return false
}