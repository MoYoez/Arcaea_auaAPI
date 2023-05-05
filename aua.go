// Package aua
package aua

import (
	"bytes"
	"github.com/tidwall/gjson"
	"image"
	"io"
	"net/http"
)

// GetUserInfo Url 为完整的 ArcaeaAPI 请求，token 为此项目需要用到的 Auth，id是需要查询使用的id。
func GetUserInfo(url string, token string, arcaeaid string) (reply []byte, err error) {
	reply, err = DrawRequestArc(url+"/arcapi/user/info?user_name="+arcaeaid+"&recent=1&with_song_info=true", token)
	if err != nil {
		return nil, err
	}
	return reply, err
}

// GetUserBest 获取用户最新成绩
func GetUserBest(url string, token string, arcaeaid string, songname string, difficuity string) (reply []byte, err error) {
	reply, err = DrawRequestArc(url+"/arcapi/user/best?user_name="+arcaeaid+"&song_name="+songname+"&difficulty="+difficuity+"&with_song_info=true", token)
	if err != nil {
		return nil, err
	}
	return reply, err
}

// GetSongRandom 随机一首曲子（
func GetSongRandom(url string, token string, start string, end string) (reply []byte, err error) {
	reply, err = DrawRequestArc(url+"/arcapi/song/random?start="+start+"&end="+end+"&with_song_info=true", token)
	if err != nil {
		return nil, err
	}
	return reply, err
}

// GetSongInfo 获得歌曲信息
func GetSongInfo(url string, token string, songname string) (reply []byte, err error) {
	reply, err = DrawRequestArc(url+"/arcapi/song/info?song_name="+songname, token)
	if err != nil {
		return nil, err
	}
	return reply, err
}

// GetSongPreview 返回谱面预览图，需要歌曲名字和难度，如果无结果则说明谱面没有（
func GetSongPreview(url string, token string, songname string, difficuity string) (images image.Image, err error) {
	reply, err := DrawRequestArc(url+"/arcapi/assets/preview?song_name="+songname+"&difficulty="+difficuity, token)
	if err != nil {
		return nil, err
	}
	images, _, err = image.Decode(bytes.NewReader(reply))
	if err != nil {
		panic(err)
	}
	return images, err
}

// GetSessionQuery Get Session (query b30,need 5 people)
func GetSessionQuery(url string, token string, id string) (sessionkey string, info string) {
	getSession, err := DrawRequestArc(url+"/arcapi/user/bests/session?user="+id, token)
	if err != nil {
		return "", ""
	}
	sessionInfo := gjson.Get(string(getSession), "content.session_info").String()
	sessionStatus := gjson.Get(string(getSession), "status").String()
	if sessionStatus == "0" {
		return sessionInfo, ""
	} else {
		sessionMsg := gjson.Get(string(getSession), "message").String()
		return sessionInfo, sessionMsg
	}
}

// GetB30BySession Get B30 By Session (wait in line mode.)
func GetB30BySession(url string, token string, sessionkey string) (reply []byte, msg string) {
	reply, _ = DrawRequestArc(url+"/arcapi/user/bests/result?session_info="+sessionkey+"&overflow=10&with_recent=false&with_song_info=true", token)
	getStatus := gjson.Get(string(reply), "status").String()
	if getStatus != "0" {
		getMsg := gjson.Get(string(reply), "message").String()
		return nil, getMsg
	}
	return reply, ""
}

//

// DrawRequestArc 发送请求结构体
func DrawRequestArc(workurl string, token string) (reply []byte, err error) {
	replyByte, err := http.NewRequest("GET", workurl, nil)
	replyByte.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	replyByte.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return nil, err
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
	return replyBack, err
}
