package main

import (
	"SMS/API"
	"SMS/Model"
	"fmt"
	"strconv"
)

// golang的中文对齐就是shit

func PrintStudent(s Model.Student)  {
	fmt.Println("\t学号"+"\t\t"+"姓名"+"\t"+"性别"+"\t"+"年龄"+"\t"+"系别"+"\t"+"奖学金")
	fmt.Println(s.Sno+"\t"+s.Sname+"\t"+s.Ssex+"\t"+strconv.Itoa(s.Sage)+"\t"+s.Sdept+"\t"+s.Scholarship)
}

func PrintCourse(c Model.Course)  {
	fmt.Println("课程号"+"\t"+"课程名"+"\t"+"先修关系"+"\t"+"学分")
	fmt.Println("\t"+c.Cno+"\t"+c.Cname+"\t\t"+c.Cpno+"\t"+strconv.Itoa(c.Ccredit))
}

func PrintSC(sc Model.SC)  {
	fmt.Println("\t学号"+"\t\t"+"课程号"+"\t"+"课程成绩")
	fmt.Println("\t"+sc.Sno+"\t"+sc.Cno+"\t\t"+strconv.Itoa(sc.Grade))
}

func PrintGradeOrder(gs API.GradeStruct)  {
	fmt.Println("\t学号"+"\t\t"+"姓名"+"\t\t"+"课程号"+"\t"+"课程名"+"\t\t"+"成绩")
	fmt.Println(gs.Sno+"\t"+gs.Sname+"\t\t"+gs.Cno+"\t\t"+ gs.Cname +"\t\t"+strconv.Itoa(gs.Grade))
}

func PrintCourseInfo(ci API.CourseChoice)  {
	fmt.Println("\t课程号"+"\t"+"课程名"+"\t\t\t"+"课程成绩")
	fmt.Printf("%6s%10s%15d\n",ci.Cno, ci.Cname, ci.Grade)
}
