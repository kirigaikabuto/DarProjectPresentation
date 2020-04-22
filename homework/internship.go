package homework

type HomeWorkInternship interface {
	AddHomework(hw *HomeWork) (*HomeWork,error)
	GetHomeWork()([]*HomeWork,error)
	CheckDataForEmpty(hw *HomeWork) (error)
}
type homeWorkInternshipClass struct {
	HomeWork Repository
}

func NewHomeWorkInternship(hw Repository) HomeWorkInternship{
	return &homeWorkInternshipClass{HomeWork:hw}
}
func(hwclass *homeWorkInternshipClass) CheckDataForEmpty(hw *HomeWork) error{
	return nil
}
func(hwclass *homeWorkInternshipClass) AddHomework(hw *HomeWork) (*HomeWork,error){
	err := hwclass.CheckDataForEmpty(hw)
	if err!=nil{
		return nil,err
	}
	newhomework,err:=hwclass.HomeWork.Add(hw)
	return newhomework,err
}
func(hwclass *homeWorkInternshipClass) GetHomeWork()([]*HomeWork,error){
	homeworks,err:=hwclass.GetHomeWork()
	if err!=nil{
		return nil,err
	}
	return homeworks,nil
}

