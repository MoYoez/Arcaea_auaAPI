package aua // Package aua

import (
	"io"
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
	replyByte, err := http.NewRequest("GET", workurl, nil)
	replyByte.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	replyByte.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(replyByte)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	replyBack, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	reply = string(replyBack)
	return reply, err
}
