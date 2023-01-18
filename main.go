package main

import (
	"github.com/FloatTech/floatbox/web"
	"github.com/fumiama/jieba/util/helper"
	"net/http"
)

// GetUserInfo Url 为完整的 ArcaeaAPI 请求，token 为此项目需要用到的 Auth，id是需要查询使用的id。
func GetUserInfo(url string, token string, arcaeaid string) (reply string, err error) {
	reply, err = DrawRequestArc(url+"/botarcapi/user/info?user="+arcaeaid+"&recent=1&withsonginfo=true", token)
	if err != nil {
		return "", err
	}
	return reply, err
}

// Best30 可以参考GetUserInfo
func Best30(url string, token string, arcaeaid string) (reply string, err error) {
	reply, err = DrawRequestArc(url+"/botarcapi/user/best30?user="+arcaeaid+"&withrecent=false&overflow=10&withsonginfo=true", token)
	if err != nil {
		return "", err
	}
	return reply, err
}

func GetUserBest(url string, token string, arcaeaid string, songname string, difficuity string) (reply string, err error) {
	reply, err = DrawRequestArc(url+"/botarcapi/user/best?user="+arcaeaid+"&songname="+songname+"&difficulty="+difficuity+"&withsonginfo=true", token)
	if err != nil {
		return "", err
	}
	return reply, err
}

func GetSongRandom(url string, token string, start string, end string) (reply string, err error) {
	reply, err = DrawRequestArc(url+"/botarcapi/song/random?start="+start+"&end="+end+"&withsonginfo=true", token)
	if err != nil {
		return "", err
	}
	return reply, err
}

func GetSongInfo(url string, token string, songname string) (reply string, err error) {
	reply, err = DrawRequestArc(url+"/botarcapi/song/info?songname="+songname, token)
	if err != nil {
		return "", err
	}
	return reply, err
}

func GetSongPreview(url string, token string, songname string, difficuity string) (reply string, err error) {
	reply, err = DrawRequestArc(url+"/botarcapi/assets/preview?songname="+songname+"&difficulty="+difficuity, token)
	if err != nil {
		return "", err
	}
	return reply, err
}

func DrawRequestArc(workurl string, token string) (reply string, err error) {
	replyByte, err := web.RequestDataWithHeaders(web.NewDefaultClient(), workurl, "GET", func(r *http.Request) error {
		r.Header.Set("accept-language", "zh,zh-CN;q=0.9,zh-HK;q=0.8,zh-TW;q=0.7,ja;q=0.6,en;q=0.5,en-GB;q=0.4,en-US;q=0.3")
		r.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
		r.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko)  Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76")
		r.Header.Set("Authorization", "Bearer "+token)
		return nil
	}, nil)
	if err != nil {
		return "", err
	}
	reply = helper.BytesToString(replyByte)
	return
}
