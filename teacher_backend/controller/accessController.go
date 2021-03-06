package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/models"
	"snail/teacher_backend/vo"
)

// TODO 邮箱校验
// TODO 密码加密
func TeacherRegister(c *gin.Context) {
	teacher := new(models.Teacher)
	if err := c.BindJSON(&teacher); err != nil {
		log.Printf("Teacher register bind json error: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.AddTeacher(teacher)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func TeacherLogin(c *gin.Context) {
	user := new(vo.LoginRequest)
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("Teacher login error: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.TeacherLogin(user)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func ResetPwdReq(c *gin.Context) {
	mail := c.Param("mail")
	baseResponse := logic.ResetPwdReq(mail)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func UpdatePwd(c *gin.Context) {
	newPwd := c.PostForm("pwd")
	proof := c.PostForm("proof")
	mail := c.PostForm("mail")
	baseResponse := logic.UpdatePwd(newPwd, proof, mail)
	c.JSON(http.StatusOK, baseResponse)
	return
}
