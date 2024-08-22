package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "golang.org/x/text/number"
	"test.com/firstgoproject/internal/db"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

type Product struct {
	Id        primitive.ObjectID `json :"_id" bson:"_id" validare:"required"`
	CreatedAt time.Time          `json: "crearedAt" bson: "created_at" validare:"required"`
	UpdatedAt time.Time          `json: "updatedAt" bson: "updated_at" validare:"required"`
	Title     string             `json: "title" bson: "title" validare:"required", min=12`
}

type Users struct {
	Id	primitive.ObjectID	`json: "_id" bson; "_id" validare:"required"`
	FullName	string 		`json: "fullName" bson:"fullName" validare:"required"`
	Phone		int			`json:"phone" bson: "phone" validare:"required"`
	Password	string		`json:"password" bson:"password" validare:"required"`
	Role		string		`json:"role" bson: "role" validare:"required"`
	CreatedAt	time.Time	`json:"createdtAt" bson:"created_at" validare:"required"`
	UpdatedAt	time.Time	`json:"updatedAt" bson:"updated_at" validare:"required"`

}

type ErrorResponse struct{
	FaildField string
	Tag string
	Value string
}


func ValidareProductStruct (p Product) []*ErrorResponse{
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(p)


	if err!=nil{
		for _, err := range err.(validator.ValidationErrors){
			var element ErrorResponse
			element.FaildField = err.StructNamespace()
			element.Tag  = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)

		}
	}
	return errors
}


func CreatProduct(c *fiber.Ctx) error {
	product := Product{
		Id:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.BodyParser((&product)); err != nil {
		return err
	}

	// err := c.BodyParser((&product))
	// if err != nil{
	// 	return err
	// }

	errors := ValidareProductStruct(product)

	if errors != nil{
		return c.JSON(errors)
	}
	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	fmt.Println("products", product)
	fmt.Println("c: ", c)

	collection := client.Database(db.Database).Collection(string(db.ProductCollection))
	_, err = collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}

	return c.JSON(product)
}

func CreatUsers(c *fiber.Ctx) error {
	user := Users{
		Id:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.BodyParser((&user)); err != nil {
		return err
	}

	// err := c.BodyParser((&user))
	// if err != nil{
	// 	return err
	// }

	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	fmt.Println("users", user)
	fmt.Println("c: ", c)

	collection := client.Database(db.Database).Collection(string(db.UserCollection))
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func GetAllProducts(c *fiber.Ctx) error {
	client , err := db.GetMongoClient()

	var products []*Product
	if err != nil{
		return err
	}

	collection := client.Database((db.Database)).Collection(string(db.ProductCollection))

	cur, err := collection.Find(context.TODO(), bson.D{
		primitive.E{},
	})
	if err != nil{
		return err
	}

	for cur.Next(context.TODO()){
		var p Product
		err:= cur.Decode(&p)

		if err != nil{
			return err
		}
		products= append(products, &p)
	}

	return c.JSON(products)
}
