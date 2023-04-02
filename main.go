package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"helloweb/ctrl"
	"html/template"
	"log"
	"net/http"
)

//	type H struct {
//		Code int         `json:"code"`
//		Data interface{} `json:"data"`
//		Msg  string      `json:"msg"`
//	}
//
//	func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//
//		h := H{
//			Code: code,
//			Data: data,
//			Msg:  msg,
//		}
//		ret, err := json.Marshal(h)
//		if err != nil {
//			log.Println(err.Error())
//		}
//		w.Write([]byte(ret))
//	}
//
//	func userLogin(writer http.ResponseWriter, request *http.Request) {
//		request.ParseForm()
//		mobile := request.PostForm.Get("mobile")
//		passwd := request.PostForm.Get("passwd")
//		loginOk := false
//		if mobile == "123" && passwd == "123" {
//			loginOk = true
//		}
//
//		if !loginOk {
//			Resp(writer, -1, nil, "密码不正确")
//		} else {
//			data := make(map[string]interface{})
//			data["id"] = 1
//			data["token"] = "test"
//			Resp(writer, 0, data, "")
//		}
//
// }
//
//	func RegisterView1() {
//		tpl, _ := template.ParseGlob("view/**/*")
//		for _, v := range tpl.Templates() {
//			name := v.Name()
//			http.HandleFunc(name, func(writer http.ResponseWriter, request *http.Request) {
//				tpl.ExecuteTemplate(writer, name, nil)
//			})
//		}
//
// }
//
//	func RegisterView() {
//		//一次解析出全部模板
//		tpl, err := template.ParseGlob("view/**/*")
//		if nil != err {
//			log.Fatal(err)
//		}
//		//通过for循环做好映射
//		for _, v := range tpl.Templates() {
//			//
//			tplname := v.Name()
//			fmt.Println("HandleFunc     " + v.Name())
//			http.HandleFunc(tplname, func(w http.ResponseWriter,
//				request *http.Request) {
//				//
//				fmt.Println("parse     " + v.Name() + "==" + tplname)
//				err := tpl.ExecuteTemplate(w, tplname, nil)
//				if err != nil {
//					log.Fatal(err.Error())
//				}
//			})
//		}
//
// }
//
// var DbEngine *xorm.Engine
//
//	func init() {
//		drivername := "mysql"
//		Dsname := "root:root@(127.0.0.1:3306)"
//		DbEngine, err := xorm.NewEngine(drivername, Dsname)
//		if nil != err {
//			log.Fatal(err.Error())
//		}
//		DbEngine.ShowSQL(true)
//		DbEngine.SetMaxIdleConns(2)
//		println("init database")
//	}
//
//	func main() {
//		//绑定请求和处理函数
//		http.HandleFunc("/user/login", userLogin)
//		//提供静态资源目录绑定支持
//		//http.Handle("/", http.FileServer(http.Dir(".")))
//		http.Handle("/asset/", http.FileServer(http.Dir(".")))
//
//		//http.HandleFunc("/user/login.shtml", func(w http.ResponseWriter, r *http.Request) {
//		//	//tpl := template.New("user/login.shtml")
//		//	//tpl.ParseFiles("./view/user/login.html")
//		//	//t1, err := template.ParseFiles("view/user/login.html")
//		//	//if err != nil {
//		//	//	panic(err)
//		//	//}
//		//	//t1.Execute(w, nil)
//		//	tpl, err := template.ParseFiles("./view/user/login.html")
//		//	if err != nil {
//		//		//log.Fatal(),应用直接退出，不在进行
//		//		log.Fatal(err.Error())
//		//	}
//		//	tpl.ExecuteTemplate(w, "user/login.shtml", nil)
//		//})
//		RegisterView()
//		//启动服务
//		http.ListenAndServe("localhost:8080", nil)
//	}
func RegisterView() {
	//一次解析出全部模板
	tpl, err := template.ParseGlob("view/**/*")
	if nil != err {
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _, v := range tpl.Templates() {
		//
		tplname := v.Name()
		fmt.Println("HandleFunc     " + v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			//
			fmt.Println("parse     " + v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w, tplname, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}

}

//func main() {
//	//绑定请求和处理函数
//	http.HandleFunc("/user/login", ctrl.UserLogin)
//	http.HandleFunc("/user/register", ctrl.UserRegister)
//
//	//1 提供静态资源目录支持
//	//http.Handle("/", http.FileServer(http.Dir(".")))
//
//	//2 指定目录的静态文件
//	http.Handle("/asset/", http.FileServer(http.Dir(".")))
//
//	RegisterView()
//
//	http.ListenAndServe(":8080", nil)
//}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	//http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/chat", ctrl.Chat)
	//1 提供静态资源目录支持
	//http.Handle("/", http.FileServer(http.Dir(".")))

	//2 指定目录的静态文件
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	RegisterView()

	http.ListenAndServe(":80", nil)
}
