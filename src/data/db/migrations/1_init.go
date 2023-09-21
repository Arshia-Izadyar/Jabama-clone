package migrations

import (
	"database/sql"

	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var log = logger.NewLogger(config.GetConfig())

func Up_01() {
	db := db.GetDB()
	createTables(db)

}

func createTables(db *gorm.DB) {

	tables := []interface{}{}

	tables = addIfNotExists(db, &models.BaseModel{}, tables)
	tables = addIfNotExists(db, &models.User{}, tables)
	tables = addIfNotExists(db, &models.Role{}, tables)
	tables = addIfNotExists(db, &models.UserRole{}, tables)
	tables = addIfNotExists(db, &models.City{}, tables)
	tables = addIfNotExists(db, &models.Province{}, tables)

	err := db.Migrator().CreateTable(tables...)
	if err != nil {
		log.Error(logger.Postgres, logger.Insert, err, nil)
		panic(err)
	}
	CreateDefaultInfo(db)

}

func addIfNotExists(db *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !db.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}
func CreateIfNotExists(db *gorm.DB, r *models.Role) {
	exists := 0
	db.Model(&models.Role{}).Select("1").Where("name = ?", r.Name).First(&exists)
	if exists == 0 {
		db.Model(&models.Role{}).Create(r)
	}
}

func CreateDefaultInfo(db *gorm.DB) {
	defaultRole := &models.Role{
		Name: "default",
	}
	CreateIfNotExists(db, defaultRole)
	adminRole := &models.Role{Name: "admin"}
	CreateIfNotExists(db, adminRole)

	u := &models.User{
		Username:    constants.AdminRoleName,
		Email:       sql.NullString{String: "arshiaa104@gmial.com", Valid: true},
		PhoneNumber: "09108624707",
		FirstName:   sql.NullString{Valid: true, String: "test"},
		LastName:    sql.NullString{Valid: true, String: "test"},
		Activated:   true,
	}
	bs, _ := bcrypt.GenerateFromPassword([]byte("a123"), bcrypt.MinCost)
	u.Password = string(bs)
	createAdmin(db, u, adminRole.Id)
}

func createAdmin(db *gorm.DB, u *models.User, roleId int) {
	exists := 0
	db.Model(&models.User{}).Select("1").Where("username = ?", u.Username).First(&exists)
	if exists == 0 {
		db.Create(u)
		userRole := models.UserRole{UserId: u.Id, RoleId: roleId}
		db.Create(&userRole)
	}
}
