package models

import (
	"snail/teacher_backend/dao"
	"time"
)

type Course struct {
	ID          int       `json:"id"`
	CourseTitle string    `json:"course_title"`
	IsPrivate   int       `json:"is_private"`
	Intro       string    `json:"intro"`
	Goal        string    `json:"goal"`
	UserFor     string    `json:"user_for"`
	SearchCode  string    `json:"search_code"`
	CreateBy    string    `json:"create_by"`
	CreateTime  time.Time `json:"create_time"`
}

func CreateCourse(course *Course) (err error) {
	err = dao.DB.Create(&course).Error
	return
}

func UpdateCourse(course *Course) (err error) {
	err = dao.DB.Model(&Course{}).Where("id = ?", course.ID).Updates(&course).Error
	return
}

func GetCourse(course *Course) (courseList []Course, total int, err error) {
	if err = dao.DB.Where(&course).Find(&courseList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return
	// if err = dao.DB.Where(&course).Limit(pageSize).Offset((page - 1) * pageSize).Find(&courseList).Count(&total).Error; err != nil {
	// 	return nil, 0, err
	// }
	// return
}

func GetSingleCourse(course *Course) (err error) {
	err = dao.DB.Where(&course).First(&course).Error
	return
}

func GetCourseByID(idList []int) (courseList []Course, total int, err error) {
	if err := dao.DB.Where("id in (?)", idList).Find(&courseList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return
}

func DeleteCourse(course *Course) (err error) {
	err = dao.DB.Delete(&course).Error
	return
}
