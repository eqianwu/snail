package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"os"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
)

func AddSelectProblem(ctx *gin.Context) {
	req := new(vo.AddSelectProblemReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Add select problem bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	org, _ := ctx.Get("user")
	user, err := models.GetToken(org)
	if err != nil {
		log.Printf("SelectProblem controller get token failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.AddSelectProblem(req, user)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func AppendSelectProblem(ctx *gin.Context) {
	req := new(vo.AppendSelectProblemReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Append select problem bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.AppendSelectProblem(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func AddSelectProblemBatch(ctx *gin.Context) {
	req := new(vo.AddSelectProblemBatchReq)
	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Add select problem batch bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	saveFilePath, err := utils.GenFilePath(req.File.Filename)
	if err != nil {
		log.Printf("Add select problem gen file path failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	err = ctx.SaveUploadedFile(req.File, saveFilePath)
	if err != nil {
		log.Printf("Add select problem save path failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	org, _ := ctx.Get("user")
	user, err := models.GetToken(org)
	if err != nil {
		log.Printf("Select problem controller get token failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.AddSelectProblemBatch(req, saveFilePath, user)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func DeleteSelectProblem(ctx *gin.Context) {
	req := new(vo.DeleteSelectProblemReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("delete select problem bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.DeleteSelectProblem(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func DeleteSelectProblemFromSet(ctx *gin.Context) {
	req := new(vo.DeleteSelectProblemFromSetReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Delete select problem from set bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.DeleteSelectProblemFromSet(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func UpdateSelectProblem(ctx *gin.Context) {
	que := new(models.SelectProblem)
	if err := ctx.ShouldBindBodyWith(&que, binding.JSON); err != nil {
		log.Printf("Update Select problem bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.UpdateSelectProblem(que)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func QuerySelectProblemList(ctx *gin.Context) {
	pageRequest := new(helper.PageRequest)
	if err := ctx.ShouldBindBodyWith(&pageRequest, binding.JSON); err != nil {
		log.Printf("Query select problem list bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	org, _ := ctx.Get("user")
	user, err := models.GetToken(org)
	if err != nil {
		log.Printf("Select problem controller get token failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.QuerySelectProblemList(pageRequest, user)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func QuerySelectProblemDetail(ctx *gin.Context) {
	req := new(vo.ProblemDetailReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Query select problem detail bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.QuerySelectProblemDetail(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func QuerySelectProblemCategory(ctx *gin.Context) {
	org, _ := ctx.Get("user")
	user, err := models.GetToken(org)
	if err != nil {
		log.Printf("Select problem controller get token failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.QuerySelectProblemCategory(user)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func FindSelectProblem(ctx *gin.Context) {
	req := new(vo.FindSelectReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Find select problem bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	org, _ := ctx.Get("user")
	user, err := models.GetToken(org)
	if err != nil {
		log.Printf("Select problem controller get token failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.FindSelectProblem(req, user)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func DownloadTemplate(c *gin.Context) {
	workDir, _ := os.Getwd()
	filePath := workDir + "/template/选择题题库模板.xlsx"
	log.Print(filePath)
	//打开文件
	fileTmp, err := os.Open(filePath)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	defer fileTmp.Close()

	//获取文件的名称
	fileName:="选择题题库模板.xlsx"
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	//if common.IsEmpty(filePath) || common.IsEmpty(fileName) || errByOpenFile != nil {
	//	logs.Error("获取文件失败")
	//	c.Redirect(http.StatusFound, "/404")
	//	return
	//}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")

	c.File(filePath)
	return
}