package main

import (
	"SMS/API"
	"SMS/Model"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UpdateFields() map[string]interface{} {
	var (
		mp   = make(map[string]interface{})
		temp = "default"
	)
	fmt.Println("请输入你想修改的字段即相应的值(end退出)：")
	for {
		reader := bufio.NewReader(os.Stdin)
		res, _, _ := reader.ReadLine()
		temp = string(res)
		array := strings.Fields(temp)
		if temp == "end" || len(array) == 1 {
			break
		}
		mp[array[0]] = array[1]
	}
	return mp
}

func main() {

	var op = 1
	fmt.Printf("------欢迎使用学生信息管理系统------\n\n")
	fmt.Printf("$学生基本信息操作：\n")
	fmt.Printf("1.增加学生信息----2.修改学生信息----3.删除学生信息----4.打印指定学生信息----5.打印所有学生信息\n\n")
	fmt.Printf("$课程信息操作\n")
	fmt.Printf("6.增加课程信息----7.修改课程信息----8.删除课程信息----9.打印指定课程信息----10.打印所有课程信息\n\n")
	fmt.Printf("选修表信息操作\n")
	fmt.Printf("11.增加选修信息----12.修改选修信息----13.删除选修信息----14.打印指定选修信息----15.打印所有选修信息\n\n")
	fmt.Printf("$高级查询：\n")
	fmt.Printf("按系查询：\n")
	fmt.Printf("16.平均成绩----17.最好和最差成绩----18.优秀率----19----不及格人数----20.成绩排名信息\n")
	fmt.Printf("21.学生基本信息和选课信息----0.退出系统\n")

	for {
		fmt.Printf("\n请输入你的选择：")
		_, _ = fmt.Scanf("%d", &op)
		if 0 == op {
			break
		}
		switch op {
		case 1:
			var student Model.Student
			fmt.Println("请输入要添加的学生信息：")
			fmt.Printf("学生学号: ")
			_, _ = fmt.Scan(&student.Sno)
			fmt.Printf("学生姓名：")
			_, _ = fmt.Scan(&student.Sname)
			fmt.Printf("学生性别：")
			_, _ = fmt.Scan(&student.Ssex)
			fmt.Printf("学生年龄：")
			_, _ = fmt.Scan(&student.Sage)
			fmt.Printf("学生系别：")
			_, _ = fmt.Scan(&student.Sdept)
			fmt.Printf("是否获取奖学金：")
			_, _ = fmt.Scan(&student.Scholarship)

			err := Model.AddStudent(&student)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			var sno string
			fmt.Println("请输入需要修改的学生的学号")
			_, _ = fmt.Scan(&sno)
			mp := UpdateFields()
			err := Model.UpdateStudent(sno, mp)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("更新后的字段：")
				student, _ := Model.GetOneStudent(sno)
				PrintStudent(student)
			}
		case 3:
			var sno string
			fmt.Println("请输入想删除学生的学号：")
			_, _ = fmt.Scan(&sno)
			err := Model.DeleteStudent(sno)
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			var sno string
			fmt.Println("请输入要打印学生的学号：")
			_, _ = fmt.Scan(&sno)
			student, err := Model.GetOneStudent(sno)
			if err != nil {
				fmt.Println(err)
			} else {
				PrintStudent(student)
			}
		case 5:
			fmt.Println("打印所有学生信息如下：")
			students := Model.GetStudents()
			for _, s := range students {
				PrintStudent(s)
			}
		case 6:
			var course Model.Course
			fmt.Println("请输入要添加的课程信息：")
			fmt.Printf("课程号: ")
			_, _ = fmt.Scan(&course.Cno)
			fmt.Printf("课程名：")
			_, _ = fmt.Scan(&course.Cname)
			fmt.Printf("先修关系：")
			_, _ = fmt.Scan(&course.Cpno)
			fmt.Printf("学分：")
			_, _ = fmt.Scan(&course.Ccredit)
			fmt.Printf("学生系别：")

			err := Model.AddCourse(&course)
			if err != nil {
				fmt.Println(err)
			}
		case 7:
			var cno string
			fmt.Println("请输入想要修改课程的课程号：")
			_, _ = fmt.Scan(&cno)
			mp := UpdateFields()
			err := Model.UpdateCourse(cno, mp)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("更新后的字段：")
				course, _ := Model.GetCourse(cno)
				PrintCourse(course)
			}
		case 8:
			var cno string
			fmt.Println("请输入想删除学生的学号：")
			_, _ = fmt.Scan(&cno)
			err := Model.DeleteCourse(cno)
			if err != nil {
				fmt.Println(err)
			}
		case 9:
			var cno string
			fmt.Println("请输入要打印课程的课程号：")
			_, _ = fmt.Scan(&cno)
			course, err := Model.GetCourse(cno)
			if err != nil {
				fmt.Println(err)
			} else {
				PrintCourse(course)
			}
		case 10:
			fmt.Println("打印所有课程信息如下：")
			courses := Model.GetCourses()
			for _, c := range courses {
				PrintCourse(c)
			}
		case 11:
			var sc Model.SC
			fmt.Println("请输入要添加的选修信息：")
			fmt.Printf("学生学号: ")
			_, _ = fmt.Scan(&sc.Sno)
			fmt.Printf("课程号：")
			_, _ = fmt.Scan(&sc.Cno)
			fmt.Printf("课程成绩：")
			_, _ = fmt.Scan(&sc.Grade)
			err := Model.AddSc(&sc)
			if err != nil {
				fmt.Println(err)
			}
		case 12:
			var sno, cno string
			fmt.Println("请输入需要修改选修信息的学号和课程号")
			_, _ = fmt.Scan(&sno, &cno)
			mp := UpdateFields()
			err := Model.UpdateSC(cno, sno, mp)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("更新后的字段：")
				sc, _ := Model.GetOneSC(cno, sno)
				PrintSC(sc)
			}
		case 13:
			var sno, cno string
			fmt.Println("请输入想删除选修信息的学号和课程号：")
			_, _ = fmt.Scan(&sno, &cno)
			err := Model.DeleteOneSC(cno, sno)
			if err != nil {
				fmt.Println(err)
			}
		case 14:
			var sno, cno string
			//fmt.Println("请选择以下打印方式")
			//fmt.Println("1.输出指定学号的所有选修信息")
			//fmt.Println("2.输出指定课程号的所有选修信息")
			//fmt.Println("3.输出指定选秀吧信息")
			fmt.Println("请输入要打印选修信息的学号和课程号：")
			_, _ = fmt.Scan(&sno, &cno)
			sc, err := Model.GetOneSC(cno, sno)
			if err != nil {
				fmt.Println(err)
			} else {
				PrintSC(sc)
			}
		case 15:
			fmt.Println("打印所有选修信息如下：")
			scs := Model.GetSC()
			for _, sc := range scs {
				PrintSC(sc)
			}
		case 16:
			var sdept string
			fmt.Println("请输入要查询的系别：")
			_, _ = fmt.Scan(&sdept)
			fmt.Println(sdept+"系的平均成绩为：", API.GetAvg(sdept))

		case 17:
			var sdept string
			fmt.Println("请输入要查询的系别：")
			_, _ = fmt.Scan(&sdept)
			maxG, minG := API.GetMAXandMin(sdept)
			fmt.Println(sdept+"系的最好成绩为：", maxG, "最差成绩为：", minG)
		case 18:
			var sdept string
			fmt.Println("请输入要查询的系别：")
			_, _ = fmt.Scan(&sdept)
			fmt.Println(sdept+"系的优秀率为：", API.GetExcellentRate(sdept))
		case 19:
			var sdept string
			fmt.Println("请输入要查询的系别：")
			_, _ = fmt.Scan(&sdept)
			fmt.Println(sdept+"系的不及格人数为：", API.GetFailedSum(sdept))
		case 20:
			var sdept string
			fmt.Println("请输入要查询的系别：")
			_, _ = fmt.Scan(&sdept)
			gos := API.GradeOrder(sdept)
			for _, gs := range gos {
				PrintGradeOrder(gs)
			}
		case 21:
			var sno string
			fmt.Println("请输入要查询学生的学号：")
			_, _ = fmt.Scan(&sno)
			student, courseInfos := API.GetStudentInfo(sno)
			fmt.Println("学生基本信息：")
			PrintStudent(student)
			fmt.Println("选修信息")
			for _, ci := range courseInfos {
				PrintCourseInfo(ci)
			}
		default:
			fmt.Println("请输入正确的选择序号!!!")
		}
	}
}
