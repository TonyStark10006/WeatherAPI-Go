package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	mysql1 "weatherAPI/utils/mysql"
	redis "weatherAPI/utils/redis"
)

func GetWeatherMsg(w http.ResponseWriter, r *http.Request) {
	val, _ := redis.GetAValue("good")
	// mysql.Table().Get()
	fmt.Println(val)

	// 重写cookie结构
	c1 := http.Cookie{
		Name:     "cookieByGo",
		Value:    "this is a cookie for broswer",
		HttpOnly: true,
	}
	// r.ParseForm() //解析表单数据
	// r.Form["city"] // 返回切片
	// city := r.FormValue("city")
	// h := r.Header // 获取所有首部字段信息，返回map
	// h := r.Header.Get("Accept-Encoding") // 通过Get方法获取首部字段信息
	// fmt.Fprint(w, h)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Set-Cookie", c1.String()) //等价于下面一行，多个cookie可以用Add()方法替换Set方法
	http.SetCookie(w, &c1)
	good := WeatherMsg{
		StatusCode: 200,
		UpdateAt:   "2018/06/19",
		Data:       []string{"good", "haha"},
	}
	// good := []string{"good", "haha"}
	// output, _ := json.MarshalIndent(&good, "", "\t\t")
	output, _ := json.Marshal(good)
	w.Write(output)
	fmt.Println(r.Cookie("cookieByGo"))
}

type WeatherMsg struct {
	StatusCode int      `json:"status_code"`
	UpdateAt   string   `json:"update_at"`
	Data       []string `json:"data"`
}

func getWeatherMsgByName(city string) (string, error) {
	if city == "" {
		return "输入城市为空", nil
	}

	good := &mysql1.STATEMENT{}
	// good.Select([]string{"id", "username"}).
	good.Insert(map[string]string{"lihaile": "gaga", "hehe": "xixi"}).
		Table("nima").
		Where(map[string]string{"wo": "detian"}).
		Get()
	//fmt.Println(good.FullStatement)

	return "找不到这个城市的天气信息啊", nil
}

// func main() {
// 	good := &STATEMENT{}
// good.Select([]string{"id", "username"}).
// good.Insert(map[string]string{"lihaile": "gaga", "hehe": "xixi"}).
// 	Table("nima").
// 	Where(map[string]string{"wo": "detian"}).
// 	Get()
//good.gather()
//fmt.Println(good.fullStatement)
// rows, err := DB.Query("select id, user_name, email from user")
// type POST struct {
// 	ID       int
// 	username string
// 	email    string
// }
// post := POST{}
// var posts []POST
// // var post map[string]string
// for rows.Next() {
// 	err = rows.Scan(&post.ID, &post.username, &post.email)
// 	posts = append(posts, post)
// }
// fmt.Println(posts)
// for key, value := range posts {
// 	fmt.Println(string(key) + ":" + value.email + "\n")
// }
// if err != nil {
// 	fmt.Println(err)
// }
// rows.Close()
// }
