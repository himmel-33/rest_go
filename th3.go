// package main //main 일 겨우 실행 프로그램으로 만듬  //mydata 데이터 여러개 들어와도 검색할 수 있게 변수설정 배열설정
// //하나의 gin 안에 query문 3개 작성, query문 값을 map을 통해 키 값으로 받고 출력
// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"
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
// 	Data string `json:"data"`
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

// 		var data string
// 		var types []string
// 		var a map[string]string
// 		a = make(map[string]string)

// 		rows, err1 := mDBConnSelect.Query("select distinct type from gas;")
// 		if err1 == nil {
// 			for rows.Next() {
// 				err1 := rows.Scan(&data)
// 				if err1 != nil {
// 					log.Println(err1)
// 				} else {
// 					types = append(types, data)
// 					fmt.Println(data)
// 				}
// 			}
// 		}
// 		for _, tyname := range types {
// 			rows, err1 := mDBConnSelect.Query("SELECT JSON_OBJECT('type', type,'pressure', pressure, 'differential', differential) FROM gas where type = ? order by date desc limit 1", tyname)
// 			if err1 == nil {
// 				for rows.Next() {
// 					err1 := rows.Scan(&data)
// 					if err1 != nil {
// 						log.Println(err1)
// 					} else {
// 						a[tyname] = data
// 						fmt.Println(data)
// 					}
// 				}
// 			}
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": "OK",
// 			"error":  "null",
// 			"data":   a,
// 		})
// 		defer rows.Close()
// 	})
// 	r.GET("/mydata/:type", func(c *gin.Context) { //특정 가스의 최신 데이터

// 		strtype := c.Param("type")

// 		var data string

// 		rows, err1 := mDBConnSelect.Query("SELECT JSON_OBJECT('type', type,'pressure', pressure, 'differential', differential) FROM gas where type = ? order by date desc limit 1", strtype)
// 		if err1 == nil {
// 			for rows.Next() {
// 				err1 := rows.Scan(&data)
// 				if err1 != nil {
// 					log.Println(err1)
// 				} else {
// 					fmt.Println(data)
// 				}
// 			}
// 			//\ 역슬래쉬 빼는 부분 성공
// 			byt, _ := json.Marshal(data)
// 			body := strings.Replace(string(byt), `\"`, `"`, -1)
// 			body = body[1 : len(body)-1]
// 			body = strings.Replace(body, `\\`, `\`, -1)
// 			var data1 map[string]interface{}
// 			if err := json.Unmarshal([]byte(body), &data1); err != nil {
// 				fmt.Println(err)
// 			} else {
// 				fmt.Println(data1)
// 			}
// 			//
// 			c.JSON(http.StatusOK, gin.H{
// 				"status": "OK",
// 				"error":  "null",
// 				"data":   data1,
// 			})
// 		}
// 		defer rows.Close()
// 	})
// 	r.GET("/mydata_all", func(c *gin.Context) { //가스 모든 데이터

// 		var data string
// 		var a []string

// 		rows, err1 := mDBConnSelect.Query("SELECT JSON_OBJECT('type', type, 'pressure', pressure, 'differential', differential) FROM gas")
// 		if err1 == nil {
// 			for rows.Next() {
// 				err1 := rows.Scan(&data)
// 				if err1 != nil {
// 					log.Println(err1)
// 				} else {
// 					a = append(a, data)
// 					fmt.Println(data)
// 				}
// 			}
// 		}
// 		//\ 역슬래쉬 빼는 부분
// 		var arrstr string
// 		arrstr = strings.Join(a, " ") //배열 문자열 변환
// 		byt, _ := json.Marshal(arrstr)
// 		body := strings.Replace(string(byt), `\"`, `"`, -1)
// 		body = body[1 : len(body)-1]
// 		body = strings.Replace(body, `\\`, `\`, -1)
// 		var data1 map[string]interface{}
// 		if err := json.Unmarshal([]byte(body), &data1); err != nil {
// 			fmt.Println(err)
// 		} else {
// 			fmt.Println(data1)
// 		}
// 		//
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": "OK",
// 			"error":  "null",
// 			"data":   data1,
// 		})
// 		defer rows.Close()
// 	})
// 	r.GET("/mydata_all/:type", func(c *gin.Context) { //특정 가스의 모든 데이터

// 		strtype := c.Param("type")

// 		var data string
// 		var a []string

// 		rows, err1 := mDBConnSelect.Query("SELECT JSON_OBJECT('type', type,'pressure', pressure, 'differential', differential) FROM gas where type =?", strtype) //tankinfo
// 		if err1 == nil {
// 			for rows.Next() {
// 				err1 := rows.Scan(&data)
// 				if err1 != nil {
// 					log.Println(err1)
// 				} else {
// 					a = append(a, data)
// 					fmt.Println(data)
// 				}
// 			}
// 			//\ 역슬래쉬 빼는 부분
// 			var arrstr string
// 			arrstr = strings.Join(a, " ") //배열 문자열 변환
// 			byt, _ := json.Marshal(arrstr)
// 			body := strings.Replace(string(byt), `\"`, `"`, -1)
// 			body = body[1 : len(body)-1]
// 			body = strings.Replace(body, `\\`, `\`, -1)
// 			var data1 map[string]interface{}
// 			if err := json.Unmarshal([]byte(body), &data1); err != nil {
// 				fmt.Println(err)
// 			} else {
// 				fmt.Println(data1)
// 			}
// 			//
// 			c.JSON(http.StatusOK, gin.H{
// 				"status": "OK",
// 				"error":  "null",
// 				"data":   data1,
// 			})
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
