# gofactory

Package to create factories for database testing in golang and gorm with ease

**Make 5 models:**

```go
package main

import "fmt"
import "github.com/vsvp21/gofactory/v1"

func main() {
	orders := gofactory.Make[*models.Order](factory.Order, 5)

	fmt.Println(orders)
}
```

**Make 5 models and override fields:**

```go
package main

import "fmt"
import "github.com/vsvp21/gofactory/v1"

func main() {
	overrideModel := &models.Order{Restaurant: &models.Restaurant{Name: "OVERRIDE"}}
	orders := gofactory.MakeOverride[*models.Order](factory.Order, 1, overrideModel)

	fmt.Println(orders)
}
```

**Create 5 models in DB:**

```go
package main

import "fmt"
import "github.com/vsvp21/gofactory/v1"

func main() {
	orders := gofactory.Create[*models.Order](factory.Order, 5)

	fmt.Println(orders)
}
```

**Create 5 models in DB and override fields:**

```go
package main

import "fmt"
import "github.com/vsvp21/gofactory/v1"

func main() {
	overrideModel := &models.Order{Restaurant: &models.Restaurant{Name: "OVERRIDE"}}
	orders := gofactory.CreateOverride[*models.Order](factory.Order, 1, overrideModel)

	fmt.Println(orders)
}
```


**How to create factory**

```go
package models

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Price int
}

func ProductFactory() *Product {
	// use gofakeit to fake data
	return &Product{Price: gofakeit.Number(1, 1000)}
}
```


**How to create factory with relations**

```go
package models

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/vsvp21/gofactory/v1"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Total      uint
	OrderItems []*OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductID uint
	Price     uint
	Amount    uint
	OrderID   uint
	Order     *Order `gorm:"foreignKey:OrderID"`
}

func OrderItemFactory() *OrderItem {
	return &OrderItem{
		Price:   uint(gofakeit.Price(0, 10000)),
		Amount:  gofakeit.UintRange(1, 3),
	}
}

func OrderFactory() *Order {
	return &Order{
		Total:      uint(gofakeit.Price(0, 10000)),
		// make 5 OrderItem models
		OrderItems: gofactory.Make[*OrderItem](OrderItemFactory, 5),
	}
}

```