package helper

import "gorm.io/gorm"

func CommitOrRollback(db *gorm.DB) {
	err := recover()
	if err != nil {
		error_ := db.Rollback().Error
		if error_ != nil {
			return
		}
	} else {
		error := db.Commit().Error
		if error != nil {
			panic(error)
		}
	}
}
