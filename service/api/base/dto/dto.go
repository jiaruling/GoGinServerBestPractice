package dto

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/4/11 16:26
*/

type FileContent struct {
	FileName string `json:"filename" binding:"required"`
	Content  string `json:"content" binding:"required"`
}
