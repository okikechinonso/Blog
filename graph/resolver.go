package graph

import (
	"digitalocean/graph/model"
	"fmt"

	"github.com/go-pg/pg/v10"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *pg.DB
}

func (r *mutationResolver) GetUserByField(field, value string)(*model.User,error){
	user := model.User{}

	err := r.DB.Model(&user).Where(fmt.Sprintf("%v = ?",field),value).First()
	return &user, err 

}

func (r *mutationResolver) UpdateUser(user *model.User)(*model.User,error){
	_,err := r.DB.Model(user).Where("id = ?",user.ID).Update()
	return user,err
}
