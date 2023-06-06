package seeder

import (
	"GoGinStarter/database/seeders"
	"fmt"
	"gorm.io/gorm"
	"path"
	"reflect"
	"runtime"
	"strings"
)

type Seeder interface {
	Seed() error
}

type seeder struct {
	db *gorm.DB
}

var Seeders = []func() (int, interface{}){
	seeders.UserSeederData,
}

func (s *seeder) Seed() error {
	fmt.Println("Seeding into database")
	for _, seederFunc := range Seeders {
		count, _ := seederFunc()
		fmt.Println("Seeding " + getFuncName(seederFunc))
		for x := 0; x < count; x++ {
			_, model := seederFunc()
			if err := s.db.Create(model).Error; err != nil {
				return err
			}
		}
	}
	fmt.Println("Seeding is ended")

	return nil
}

func getFuncName(ff func() (int, interface{})) string {
	f := reflect.ValueOf(ff)
	if f.Kind() == reflect.Func {
		funcName := runtime.FuncForPC(f.Pointer()).Name()
		funcName = path.Base(funcName)
		return strings.Split(funcName, ".")[1]
	}

	return ""
}

func ProvideSeeder(db *gorm.DB) Seeder {
	return &seeder{db}
}
