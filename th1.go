// package main //main 일 겨우 실행 프로그램으로 만듬  //템플릿 없고 data 한번에 처리랑 나눠서 처리 두개다 있는것

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-sql-driver/mysql"
// 	_ "github.com/go-sql-driver/mysql" // 이파일을 다운 받기 위해서는 go.mod 파일이 있는곳에 해야된다. (github.com 위치에 설치하는 것이 아닌 여기에 설치) //설치명령어 go get github.com/go-sql-driver/mysql(설치되면 go.sum 파일이 생김)
// )

// type GatewayDBInfo struct { //conf 안 정보
// 	DBUser   string
// 	DBPW     string
// 	DBIP     string
// 	DBIP4    string
// 	DBPort   int
// 	RESTPort int
// 	DBName   string
// }

// type Tankinfos struct {
// 	Data         string  `json:"data"`
// 	Type         string  `json:"type"`
// 	Differential float64 `json:"differential"`
// 	Pressure     float64 `json:"pressure"`
// }

// var mDBConnSelect *sql.DB

// func readConf(conf *GatewayDBInfo) { //conf.json 파일 읽어들이고 사용하기
// 	file, _ := os.Open("./conf.json") //os import 통해 conf.jon 열기 > 파일을 file 에 저장
// 	defer file.Close()
// 	decoder := json.NewDecoder(file) //마샬링,언마샬링(정수형이나 구조체를 바이트 슬라이스로 변경) 말고 많은 데이터를 처리할때
// 	err := decoder.Decode(&conf)     //json 문자열을 go 밸류로 바꾸는 것 (디코딩)

// 	if err != nil { //에러 처리 팡일 열기에 실패했다면 nil 이 아닌 error 값 반환
// 		fmt.Println("error: ", err)
// 	}

// 	fmt.Println("DB User : ", conf.DBUser) //출력
// 	fmt.Println("DB PassWord : ", conf.DBPW)
// 	fmt.Println("DB IP : ", conf.DBIP)
// 	fmt.Println("DB IP4 : ", conf.DBIP4)
// 	fmt.Println("DB Port : ", conf.DBPort)
// 	fmt.Println("restPort : ", conf.RESTPort)
// 	fmt.Println("DB Name : ", conf.DBName)
// }

// func GetConnector(conf *GatewayDBInfo) *sql.DB { //DB연결
// 	cfg := mysql.Config{
// 		User:                 conf.DBUser, //"279", //readConf때문에 대신하여 사용가능
// 		Passwd:               conf.DBPW,   //"279developer",
// 		Net:                  "tcp",
// 		Addr:                 conf.DBIP + ":" + strconv.Itoa(conf.DBPort), //"127.0.0.1:3306", //strconv.Itoa 숫자를 문자로
// 		Collation:            "utf8mb4_general_ci",
// 		Loc:                  time.UTC,
// 		MaxAllowedPacket:     4 << 20.,
// 		AllowNativePasswords: true,
// 		CheckConnLiveness:    true,
// 		DBName:               conf.DBName, //"dummy"(스키마이름인것을 기억할것, no table)
// 	}
// 	connector, err := mysql.NewConnector(&cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	db := sql.OpenDB(connector)
// 	return db
// }

// func runResever(conf *GatewayDBInfo) {
// 	r := gin.Default()

// 	r.GET("/mydata", func(c *gin.Context) { //3개 가스 최신 데이터

// 		var type1 string
// 		var differential float64
// 		var pressure float64
// 		var tankinfos []Tankinfos

// 		rows, err1 := mDBConnSelect.Query("SELECT type, pressure, differential FROM gas order by timestamp desc limit 3") //tankinfo
// 		if err1 == nil {
// 			for rows.Next() {
// 				err1 := rows.Scan(&type1, &differential, &pressure)
// 				if err1 != nil {
// 					log.Println(err1)
// 				} else {
// 					tankinfo := new(Tankinfos)

// 					tankinfo.Type = type1
// 					tankinfo.Differential = differential
// 					tankinfo.Pressure = pressure

// 					tankinfos = append(tankinfos, *tankinfo)

// 					fmt.Println(type1)
// 					fmt.Println(differential)
// 					fmt.Println(pressure)
// 				}
// 			}
// 			print("sucess")

