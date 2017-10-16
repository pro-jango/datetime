// @package datetime
// @author caiyujiang
// @version 2017-10-16

// 提供类似php的strtotime(),time(),date()格式化等方法
// making http request.
package datetime

import (
	"errors"
	"strings"
	"time"
	"strconv"
)

/**
 * Date 格式化字符串
 * timestamp 时间戳
 * 例1：Y-m-d 返回2017-08-24
 * 例2：y年m月d日 返回 17年08月24日
 * 例3：H:i:s 返回 17:04:57
 */
func Date(date string,timestamp int64) string{
	//创建一个多对替换的规则
	replaceRule := strings.NewReplacer("Y","2006","y","06","m","01","d","02","H","15","i","04","s","05");
	date = replaceRule.Replace(date);

	_timeModel := time.Unix(timestamp, 0);
	
	return _timeModel.Format(date);
}


/**
 * 获取当前系统时间戳
 */
func Time() int64{
	_timeNow := time.Now();
	return _timeNow.Unix();
}

/**
 * 通过字符串获取时间戳
 * return int64 时间戳
 * return err 错误信息
 * 用例：newtimes,_ := getTime("-1days")
 * 例2：newtimes,_ := getTime("+12hours")
 * 例3：newtimes,_ := getTime("+1years")
 */
func Strtotime(str string) int64{
	str = strings.ToLower(str);

	arrTmpMap := make(map[string]int64);
	arrTmpMap["seconds"] = 1;
	arrTmpMap["minutes"] = 60;
	arrTmpMap["hours"] = 3600;
	arrTmpMap["days"] = 86400;

	_timeNow := time.Now();

	//解析*seconds,*minutes,*hours,*days
	for key := range arrTmpMap{
		if (strings.HasSuffix(str,key)){
			number,err := getFormatInt(str,key);
			if(err != nil){
				return 0;
			}
			number = _timeNow.Unix() + number * arrTmpMap[key];
			return number;
		}
	}
	//解析*months
	if strings.HasSuffix(str,"months") {
		numberMonths,err := getFormatInt(str,"months");
		if(err != nil){
			return 0;
		}
		_timeNow = _timeNow.AddDate(0,int(numberMonths),0);
		return _timeNow.Unix();
	}

	//解析*years
	if strings.HasSuffix(str,"years") {
		numberYears,err := getFormatInt(str,"years");
		if(err != nil){
			return 0;
		}
		_timeNow = _timeNow.AddDate(int(numberYears),0,0);
		return _timeNow.Unix();
	}

	//解析 xxxx-xx-xx和xxxx-xx-xx xx:xx:xx格式的字符串
	strLen := len(str);
	if(strLen == 10){
		str = str + " 00:00:00";
	}
	strLen = len(str);
	if(strLen == 19){
		//不使用Parse解析是因为有8个小时的时差
		//_timeModel,_ := time.Parse("2006-01-02 15:04:05",str); 
		_timeModel,_ := time.ParseInLocation("2006-01-02 15:04:05",str,time.Local);
		return _timeModel.Unix();
	}
	return 0;
}

func getFormatInt(str string,key string) (int64,error){
	strNumber := strings.Trim(str,key);
	strNumber = strings.TrimSpace(strNumber);
	number,err := strconv.ParseInt(strNumber,10,0);
	if(err != nil){
		return 0,errors.New("字符类型错误");
	}else{
		return number,nil;
	}
}