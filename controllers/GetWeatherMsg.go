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
	mysql.Table().Get()
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

	DB := &mysql1.STATEMENT{}

	return "找不到这个城市的天气信息啊", nil
}
