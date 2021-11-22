package Model

import "fmt"

type SC struct {
	Sno string	`gorm:"type:char(9)"`
	Cno string	`gorm:"type:char(4)"`
	Grade int	`gorm:"type:smallint"`
}

func AddSc(sc *SC) error{
	err := Db.Create(&sc).Error
	if err != nil{
		return err
	}
	fmt.Println("加入成功！")
	return nil
}

func GetInfoBySno(sno string) []SC{
	var scs []SC
	Db.Where("sno=?", sno).Find(&scs)
	return scs
}

func GetInfoByCno(cno string) []SC{
	var scs []SC
	Db.Where("cno=?", cno).Find(&scs)
	return scs
}

func GetOneSC(cno, sno string) (SC, error){
	var sc SC
	err := Db.Where("cno = ? AND sno = ?", cno, sno).First(&sc).Error
	if err != nil{
		return SC{}, err
	}
	return sc, nil
}

func GetSC() (scs []SC) {
	Db.Find(&scs)
	return
}

func DeleteOneSC(cno, sno string) error  {
	err := Db.Where("cno = ? AND sno = ?", cno, sno).Delete(&SC{}).Error
	if err != nil{
		fmt.Println("删除失败！")
		return err
	}
	fmt.Println("删除成功！")
	return nil
}

func UpdateSC(cno, sno string, data map[string]interface{}) error{
	err := Db.Model(&SC{}).Where("cno=? AND sno=?", cno, sno).Updates(&data).Error
	if err != nil{
		return err
	}
	return nil
}



