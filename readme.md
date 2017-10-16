### GOLANG datetime包
#### 提供类似php的strtotime(),time(),date()格式化等方法
```
data := make(map[string]interface{});

time := datetime.Time();
data["time"] = time;
data["date"] = datetime.Date("Y-m-d H:i:s",time + 100);
_strtotime  := datetime.Strtotime("2017-10-16 23:45:34");
data["strtotime1"] = datetime.Date("Y-m-d H:i:s",_strtotime);
_strtotime = datetime.Strtotime("-1days");
data["strtotime2"] = datetime.Date("Y年m月d日 H时i分s秒",_strtotime);
data["strtotime3"] = datetime.Date("Y/m/d H:i",_strtotime);
```