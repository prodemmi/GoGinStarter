package seeders

import (
	"GoGinStarter/app/models"
	"github.com/bxcodec/faker/v4"
)

func UserSeederData() (int, interface{}) {
	return 5, &models.User{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  "$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi", // password
	}
}
