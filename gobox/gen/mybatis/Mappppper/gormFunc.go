package Mappppper

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	GormDB, e := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/im?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true")
	if e != nil {
		panic(e)
	} else {
		DB = GormDB
	}
}

var (
	DB            *gorm.DB
	AlbumsDB      = AlbumsOpe{DB: DB}
	BizPropertyDB = BizPropertyOpe{DB: DB}
	PhotosDB      = PhotosOpe{DB: DB}
	BizActivityDB = BizActivityOpe{DB: DB}
	UsersDB       = UsersOpe{DB: DB}
)

type PhotosOpe struct {
	DB *gorm.DB
}

func (a PhotosOpe) GetByID(id int) (Photos, error) {
	obj := Photos{}
	err := a.DB.Model(&Photos{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a PhotosOpe) ListByLimit(o, l int) ([]Photos, error) {
	var list []Photos
	err := a.DB.Model(&Photos{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a PhotosOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Photos{}).Find(&[]Photos{}).Count(&count).Error
	return count, err
}
func (a PhotosOpe) UpdateById(up Photos) error {
	return a.DB.Model(&Photos{}).Update(up).Error
}
func (a PhotosOpe) DeleteById(id int) error {
	return a.DB.Model(&Photos{}).Where("`id`=?", id).Delete(nil).Error
}
func (a PhotosOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Photos{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------photos------------------------//
type AlbumsOpe struct {
	DB *gorm.DB
}

func (a AlbumsOpe) GetByID(id int) (Albums, error) {
	obj := Albums{}
	err := a.DB.Model(&Albums{}).Where("`cid`=?", id).Find(&obj).Error
	return obj, err
}

func (a AlbumsOpe) ListByLimit(o, l int) ([]Albums, error) {
	var list []Albums
	err := a.DB.Model(&Albums{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a AlbumsOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Albums{}).Find(&[]Albums{}).Count(&count).Error
	return count, err
}
func (a AlbumsOpe) UpdateById(up Albums) error {
	return a.DB.Model(&Albums{}).Update(up).Error
}
func (a AlbumsOpe) DeleteById(id int) error {
	return a.DB.Model(&Albums{}).Where("`cid`=?", id).Delete(nil).Error
}
func (a AlbumsOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Albums{}).Where("`cid` in (?)", id).Delete(nil).Error
}

//------------------------albums------------------------//
type BizPropertyOpe struct {
	DB *gorm.DB
}

func (a BizPropertyOpe) GetByID(id int) (BizProperty, error) {
	obj := BizProperty{}
	err := a.DB.Model(&BizProperty{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a BizPropertyOpe) ListByLimit(o, l int) ([]BizProperty, error) {
	var list []BizProperty
	err := a.DB.Model(&BizProperty{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a BizPropertyOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&BizProperty{}).Find(&[]BizProperty{}).Count(&count).Error
	return count, err
}
func (a BizPropertyOpe) UpdateById(up BizProperty) error {
	return a.DB.Model(&BizProperty{}).Update(up).Error
}
func (a BizPropertyOpe) DeleteById(id int) error {
	return a.DB.Model(&BizProperty{}).Where("`id`=?", id).Delete(nil).Error
}
func (a BizPropertyOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&BizProperty{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------biz_property------------------------//
type BizActivityOpe struct {
	DB *gorm.DB
}

func (a BizActivityOpe) GetByID(id int) (BizActivity, error) {
	obj := BizActivity{}
	err := a.DB.Model(&BizActivity{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a BizActivityOpe) ListByLimit(o, l int) ([]BizActivity, error) {
	var list []BizActivity
	err := a.DB.Model(&BizActivity{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a BizActivityOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&BizActivity{}).Find(&[]BizActivity{}).Count(&count).Error
	return count, err
}
func (a BizActivityOpe) UpdateById(up BizActivity) error {
	return a.DB.Model(&BizActivity{}).Update(up).Error
}
func (a BizActivityOpe) DeleteById(id int) error {
	return a.DB.Model(&BizActivity{}).Where("`id`=?", id).Delete(nil).Error
}
func (a BizActivityOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&BizActivity{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------biz_activity------------------------//
type UsersOpe struct {
	DB *gorm.DB
}

func (a UsersOpe) GetByID(id int) (Users, error) {
	obj := Users{}
	err := a.DB.Model(&Users{}).Where("`id`=?", id).Find(&obj).Error
	return obj, err
}

func (a UsersOpe) ListByLimit(o, l int) ([]Users, error) {
	var list []Users
	err := a.DB.Model(&Users{}).Offset(o).Limit(l).Find(&list).Error
	return list, err
}
func (a UsersOpe) Count() (int, error) {
	var count int
	err := a.DB.Model(&Users{}).Find(&[]Users{}).Count(&count).Error
	return count, err
}
func (a UsersOpe) UpdateById(up Users) error {
	return a.DB.Model(&Users{}).Update(up).Error
}
func (a UsersOpe) DeleteById(id int) error {
	return a.DB.Model(&Users{}).Where("`id`=?", id).Delete(nil).Error
}
func (a UsersOpe) DeleteByIds(id ...int) error {
	return a.DB.Model(&Users{}).Where("`id` in (?)", id).Delete(nil).Error
}

//------------------------users------------------------//
