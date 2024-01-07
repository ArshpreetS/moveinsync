package handlers

import (
	"fmt"
	"net/http"

	"github.com/ArshpreetS/moveinsync/gateway/database"
	"github.com/ArshpreetS/moveinsync/gateway/models"
	"github.com/ArshpreetS/moveinsync/gateway/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(c *gin.Context) {

	var data models.NewUser

	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Wrong Request Body")
		return
	}

	data.UserID = uuid.New().String()
	db_client := database.GetDBClient()
	defer db_client.Close()

	// check if email or username pre exists
	row, err := db_client.Query(fmt.Sprintf("SELECT * from users where email='%s' or username='%s'", string(data.Email), data.Username))

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "There was error while querying")
		return
	}
	defer row.Close()

	if row.Next() {
		c.JSON(http.StatusConflict, struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{
			Status:  0,
			Message: "There already exists such user",
		})
		return
	}

	stmt, err := db_client.Prepare("INSERT INTO users (id, username,email,password) VALUES (?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error with preparing insert into mysql")
		return
	}
	defer stmt.Close()

	hashedPass, err := utils.GetHashPass(data.Password)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "There was some error with hashing password")
		return
	}

	token, err := utils.CreateToken(data.Username, data.Email, data.UserID)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "There was some error while creating jwt token")
		return
	}

	_, err = stmt.Exec(data.UserID, data.Username, data.Email, hashedPass)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error while adding data to database")
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Token:  token,
		Status: 1,
	})
}
