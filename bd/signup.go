package bd

import (
	"fmt"

	"github.com/CrisValens666/valensuser/models"
	"github.com/CrisValens666/valensuser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) Values ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySql() + "')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("SignUp > Ejecución Exitosa")
	return nil
}
