// Package db Tương tác với LIS database
//
// @Author	Đào Văn Thanh
// @Date		2019-05-31
//
package db.mysql

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//MySQLLis Lis database
func MySQLLis() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "itcntt@123"
	dbName := "lis"
	domain := "192.168.0.16"

	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, domain, dbName)
	db, err = sql.Open(dbDriver, conn)

	return db, err
}

//ValidUser kiểm tra thông tin user trên LIS
func ValidUser(user, pwd string) (userID int64, err error) {
	db, err := MySQLLis()
	if err != nil {
		return 0, errors.New("Không kết nối với database")
	}
	defer db.Close()

	matkhau := strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(pwd))))
	sqlString := "SELECT employee_id FROM employees WHERE user_name=? AND password=?"
	err = db.QueryRow(sqlString, user, matkhau).Scan(&userID)
	if err != nil {
		return 0, err
	}

	if userID != 0 {
		return userID, nil
	}
	return 0, errors.New("username hoặc password không hợp lệ")
}
