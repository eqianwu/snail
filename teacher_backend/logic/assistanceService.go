package logic

import (
	"log"
	"snail/teacher_backend/common"
	"snail/teacher_backend/models"
)

// TODO 助教无权限
func AddAssistance(assistance *models.Assistance) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if isStudent(assistance.StuID) {
		if isAssistanceExist(assistance.StuID, assistance.CourseID) {
			log.Printf("Assistance exist.")
			baseResponse.Code = common.Error
			baseResponse.Msg = "该学生已是课程助教"
			return
		}
		if err := models.CreateAssistance(assistance); err != nil {
			log.Printf("Assistance Server create assistance failed: %v\n", err)
			baseResponse.Code = common.ServerError
		}
	} else {
		baseResponse.Code = common.Error
		baseResponse.Msg = "添加助教失败"
	}
	return
}

func isStudent(stuID string) bool {
	student := new(models.Student)
	student.StudentID = stuID
	err := models.GetSingleStudent(student)
	if err != nil {
		log.Printf("Assistance Service get single student failed: %v\n", err)
		return false
	}
	if student.ID < 1 {
		log.Printf("User Invalid: %v\n", student.StudentID)
		return false
	}
	return true
}

func isAssistanceExist(stuID string, courseID int) bool {
	assistance := new(models.Assistance)
	assistance.StuID = stuID
	assistance.CourseID = courseID
	assistanceList, err := models.GetAssistance(assistance)
	if err != nil {
		log.Printf("Assistance service get assistance failed: %v\n", err)
		return false
	}
	return len(assistanceList) != 0
}

func DeleteAssistance(assistance *models.Assistance) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if err := models.DeleteAssistance(assistance); err != nil {
		log.Printf("Assistance service delete assistance failed: %v\n", err)
		baseResponse.Code = common.Error
		baseResponse.Msg = "删除失败"
	}
	return baseResponse
}
