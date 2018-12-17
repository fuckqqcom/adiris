package mysql

import (
	"fmt"
	"github.com/xormplus/xorm"
)

func CheckErrInsert(s *xorm.Session, a interface{}) bool {

	if _, err := s.Insert(a); err != nil {
		if err := recover(); err != nil {
			//TODO
			fmt.Println("error--->", err)
			//conf.Logger.Warn("sql panic", err)
		}
		fmt.Println("error--->", err, a)
		//conf.Logger.Warn("sql", err)
		s.Rollback()
		return false
	}
	return true
}

func CheckErrUpdate(s *xorm.Session, a interface{}, id int) bool {
	if _, err := s.ID(id).Update(a); err != nil {
		if err := recover(); err != nil {
			//TODO
			fmt.Println(err)
			//conf.Logger.Warn("sql panic", err)
		}
		//conf.Logger.Warn("sql", err)
		s.Rollback()
		return false
	}
	return true
}
