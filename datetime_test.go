// @package datetime
// @author caiyujiang
// @version 2017-10-16

// 单元测试用例
// 使用方法：当前目录下执行：go test -v
package datetime

import (
    "testing"
)

func Test_Date_1(t * testing.T){
	if date1:= Date("Y-m-d",1508126777);date1 == "2017-10-16" {
		t.Log(date1);
		t.Log("date1测试通过了");
	}else{
		t.Error("date1未通过");
	}
}

func Test_Date_2(t * testing.T){
	if date1:= Date("Y/m/d H:i:s",1508126777);date1 == "2017/10/16 12:06:17" {
		t.Log(date1);
		t.Log("date2测试通过了");
	}else{
		t.Error("date2未通过");
	}
}

func Test_Date_3(t * testing.T){
	if date1:= Date("Y年m月d日 H时i分s秒",1508126777);date1 == "2017年10月16日 12时06分17秒" {
		t.Log(date1);
		t.Log("date3测试通过了");
	}else{
		t.Error("date3未通过");
	}
}

func Test_Time(t * testing.T){
	time := Time();
	if time >0 {
		t.Log("Time测试通过了");
	}else{
		t.Error("Time未通过");
	}
}

func Test_Strtotime_1(t * testing.T){
	time := Strtotime("2017-10-16 12:06:17");
	t.Log(time);
	if time == 1508126777 {
		t.Log("Strtotime_1测试通过了");
	}else{
		t.Error("Strtotime_1未通过");
	}
}

func Test_Strtotime_2(t * testing.T){
	time := Strtotime("+1days");
	now := Time();
	if time > now {
		t.Log(Date("Y-m-d H:i:s",time));
		t.Log("Strtotime_2测试通过了");
	}else{
		t.Error("Strtotime_2未通过");
	}
}