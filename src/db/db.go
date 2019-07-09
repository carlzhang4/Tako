package db

import(
    "database/sql"
    "fmt"
    _"github.com/go-sql-driver/mysql"
    "strconv"
)

//sqlaccount := "root:19050817@/designer?charset=utf8"
type UserInfo struct{
    UserName string
    UserPhone string
    UserId int
}
type Contract struct{
    Name string
    UserID int
    Price string
    Description string
    Content string
    ID int

}

func SignUp(name string,phone string,password string,sqlaccount string) int{
	db, err := sql.Open("mysql", sqlaccount)
    checkErr(err)

    rows, err := db.Query("SELECT * FROM users WHERE name = ? or phone = ?", name,phone)
    for rows.Next() {
        return 1
    }

	stmt, err := db.Prepare("INSERT users SET name=?,phone=?,password=?,money=100")
    checkErr(err)
    res, err := stmt.Exec(name,phone,password)
    checkErr(err)
    afctedRows,err := res.RowsAffected()
    checkErr(err)
    fmt.Println(afctedRows)
    return 0;
}
func QueryUser(name string,phone string,sqlaccount string)UserInfo{
    db, err := sql.Open("mysql", sqlaccount)
	checkErr(err)
    rows, err := db.Query("SELECT name,phone,id FROM users WHERE name = ? and phone = ?", name,phone)
    var resName string
    var resPhone string
    var resId int
    for rows.Next() {
        rows.Scan(&resName,&resPhone,&resId)
        return UserInfo{resName,resPhone,resId}
    }
    return UserInfo{}
}

func SignIn(phone string,password string,sqlaccount string)UserInfo{
    db, err := sql.Open("mysql", sqlaccount)
	checkErr(err)
    rows, err := db.Query("SELECT name,phone,id FROM users WHERE phone = ? and password = ?", phone,password)
    var resName string
    var resPhone string
    var resId int
    for rows.Next() {
        rows.Scan(&resName,&resPhone,&resId)
        return UserInfo{resName,resPhone,resId}
    }
    return UserInfo{"","",-1}
}

func UploadUserContract(name string,userID int,price string,description string,content string,sqlaccount string)int{
    db, err := sql.Open("mysql", sqlaccount)
    checkErr(err)

    rows, err := db.Query("SELECT * FROM contract WHERE name = ? and userID = ?", name,userID)
    for rows.Next() {
        return 1
    }
    stmt, err := db.Prepare("INSERT contract SET name=?,userID=?,price=?,description=?,content=?")
    checkErr(err)
    res, err := stmt.Exec(name,userID,price,description,content)
    checkErr(err)
    afctedRows,err := res.RowsAffected()
    checkErr(err)
    fmt.Println(afctedRows)
    return 0
}

func UploadMarketContract(name string,userID int,price string,description string,content string,sqlaccount string)int{
    db, err := sql.Open("mysql", sqlaccount)
    checkErr(err)

    rows, err := db.Query("SELECT * FROM market WHERE name = ? and userID = ?", name,userID)
    for rows.Next() {
        fmt.Println("1111")
        return 1
    }
    fmt.Println("2222")
    stmt, err := db.Prepare("INSERT market SET name=?,userID=?,price=?,description=?,content=?")
    checkErr(err)
    res, err := stmt.Exec(name,userID,price,description,content)
    checkErr(err)
    afctedRows,err := res.RowsAffected()
    checkErr(err)
    fmt.Println(afctedRows)
    return 0
}

func GetUserContract(userID int,sqlaccount string) []Contract{
    db, err := sql.Open("mysql", sqlaccount)
    checkErr(err)
    fmt.Println(userID)
    rows, err := db.Query("SELECT name,price,description,content,id FROM contract WHERE userID = ?", userID)

    var name string
    var price string
    var description string
    var content string
    var id int
    var contracts []Contract

    for rows.Next() {
        rows.Scan(&name,&price,&description,&content,&id)
        fmt.Println(contracts)
        contracts = append(contracts,Contract{name,userID,price,description,content,id})
    }
    return contracts
}

func GetMarketContract(userID int,sqlaccount string) []Contract{
    db, err := sql.Open("mysql", sqlaccount)
    checkErr(err)
    rows, err := db.Query("SELECT name,price,description,content,id FROM market WHERE userID = ?", userID)

    var name string
    var price string
    var description string
    var content string
    var id int
    var contracts []Contract

    for rows.Next() {
        rows.Scan(&name,&price,&description,&content,&id)
        fmt.Println(contracts)
        contracts = append(contracts,Contract{name,userID,price,description,content,id})
    }
    return contracts
}
func BuyContract(userID int,contractID int,sqlaccount string)int{
    db, err := sql.Open("mysql", sqlaccount)
    checkErr(err)
    rows, err := db.Query("SELECT name,price,userID,description,content FROM market WHERE id = ?", contractID)
    var name string
    var price string
    var sellerID int
    var description string
    var content string

    var money string
    for rows.Next(){
        rows.Scan(&name,&price,&sellerID,&description,&content)
    }
    fmt.Println(userID,contractID)
    if sellerID <=0{
        return 1
    }
    rows, err = db.Query("SELECT * FROM contract WHERE name = ? and userID = ?", name,userID)
    for rows.Next(){
        return 1
    }
    rows,err = db.Query("SELECT money FROM users WHERE id = ?",userID)
    for rows.Next(){
        rows.Scan(&money)
    }
    moneynum,err := strconv.ParseFloat(money,64)
    checkErr(err)
    pricenum,err := strconv.ParseFloat(price,64)
    checkErr(err)
    if moneynum<pricenum{
        return 1
    }else{
        stmt, err := db.Prepare("INSERT contract SET name=?,userID=?,price=?,description=?,content=?")
        checkErr(err)
        _, err = stmt.Exec(name,userID,price,description,content)
        checkErr(err)

        stmt, err = db.Prepare("UPDATE users SET money=? where id = ?")
        checkErr(err)
        _, err = stmt.Exec(moneynum-pricenum,userID)
        checkErr(err)

        rows,err = db.Query("SELECT money FROM users WHERE id = ?",sellerID)
        for rows.Next(){
            rows.Scan(&money)
        }
        moneynum,err := strconv.ParseFloat(money,64)
        checkErr(err)

        stmt, err = db.Prepare("UPDATE users SET money=? where id = ?")
        checkErr(err)
        _, err = stmt.Exec(moneynum+pricenum,sellerID)
        checkErr(err)
        return 0
    }
}

func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
