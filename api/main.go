package api

import (
	"encoding/json"
	"errors"
	"github.com/microcosm-cc/bluemonday"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var Cookies string
var CodeNot0Error = errors.New("resp code not zero")

// getReq
//
//	@Description: 发送Get请求
//	@param data url.Values 请求参数
//	@param getUrl string 请求地址
//	@param cookies string cookie
//	@return []byte 返回的数据
//	@return string 返回的cookie
//	@return error 错误信息
func getReq(data *url.Values, getUrl string, cookies string) ([]byte, string, error) {
	u, err := url.ParseRequestURI(getUrl)
	if err != nil {
		return nil, cookies, err
	}
	u.RawQuery = data.Encode()
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header = http.Header{
		"accept":     {"application/json, text/plain, */*"},
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"},
		"Cookie":     {cookies},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, cookies, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != 200 {
		return nil, cookies, CodeNot0Error
	}
	newCookies := ""
	if resp.Header.Get("Set-Cookie") != "" {
		for _, v := range resp.Header["Set-Cookie"] {
			newCookies += v + ";"
		}
	} else {
		newCookies = cookies
	}
	s, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, cookies, err
	}
	return s, newCookies, nil
}

// getCookies
//
//	@Description: 获取B站无登录的cookies
func getCookies() {
	getUrl := "https://www.bilibili.com/"
	_, Cookies, _ = getReq(&url.Values{"spm_id_from": {"333.999.0.0"}}, getUrl, "")
}

// getShortReview
//
//	@Description:  获取部分短评
//	@param mediaId int64 媒体id
//	@param cursor int64 游标
//	@return *ShortReviewResp 短评响应
func getShortReview(mediaId int64, cursor int64) *ShortReviewResp {
	getUrl := "https://api.bilibili.com/pgc/review/short/list"
	data := url.Values{}
	data.Set("media_id", strconv.FormatInt(mediaId, 10))
	data.Set("ps", "1000")
	data.Set("sort", "0")
	data.Set("cursor", strconv.FormatInt(cursor, 10))
	if s, _, err := getReq(&data, getUrl, ""); err == nil {
		var rep ShortReviewResp
		if err := json.Unmarshal(s, &rep); err == nil {
			return &rep
		}
	}
	return &ShortReviewResp{}
}

// getLongReview
//
//	@Description: 获取部分长评
//	@param mediaId int64 媒体id
//	@param cursor 	int64 游标
//	@return *LongReviewResp 长评响应
func getLongReview(mediaId int64, cursor int64) *LongReviewResp {
	getUrl := "https://api.bilibili.com/pgc/review/long/list"
	data := url.Values{}
	data.Set("media_id", strconv.FormatInt(mediaId, 10))
	data.Set("sort", "0")
	data.Set("ps", "1000")
	data.Set("cursor", strconv.FormatInt(cursor, 10))
	if s, _, err := getReq(&data, getUrl, ""); err == nil {
		var rep LongReviewResp
		if err := json.Unmarshal(s, &rep); err == nil {
			return &rep
		}
	}
	return &LongReviewResp{}
}

// GetAllLongReview
//
//	@Description: 获取所有长评
//	@param mediaId int64 媒体id
//	@param respChan chan *LongReview 长评通道
func GetAllLongReview(mediaId int64, respChan chan *LongReview) {
	nowCursor := int64(0)
	for {
		resp := getLongReview(mediaId, nowCursor)
		if resp.Code == 0 {
			for _, v := range resp.Data.List {
				respChan <- &v
			}
		}
		if resp.Data.Next == 0 {
			break
		}
		nowCursor = resp.Data.Next
		time.Sleep(time.Millisecond * 100)
	}
	respChan <- &LongReview{Mid: -1}
}

// GetAllShortReview
//
//	@Description: 获取所有短评
//	@param mediaId int64 媒体id
//	@param respChan chan *ShortReview 短评通道
func GetAllShortReview(mediaId int64, respChan chan *ShortReview) {
	nowCursor := int64(0)
	for {
		resp := getShortReview(mediaId, nowCursor)
		if resp.Code == 0 {
			for _, v := range resp.Data.List {
				respChan <- &v
			}
		}
		if resp.Data.Next == 0 {
			break
		}
		nowCursor = resp.Data.Next
		time.Sleep(time.Millisecond * 100)
	}
	respChan <- &ShortReview{Mid: -1}
}

// GetReviewSum
//
//	@Description: 获取评论总数
//	@param mediaId int64 媒体id
//	@return res *ReviewSumResp 评论总数响应
func GetReviewSum(mediaId int64) (res int) {
	r1 := getShortReview(mediaId, 0)
	res += r1.Data.Total
	r2 := getLongReview(mediaId, 0)
	res += r2.Data.Total
	return res
}

// SearchVideo
//
//	@Description: 搜索视频
//	@param keyWord string 关键词
//	@return *VideInfo 视频信息
func SearchVideo(keyWord string) *VideInfo {
	if Cookies == "" {
		getCookies()
	}
	getUrl := "https://api.bilibili.com/x/web-interface/search/type"
	data := url.Values{}
	data.Set("keyword", keyWord)
	for _, v := range []string{"media_ft", "media_bangumi"} {
		data.Set("search_type", v)
		if s, _, err := getReq(&data, getUrl, Cookies); err == nil {
			var rep SearchVideResp
			if err = json.Unmarshal(s, &rep); err == nil {
				if rep.Code == 0 {
					for _, x := range rep.Data.Result {
						x.Title = bluemonday.StripTagsPolicy().Sanitize(x.Title)
						if x.Title == keyWord {
							return &x
						}
					}
				}
			}
		}
	}
	return &VideInfo{}
}
