package src

import (
	"errors"
	"mime/multipart"
)

// StoreFile 检查文件合法性并存储文件
func StoreFile(file *multipart.FileHeader, title, track string) error {
	fileName := file.Filename

	//检查文件后缀名是是word文档后缀(后缀名为docx或doc)
	if !checkFileSuffix(fileName) {
		return errors.New("文件格式错误")
	}
}
