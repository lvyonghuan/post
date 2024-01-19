package src

import "strings"

// examineWord 审查文档内部是否包含作者个人信息
// TODO：事在人为
func examineWord() {

}

// 检查文件后缀名是是word文档后缀(后缀名为docx或doc)
func checkFileSuffix(fileName string) bool {
	name := strings.Split(fileName, ".")
	if name[len(name)-1] == "docx" || name[len(name)-1] == "doc" {
		return true
	}
	return false
}
