package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"db"
	"encoding/json"
	"strconv"

)
var sqlaccount string
func main(){
	sqlaccount = "tako:Takos4incas@/designer?charset=utf8"

	engine := gin.Default()
	engine.POST("/test",test)
	engine.POST("/auth/signup",signup)
	engine.POST("/auth/signin",signin)
	engine.POST("/fnc/getusercontract",getusercontract)
	engine.POST("/fnc/postusercontract",postusercontract)
	engine.POST("/fnc/getmarketcontract",getmarketcontract)
	engine.POST("/fnc/postmarketcontract",postmarketcontract)
	engine.POST("/fnc/buycontract",buycontract)
	engine.Run(":8080")
}

func signup(c *gin.Context){
	name := c.PostForm("userName")
	phone := c.PostForm("userPhone")
	password := c.PostForm("userPsw")

	res := db.SignUp(name,phone,password,sqlaccount)
	if res == 0{
		user := db.QueryUser(name,phone,sqlaccount)
		m,_ := json.Marshal(user)
		c.JSON(http.StatusOK,gin.H{
			"status":"0",
			"msg":"SUCCESS",
			"userInfo":string(m),
			"token":"sodfmamdakfp34k3pk2p42",
		},)
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":"1",
			"msg":"Signup Failed",
		},)
	}
}

func signin(c *gin.Context){
	phone := c.PostForm("userPhone")
	password := c.PostForm("userPsw")
	user := db.SignIn(phone,password,sqlaccount)
	if user.UserId == -1{
		c.JSON(http.StatusOK,gin.H{
			"status":"1",
			"msg":"Signin Failed",
		},)
	}else{
		m,_ := json.Marshal(user)
		c.JSON(http.StatusOK,gin.H{
			"status":"0",
			"msg":"SUCCESS",
			"userInfo":string(m),
			"token":"sodfmamdakfp34k3pk2p42",
		},)
	}
}

func getusercontract(c *gin.Context){
	userIdString := c.PostForm("userId")
	userId,_ := strconv.Atoi(userIdString)
	contracts := db.GetUserContract(userId,sqlaccount)

	m,_ := json.Marshal(contracts)
	c.JSON(http.StatusOK,gin.H{
		"status":"0",
		"msg":"SUCCESS",
		"contracts":string(m),
		"token":"sodfmamdakfp34k3pk2p42",
	},)
}

func postusercontract(c *gin.Context){
	userIdString := c.PostForm("userId")
	userId,_ := strconv.Atoi(userIdString)
	contractName := c.PostForm("contractName")
	contractPrice := c.PostForm("contractPrice")
	contractDesc := c.PostForm("contractDesc")
	contractCont := c.PostForm("contractCont")
	res := db.UploadUserContract(contractName,userId,contractPrice,contractDesc,contractCont,sqlaccount)
	if res == 0{
		c.JSON(http.StatusOK,gin.H{
			"status":"0",
			"msg":"SUCCESS",
		},)
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":"1",
			"msg":"PostContract Failed",
		},)
	}

}

func getmarketcontract(c *gin.Context){
	userIdString := c.PostForm("userId")
	userId,_ := strconv.Atoi(userIdString)
	contracts := db.GetMarketContract(userId,sqlaccount)

	m,_ := json.Marshal(contracts)
	c.JSON(http.StatusOK,gin.H{
		"status":"0",
		"msg":"SUCCESS",
		"contracts":string(m),
		"token":"sodfmamdakfp34k3pk2p42",
	},)
}

func postmarketcontract(c *gin.Context){
	userIdString := c.PostForm("userId")
	userId,_ := strconv.Atoi(userIdString)
	contractName := c.PostForm("contractName")
	contractPrice := c.PostForm("contractPrice")
	contractDesc := c.PostForm("contractDesc")
	contractCont := c.PostForm("contractCont")
	res := db.UploadUserContract(contractName,userId,contractPrice,contractDesc,contractCont,sqlaccount)
	if res == 0{
		db.UploadMarketContract(contractName,userId,contractPrice,contractDesc,contractCont,sqlaccount)
		c.JSON(http.StatusOK,gin.H{
			"status":"0",
			"msg":"SUCCESS",
		},)
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":"1",
			"msg":"PostMarketContract Failed",
		},)
	}
}

func buycontract(c *gin.Context){
	userIdString := c.PostForm("userId")
	userId,_ := strconv.Atoi(userIdString)
	contractIdString := c.PostForm("contractId")
	contractId,_ := strconv.Atoi(contractIdString)
	res := db.BuyContract(userId,contractId,sqlaccount)
	if res == 0{
		c.JSON(http.StatusOK,gin.H{
			"status":"0",
			"msg":"SUCCESS",
		},)
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":"1",
			"msg":"BuyContract Failed",
		},)
	}
}


func test(c *gin.Context){
	bookname := c.PostForm("bookname")
	fmt.Println(bookname)
	c.JSON(http.StatusOK,gin.H{
		"books":"123",
	},
	)
}