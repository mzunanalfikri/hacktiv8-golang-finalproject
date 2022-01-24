package service

import (
	"project-3/config"
	"project-3/model"
	"time"
)

func GetAllTasks() ([]*model.Task, error) {
	var tasks []*model.Task

	db := config.GetDB()

	err := db.Model(&model.Task{}).Find(&tasks).Error

	return tasks, err
}

func AddTask(task model.Task) (*model.Task, error) {
	db := config.GetDB()

	err := db.Model(&model.Task{}).Create(&task).Error

	return &task, err
}

func UpdateTask(task model.Task) (*model.Task, error) {
	var (
		currentTime = time.Now()
		db          = config.GetDB()
	)

	task.UpdatedAt = &currentTime

	err := db.Model(&model.Task{}).Where("id = ?", task.ID).Updates(&task).Error
	if err != nil {
		return nil, err
	}

	return GetTaskDetail(task.ID)
}

func GetTaskDetail(id int) (*model.Task, error) {
	var (
		task model.Task
		db   = config.GetDB()
	)

	err := db.Model(&model.Task{}).Where("id = ?", id).First(&task).Error

	return &task, err
}

func DeleteTaskById(id int) (interface{}, error) {
	db := config.GetDB()

	err := db.Model(&model.Task{}).Where("id = ?", id).Delete(&model.Task{}).Error

	return map[string]string{"message": "Task has been successfully deleted"}, err
}

func UpdateStatusByTaskId(task model.Task) (*model.Task, error) {

	var (
		currentTime = time.Now()
		db          = config.GetDB()
	)

	task.UpdatedAt = &currentTime
	err := db.Model(model.Task{}).Where("id = ?", task.ID).Updates(&task).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateCategoryByTaskId(task model.Task) (*model.Task, error) {

	var (
		currentTime = time.Now()
		db          = config.GetDB()
	)

	task.UpdatedAt = &currentTime
	err := db.Model(model.Task{}).Where("id = ?", task.ID).Updates(&task).Error
	if err != nil {
		return nil, err
	}

	return &task, err
}
