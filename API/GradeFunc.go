package API

import (
	"SMS/Model"
)

// 按系返回平均成绩
func GetAvg(sdept string) (avg float32) {
	Model.Db.Raw("SELECT AVG(grade) FROM sc, student WHERE student.sdept=? "+
		"AND sc.sno = student.sno", sdept).Scan(&avg)
	return
}

// 按系返回最好成绩和最差成绩
func GetMAXandMin(sdept string) (maxG, minG float32) {
	Model.Db.Raw("SELECT MAX(grade) FROM sc, student WHERE student.sdept=? "+
		"AND sc.sno = student.sno", sdept).Scan(&maxG)
	Model.Db.Raw("SELECT MIN(grade) FROM sc, student WHERE student.sdept=? "+
		"AND sc.sno = student.sno", sdept).Scan(&minG)
	return
}

// 按系查询优秀率即成绩在90分以上的占比
func GetExcellentRate(sdept string) (rate float32) {
	var sum, temp int
	Model.Db.Raw("SELECT COUNT(*) FROM sc, student WHERE student.sdept=? "+
		"AND sc.sno = student.sno", sdept).Scan(&sum)
	Model.Db.Raw("SELECT COUNT(*) FROM sc, student WHERE student.sdept=?"+
		" AND sc.sno = student.sno AND grade >= 90", sdept).Scan(&temp)
	rate = float32(temp) / float32(sum)
	return
}

// 按系查询不及格人数
func GetFailedSum(sdept string) (sum int) {
	Model.Db.Raw("SELECT COUNT(*) FROM (SELECT DISTINCT sc.sno from sc, student WHERE student.sdept=?"+
		" AND sc.sno = student.sno) as temp", sdept).Scan(&sum)
	return
}
