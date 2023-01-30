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
	shortScore []Review
	shortSum   int64
	longScore  []Review
	longSum    int64
	bar        *progressbar.ProgressBar
}

type Review struct {
	Mid     *int64
	Name    *string
	Time    *string
	Score   *int
	Content *string
}

// init
//
//	@Description: 初始化
//	@receiver m
func (m *Mark) init() {
	m.longScore = make([]Review, 0)
	m.shortScore = make([]Review, 0)
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
		xc := make(chan api.ShortReview, 1000)
		go api.GetAllShortReview(m.mediaId, xc)
		for {
			select {
			case v := <-xc:
				if v.Mid == -1 {
					done <- true
				}
				t := time.Unix(v.Ctime, 0).Format("2006-01-02 15:04:05")
				nowReview := Review{
					Mid:     &v.Mid,
					Name:    &v.Author.Uname,
					Time:    &t,
					Score:   &v.Score,
					Content: &v.Content,
				}
				m.shortScore = append(m.shortScore, nowReview)
				m.shortSum++
				_ = m.bar.Add(1)
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
		xc := make(chan api.LongReview, 1000)
		go api.GetAllLongReview(m.mediaId, xc)
		for {
			select {
			case v := <-xc:
				if v.Mid == -1 {
					done <- true
				}
				t := time.Unix(v.Ctime, 0).Format("2006-01-02 15:04:05")
				longReview := Review{
					Mid:     &v.Mid,
					Name:    &v.Author.Uname,
					Time:    &t,
					Score:   &v.Score,
					Content: &v.Content,
				}
				m.longScore = append(m.longScore, longReview)
				m.longSum++
				_ = m.bar.Add(1)
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

func (m *Mark) Save2Excel() {
	key := []string{"mid", "name", "time", "score", "content"}
	keys := map[string][]string{}
	keys["short"] = key
	keys["long"] = key
	short := make(map[string][]string, 0)
	long := make(map[string][]string, 0)
	for _, v := range key {
		short[v] = make([]string, 0)
		long[v] = make([]string, 0)
	}
	for _, v := range m.shortScore {
		short["mid"] = append(short["mid"], fmt.Sprintf("%d", *v.Mid))
		short["name"] = append(short["name"], *v.Name)
		short["time"] = append(short["time"], *v.Time)
		short["score"] = append(short["score"], fmt.Sprintf("%d", *v.Score))
		short["content"] = append(short["content"], *v.Content)
	}
	for _, v := range m.longScore {
		long["mid"] = append(long["mid"], fmt.Sprintf("%d", *v.Mid))
		long["name"] = append(long["name"], *v.Name)
		long["time"] = append(long["time"], *v.Time)
		long["score"] = append(long["score"], fmt.Sprintf("%d", *v.Score))
		long["content"] = append(long["content"], *v.Content)
	}
	review := make(map[string]map[string][]string, 0)
	review["short"] = short
	review["long"] = long
	err := Write2Excel(m.Name+".xlsx", review, keys)
	if err != nil {
		fmt.Printf("保存失败: %s", err)
	} else {
		fmt.Printf("保存成功: %s.xlsx", m.Name)
	}
}

// printMark
//
//	@Description: 打印评分
//	@receiver m Mark
func (m *Mark) printMark() {
	fmt.Printf("\n获取到: 短评：%d 个, 长评 %d 个, 共 %d 个\n", m.shortSum, m.longSum, m.shortSum+m.longSum)
	shortScore, longScore := make(map[int]int, 10), make(map[int]int, 10)
	var short, long int64
	for _, v := range m.shortScore {
		shortScore[*v.Score]++
		short += int64(*v.Score)
	}
	for _, v := range m.longScore {
		longScore[*v.Score]++
		long += int64(*v.Score)
	}
	for i := 10; i > 0; i -= 2 {
		if _, ok := shortScore[i]; ok {
			fmt.Printf("短评 %d 星: %d 个\n", i, shortScore[i])
		}
	}
	for i := 10; i > 0; i -= 2 {
		if _, ok := longScore[i]; ok {
			fmt.Printf("长评 %d 星: %d 个\n", i, longScore[i])
		}
	}
	fmt.Printf("短评平均分 %f \n", float64(short)/float64(m.shortSum))
	fmt.Printf("长评平均分 %f \n", float64(long)/float64(m.longSum))
	fmt.Printf("总体平均分 %f \n", float64(short+long)/float64(m.shortSum+m.longSum))
	m.Save2Excel()
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
