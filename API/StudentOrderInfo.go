package API

import (
	"SMS/Model"
	"fmt"
)

type GradeStruct struct {
	Sno string
	Sname string
	Cno string
	Cname string
	Grade int
}

type CourseChoice struct{
	Cno string
	Cname string
	Grade int
}

// 按系对该系的学生进行排名(按成绩从高到底)
func GradeOrder(sdept string) (gs []GradeStruct){
	Model.Db.Raw("SELECT student.sno,student.sname ,sc.cno, course.cname, sc.grade" +
		" FROM sc,student,course WHERE student.sdept=? " +
		"AND student.sno=sc.sno AND course.cno=sc.cno ORDER BY grade DESC", sdept).Scan(&gs)
	return
}

func GetStudentInfo(sno string) (stuInfo Model.Student, courseInfo []CourseChoice) {
	stuInfo, err := Model.GetOneStudent(sno)
	if err != nil{
		fmt.Println(err)
	}
	Model.Db.Raw("select course.cno, course.cname, sc.grade" +
		" from course, sc where course.cno=sc.cno and sc.sno=?", sno).Scan(&courseInfo)
	return
}
