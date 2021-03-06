package db

import (
	"context"
	"fmt"
	"time"

	"github.com/phihdn/nc_user/models"

	"github.com/phihdn/nc_user/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(req models.LoginReq) (*models.LoginResp, error) {
	var user models.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"email": req.Email, "password": utils.MD5(req.Password)} //map[string]interface{}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)
	token := utils.GenerateToken(user.ID, user.Phone, user.Email)
	resp := models.LoginResp{&user, token}
	fmt.Println(err, resp)
	return &resp, err
}

func FindUserByID(ID int) (*models.User, error) {
	var user models.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": ID} //map[string]interface{}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)

	return &user, err
}
