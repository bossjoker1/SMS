package Model

import "fmt"

type Course struct {
	Cno     string `gorm:"primary_key;type:char(4)"`
	Cname   string `gorm:"type:char(40)"`
	Cpno    string `gorm:"type:char(4)"`
	Ccredit int    `gorm:"type:smallint"`
}

func GetCourse(cno string) (Course, error) {
	var c Course
	err := Db.Where("cno = ?", cno).First(&c).Error
	if err != nil {
		return Course{}, err
	}
	return c, nil
}

func AddCourse(course *Course) error {
	err := Db.Exec("INSERT INTO course(cno, cname, ccredit) VALUES(?, ?, ?)", course.Cno, course.Cname, course.Ccredit).Error
	// err := Db.Create(&course).Error
	if err != nil {
		return err
	}
	fmt.Println("加入成功！")
	return nil
}

func DeleteCourse(cno string) error {
	err := Db.Where("cno = ?", cno).Delete(&Course{}).Error
	if err != nil {
		fmt.Println("删除失败！")
		return err
	}
	fmt.Println("删除成功！")
	return nil
}

func UpdateCourse(cno string, data map[string]interface{}) error {
	err := Db.Model(&Course{}).Where("cno=?", cno).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCourses() []Course {
	var courses []Course
	Db.Find(&courses)
	return courses
}
