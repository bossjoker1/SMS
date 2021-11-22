package Model

import "fmt"

type Student struct {
	Sno string `gorm:"primary_key;type:char(9)"json:"sno"`
	Sname string `gorm:"type:char(20)"json:"sname"`
	Ssex string	`gorm:"type:char(2)"json:"ssex"`
	Sage int	`gorm:"type:smallint"json:"sage"`
	Sdept string	`gorm:"type:char(20)"json:"sdept"`
	Scholarship string	`gorm:"type:char(2)"json:"scholarship"`
}

func AddStudent(data *Student) error {
	err := Db.Create(&data).Error
	if err != nil{
		return err
	}
	fmt.Println("加入成功！")
	return nil
}

func DeleteStudent(sno string) error{
	err := Db.Where("sno = ?", sno).Delete(&Student{}).Error
	if err != nil{
		return err
	}
	fmt.Println("删除成功！")
	return nil
}

func UpdateStudent(sno string, data map[string]interface{}) error{
	err := Db.Model(&Student{}).Where("sno = ?", sno).Updates(&data).Error
	if err != nil{
		return err
	}
	return nil
}

func GetOneStudent(sno string) (Student, error) {
	var s Student
	err := Db.Where("sno = ?", sno).First(&s).Error
	if err != nil {
		return Student{}, err
	}
	return s, nil
}

func GetStudents() []Student {
	var students []Student
	Db.Find(&students)
	return students
}