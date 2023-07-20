package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorGetTagsFail    = NewError(20010002, "获取多个标签失败")
	ErrorGetTagFail     = NewError(20010003, "获取单个标签失败")
	ErrorCreateTagFail  = NewError(20010004, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010005, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010006, "删除标签失败")
	ErrorCountTagFail   = NewError(20010007, "统计标签失败")

	ErrorGetArticleListFail = NewError(20020001, "获取文章列表失败")
	ErrorGetArticlesFail    = NewError(20020002, "获取多个文章失败")
	ErrorGetArticleFail     = NewError(20020003, "获取单个文章失败")
	ErrorCreateArticleFail  = NewError(20020004, "创建文章失败")
	ErrorUpdateArticleFail  = NewError(20020005, "更新文章失败")
	ErrorDeleteArticleFail  = NewError(20020006, "删除文章失败")
	ErrorCountArticleFail   = NewError(20020007, "统计文章失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)
