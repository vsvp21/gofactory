package gofactory

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"testing"
)

type User struct {
	ID   int
	Name string
}

func UserFactory() *User {
	return &User{
		Name: gofakeit.Name(),
	}
}

func TestMakeOverride(t *testing.T) {
	c := rand.Intn(5)
	name := "Test"
	users := MakeOverride[*User](UserFactory, c, &User{Name: name})

	if len(users) != c {
		t.Errorf("len assertion failed\nexpected: %d\nactual: %d", c, len(users))
	}

	for _, m := range users {
		if users[0].Name != name {
			t.Errorf("model assertion failed\nexpected: %s\nactual: %s", name, m.Name)
		}
	}
}

func TestMake(t *testing.T) {
	c := rand.Intn(5)
	users := Make[*User](UserFactory, c)

	if len(users) != c {
		t.Errorf("len assertion failed\nexpected: %d\nactual: %d", c, len(users))
	}

	for _, m := range users {
		if m.Name == "" {
			t.Errorf("model assertion failed, empty")
		}
	}
}

func TestCreateOverride(t *testing.T) {
	dsn := "host=localhost user=db_user password=secretsecret dbname=factory port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	c := rand.Intn(5)
	name := "Test"

	users, err := CreateOverride[*User](db, UserFactory, c, &User{Name: name})
	if err != nil {
		t.Error(err)
	}

	if len(users) != c {
		t.Errorf("len assertion failed\nexpected: %d\nactual: %d", c, len(users))
	}

	for _, m := range users {
		if m.Name != name {
			t.Errorf("model assertion failed\nexpected: %s\nactual: %s", name, m.Name)
		}
	}

	err = db.Migrator().DropTable(&User{})
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	dsn := "host=localhost user=db_user password=secretsecret dbname=factory port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	c := rand.Intn(5)

	users, err := Create[*User](db, UserFactory, c)
	if err != nil {
		t.Error(err)
	}

	if len(users) != c {
		t.Errorf("len assertion failed\nexpected: %d\nactual: %d", c, len(users))
	}

	for _, m := range users {
		if m.Name == "" {
			t.Errorf("model assertion failed, empty")
		}
	}
}
