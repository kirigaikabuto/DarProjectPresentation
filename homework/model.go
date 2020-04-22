package homework

import "time"

type Repository interface {
	Add(l *HomeWork) (*HomeWork,error)
	Get() ([]*HomeWork,error)
	GetById(id int64) (*HomeWork,error)
	Remove(l *HomeWork) error
	Update(l *HomeWork)  (*HomeWork,error)
}
type HomeWork struct {
	Id int64 `json:"id"`
	TaskDescription string `json:"taskdescription"`
	TaskMaterial string `json:"taskmaterial"`
	EndTime time.Time `json:"endtime"`
	LessonId int64 `json:"lessonid"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updateddat"`
}