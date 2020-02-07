package db

import (
	"context"
	"fmt"
	"github.com/phihdn/nc_user/model"
	"time"

	"github.com/phihdn/nc_user/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(req model.LoginReq) (*model.LoginResp, error) {
	var user model.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"email": req.Email, "password": utils.MD5(req.Password)} //map[string]interface{}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)
	token := utils.GenerateToken(user.ID, user.Phone, user.Email)
	resp := model.LoginResp{&user, token}
	fmt.Println(err, resp)
	return &resp, err
}

func FindUserByID(ID int) (*model.User, error) {
	var user model.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": ID} //map[string]interface{}
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&user)

	return &user, err
}
