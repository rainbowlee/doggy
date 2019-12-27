package common

import(
	"os"
	"strings"
)

func FileExist(path string) bool {
  _, err := os.Lstat(path)
  return !os.IsNotExist(err)
}


func StingsEndWith(strsrc string, strend string) bool {
	findindex := strings.Index(strsrc, strend)
	if findindex != -1{
		if findindex + len(strend) == len(strsrc){
			return true
		}
	}
	
	return false
}