package config

import (
   "fmt"
   "time"
   "example.com/LAB2/backend/entity"
   "gorm.io/driver/sqlite"
   "gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
   return db
}

func ConnectionDB() {
   database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
   if err != nil {
        panic(fmt.Sprintf("failed to connect database: %v", err))
    }
   fmt.Println("connected database")
   db = database
}

func SetupDatabase() {
   db.AutoMigrate(
       &entity.Users{},
       &entity.Genders{},
   )
   
   GenderMale := entity.Genders{Gender: "Male"}
   GenderFemale := entity.Genders{Gender: "Female"}

   db.FirstOrCreate(&GenderMale, &entity.Genders{Gender: "Male"})
   db.FirstOrCreate(&GenderFemale, &entity.Genders{Gender: "Female"})

   hashedPassword, _ := HashPassword("123456")
   BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")
   User := &entity.Users{
       FirstName: "Software",
       LastName:  "Analysis",
       Email:     "sa@gmail.com",
       Age:       80,
       Password:  hashedPassword,
       BirthDay:  BirthDay,
	   Address: "F11",
       GenderID:  1,
	
   }
   db.FirstOrCreate(User, &entity.Users{
       Email: "sa@gmail.com",
   })
}