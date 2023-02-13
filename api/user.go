package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type Userform struct {
	UserID          string `json:"userID" form:"userID"`
	Title           string `json:"title" form:"title"`
	Site            string `json:"site" form:"site"`
	Confirmpassword string `json:"confirmpassword" form:"confirmpassword"`
	AgentID         string `json:"agentID" form:"agentID"`
	Username        string `json:"username" form:"username"`
	AgentName       string `json:"agentName" form:"agentName"`
	CompanyID       string `json:"companyID" form:"companyID"`
	Permission      string `json:"permission" form:"permission"`
	Currency        string `json:"currency" form:"currency"`
	Role            string `json:"role" form:"role"`
	Token           string `json:"token" form:"token"`
}

func UserLogin(c *gin.Context) {

	values := url.Values{}
	userID := c.PostForm("userID")
	password := c.PostForm("password")
	agentID := c.PostForm("agentID")

	values.Set("userID", userID)
	values.Set("password", password)
	values.Set("agentID", agentID)

	req, err := http.NewRequest("POST", "https://uat-backoffice-api.bpweg.com/login", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	clt := &http.Client{}
	r, _ := clt.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	//读取整个响应体
	body, _ := ioutil.ReadAll(r.Body)
	var data Userform
	json.Unmarshal(body, &data)

	c.JSON(200, gin.H{
		"token": data.Token,
	})

}


func GetOpInfo()