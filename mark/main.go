package mark

import (
	"Bili-TrueMark/api"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"time"
)

type Mark struct {
	Name       string
	mediaId    int64
	totalSum   int
	shortScore map[int]int
	shortSum   int64
	longScore  map[int]int
	longSum    int64
	bar        *progressbar.ProgressBar
}

// init
//
//	@Description: 初始化
//	@receiver m
func (m *Mark) init() {
	m.longScore = make(map[int]int)
	m.shortScore = make(map[int]int)
	m.bar = getBar(m.totalSum)
}

// getBar
//
//	@Description: 获取进度条
//	@param sum int 总数
//	@return *progressbar.ProgressBar
func getBar(sum int) *progressbar.ProgressBar {
	return progressbar.NewOptions(sum,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetItsString("it/s"),
		progressbar.OptionShowCount(),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("[cyan][reset]获取评论中"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
}

// getShortMark
//
//	@Description: 获取短评
//	@receiver m Mark
//	@param over chan bool 通知结束
func (m *Mark) getShortMark(over chan bool) {
	done := make(chan bool)
	go func() {
		xc := make(chan *api.ShortReview, 1000)
		go api.GetAllShortReview(m.mediaId, xc)
		for {
			select {
			case v := <-xc:
				if v.Mid == -1 {
					done <- true
				}
				if _, ok := m.shortScore[v.Score]; !ok {
					m.shortScore[v.Score] = 1
				} else {
					m.shortScore[v.Score]++
				}
				m.shortSum++
				_ = m.bar.Add(1)
			default:
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	for {
		select {
		case <-done:
			over <- true
			return
		}
	}
}

// getLongMark
//
//	@Description: 获取长评
//	@receiver m Mark
//	@param over chan bool 通知结束
func (m *Mark) getLongMark(over chan bool) {
	done := make(chan bool)
	go func() {
		xc := make(chan *api.LongReview, 1000)
		go api.GetAllLongReview(m.mediaId, xc)
		for {
			select {
			case v := <-xc:
				if v.Mid == -1 {
					done <- true
				}
				if _, ok := m.longScore[v.Score]; !ok {
					m.longScore[v.Score] = 1
				} else {
					m.longScore[v.Score]++
				}
				m.longSum++
				_ = m.bar.Add(1)
			default:
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	for {
		select {
		case <-done:
			over <- true
			return
		}
	}
}

// printMark
//
//	@Description: 打印评分
//	@receiver m Mark
func (m *Mark) printMark() {
	fmt.Printf("\n获取到: 短评：%d 个, 长评 %d 个, 共 %d 个\n", m.shortSum, m.longSum, m.shortSum+m.longSum)
	var short, long int64
	for i, v := range m.shortScore {
		short += int64(v * i)
		fmt.Printf("短评 %d 星: %d 个\n", i, v)
	}
	for i, v := range m.longScore {
		long += int64(v * i)
		fmt.Printf("长评 %d 星: %d 个\n", i, v)
	}
	fmt.Printf("短评平均分 %f \n", float64(short)/float64(m.shortSum))
	fmt.Printf("长评平均分 %f \n", float64(long)/float64(m.longSum))
	fmt.Printf("总体平均分 %f \n", float64(short+long)/float64(m.shortSum+m.longSum))
}

// getAllMark
//
//	@Description: 获取所有评分
//	@receiver m Mark
func (m *Mark) getAllMark() {
	over := make(chan bool)
	go m.getShortMark(over)
	go m.getLongMark(over)
	sum := 0
	for {
		select {
		case <-over:
			sum++
			if sum == 2 {
				m.printMark()
				return
			}
		}
	}
}

// Start
//
//	@Description: 开始
//	@receiver m
func (m *Mark) Start() {
	videoInfo := api.SearchVideo(m.Name)
	if videoInfo.MediaId == 0 {
		fmt.Printf("未找到视频：%s\n", m.Name)
		return
	}
	fmt.Printf("%v (%v)[%d] 评分: %f\n\n", videoInfo.Title, videoInfo.SeasonTypeName, videoInfo.MediaId, videoInfo.MediaScore.Score)
	m.mediaId = videoInfo.MediaId
	m.totalSum = api.GetReviewSum(m.mediaId)
	m.init()
	m.getAllMark()
}

// Cmd
//
//	@Description: 命令
func Cmd() {
	m := Mark{}
	fmt.Println("请输入影视/番剧名称：")
	_, _ = fmt.Scanln(&m.Name)
	m.Start()
}
