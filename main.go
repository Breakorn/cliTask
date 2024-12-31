// 从控制台读取输入:
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Command struct{}
type Task struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdaredAt   string `json:"updaredAt"`
}

func (com Command) Add(s string) {
	// 获取当前的本地时间
	loc, _ := time.LoadLocation("Asia/Shanghai")
	currentTime := time.Now().In(loc).Format("2006-01-02 15:04")
	var data []Task
	file, err := GetJson("work.json")
	if err != nil {
		d := []Task{}
		d = append(d, Task{
			Id:          "1",
			Description: s,
			Status:      "todo",
			CreatedAt:   currentTime,
			UpdaredAt:   currentTime})
		SetJson("work.json", d)
		return
	}
	data = file
	item := data[len(data)-1]

	id, err := strconv.Atoi(item.Id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	id = id + 1
	taskid := strconv.Itoa(id)

	task := Task{
		Id:          taskid,
		Description: s,
		Status:      "todo",
		CreatedAt:   currentTime,
		UpdaredAt:   currentTime,
	}
	data = append(data, task)
	SetJson("work.json", data)

}
func (com Command) List(s string) {
	var data []Task
	file, err := GetJson("work.json")
	if err != nil {
		println("粗我")
	}
	data = file
	var str = "id 状态  创建时间            更新时间             描述\n"

	switch s {
	case "done":
		for i := 0; i < len(data); i++ {
			if data[i].Status == "done" {
				str += data[i].Id + "  " + data[i].Status + " " + data[i].CreatedAt + " " + data[i].UpdaredAt + " " + data[i].Description + "\n"
			}
		}
		println(str)
	case "todo":
		for i := 0; i < len(data); i++ {
			if data[i].Status == "todo" {
				str += data[i].Id + "  " + data[i].Status + " " + data[i].CreatedAt + " " + data[i].UpdaredAt + " " + data[i].Description + "\n"
			}
		}
		println(str)
	case "doing":
		for i := 0; i < len(data); i++ {
			if data[i].Status == "doing" {

				str += data[i].Id + "  " + data[i].Status + " " + data[i].CreatedAt + " " + data[i].UpdaredAt + " " + data[i].Description + "\n"
			}
		}
		println(str)
	default:
		for i := 0; i < len(data); i++ {
			str += data[i].Id + "  " + data[i].Status + " " + data[i].CreatedAt + " " + data[i].UpdaredAt + " " + data[i].Description + "\n"
		}
		println(str)
	}

}
func (com Command) Update(s string, s2 string) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	currentTime := time.Now().In(loc).Format("2006-01-02 15:04")
	var data []Task
	file, err := GetJson("work.json")
	if err != nil {
		println("获取json失败")
		return
	}
	data = file

	index := -1
	for i := 0; i < len(data); i++ {
		if s == data[i].Id {

			data[i].Description = s2
			data[i].UpdaredAt = currentTime
			index = i
			break
		}
	}
	if index == -1 {
		println("不存在该任务")
		return
	}
	SetJson("work.json", data)

}
func (com Command) MarkDone(s string) {
	mark("done", s)

}
func (com Command) MarkDoing(s string) {
	mark("doing", s)
}

func (com Command) Delete(s string) {
	var data []Task
	file, err := GetJson("work.json")
	if err != nil {
		println("获取json失败")
		return
	}
	data = file

	index := -1
	for i := 0; i < len(data); i++ {
		if data[i].Id == s {
			index = i
			break
		}
	}

	if index != -1 {
		data = append(data[:index], data[index+1:]...)
		SetJson("work.json", data)
	}

}
func GetJson(name string) ([]Task, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	// 创建一个 Task 类型的变量
	var list []Task

	// 使用 json.Unmarshal 解析 JSON 数据
	err = json.Unmarshal(data, &list)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}
	return list, nil
}
func SetJson(name string, data []Task) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	err1 := os.WriteFile(name, jsonData, 0644)
	if err1 != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	println("写入成功！")
}
func mark(name string, s string) {
	var data []Task
	file, err := GetJson("work.json")
	if err != nil {
		println("不存在")
		return
	}
	data = file
	index := -1
	for i := 0; i < len(data); i++ {
		if s == data[i].Id {
			loc, _ := time.LoadLocation("Asia/Shanghai")
			currentTime := time.Now().In(loc).Format("2006-01-02 15:04")
			data[i].Status = name
			data[i].UpdaredAt = currentTime
			index = i
			break
		}
	}
	if index == -1 {
		println("不存在该任务")
		return
	}

	SetJson("work.json", data)
}

// func
func main() {
	var classCom = Command{}
	var (
		com, arg, arg2 string
	)
	fmt.Scanln(&com, &arg, &arg2)
	value := reflect.ValueOf(classCom)
	method := value.MethodByName(com)
	if method.IsValid() {
		// 准备传递给方法的参数
		params := []reflect.Value{
			reflect.ValueOf(arg), // 传递的参数
		}
		if com == "Update" {
			params1 := []reflect.Value{
				reflect.ValueOf(arg), // 传递的参数
				reflect.ValueOf(arg2),
			}
			method.Call(params1)
			return
		}
		method.Call(params) // 调用方法，传递空参数
	} else {
		fmt.Println("未发现该指令", com, arg)
	}

}
