package main
import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)
var directoryPath string
func main() {
	for {
		fmt.Println("input a directory:")
		fmt.Scanf("%s", &directoryPath)
		GetFilesAndDirs(directoryPath);
	}
}
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			byteArray, error2 := ioutil.ReadFile(dirPth + PthSep + fi.Name())
			if error2 != nil {
				fmt.Println("I/O error for reading " + dirPth + PthSep + fi.Name())
			}
			fmt.Print(md5.Sum(byteArray));
			fmt.Print(" " + dirPth + PthSep + fi.Name());
			fmt.Println();
		}
	}

	return files, dirs, nil
}