// 			ret, err2 := json.Marshal(tankinfos)
// 			if err2 != nil {
// 				log.Println(err2)
// 			} else {
// 				c.JSON(http.StatusOK, gin.H{
// 					"error":  "null",
// 					"status": "ok",
// 					"data":   string(ret),
// 				})
// 			}
// 		}
// 		defer rows.Close()
// 	})
// 	r.GET("/mydata/:type", func(c *gin.Context) { //특정 가스의 최신 데이터

// 		strtype := c.Param("type")

// 		var data string
// 		// var type1 string
// 		// var differential float64
// 		// var pressure float64
// 		var tankinfos []Tankinfos

// 		rows, err1 := mDBConnSelect.Query("SELECT JSON_OBJECT('type', type,'pressure', pressure, 'differential', differential) FROM gas where type = ? order by timestamp desc limit 1", strtype) //tankinfo
// 		if err1 == nil {
// 			for rows.Next() {
// 				err1 := rows.Scan(&data) //&type1, &differential, &pressure
// 				if err1 != nil {
// 					log.Println(err1)
// 				} else {
// 					tankinfo := new(Tankinfos)
// 					tankinfo.Data = data
// 					tankinfos = append(tankinfos, *tankinfo)

// 					fmt.Println(data)

// 					c.JSON(http.StatusOK, gin.H{
// 						"status": "ok",
// 						"error":  "null",
// 						"data":   string(data),
// 					})
// 				}
// 			}
// 		}
// 		defer rows.Close()
// 	})
// 	r.GET("/mydata_all", func(c *gin.Context) { //가스 모든 데이터

// 		c.HTML(http.StatusOK, "header.html", gin.H{
// 			"title": "3개 가스 모든 데이터",
// 		})
// 		var data string
// 		var tankinfos []Tankinfos

// 		rows, err1 := mDBConnSelect.Query("SELECT JSON_OBJECT('type', type, 'pressure', pressure, 'differential', differential) FROM gas") //tankinfo
// 		if err1 == nil {
// 			for rows.Next() {
// 				err1 := rows.Scan(&data) //&type1, &differential, &pressure
// 				if err1 != nil {
// 					log.Println(err1)
// 				} else {
// 					tankinfo := new(Tankinfos)
// 					tankinfo.Data = data
// 					tankinfos = append(tankinfos, *tankinfo)

// 					fmt.Println(data)

// 					c.HTML(http.StatusOK, "index.html", gin.H{
// 						"message": data,
// 					})
// 				}
// 			}
// 		}
// 		defer rows.Close()
// 	})
// 	r.GET("/mydata_all/:type", func(c *gin.Context) { //특정 가스의 모든 데이터

// 		strtype := c.Param("type")

// 		c.HTML(http.StatusOK, "header.html", gin.H{
// 			"title": strtype + "모든 데이터",
// 		})
// 		var data string
// 		var tankinfos []Tankinfos

// 		rows, err1 := mDBConnSelect.Query("SELECT JSON_OBJECT('type', type,'pressure', pressure, 'differential', differential) FROM gas where type =?", strtype) //tankinfo
// 		if err1 == nil {
// 			for rows.Next() {
// 				err1 := rows.Scan(&data)
// 				if err1 != nil {
// 					log.Println(err1)
// 				} else {
// 					tankinfo := new(Tankinfos)
// 					tankinfo.Data = data
// 					tankinfos = append(tankinfos, *tankinfo)

// 					fmt.Println(data)

// 					c.HTML(http.StatusOK, "index.html", gin.H{
// 						"message": data,
// 					})
// 				}
// 			}
// 		}
// 		defer rows.Close()
// 	})
// 	r.Run(":7000")
// }

// func main() { //Entry Point(시작점)
// 	conf := GatewayDBInfo{} //구조체 객체 생성
// 	readConf(&conf)         //구조체 함수에 구조체 대입(변수가 구조체로 들어가 있기 때문)

// 	db := GetConnector(&conf) //(구조체 변수 대입)
// 	err := db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	mDBConnSelect = GetConnector(&conf)
// 	runResever(&conf)
// }
