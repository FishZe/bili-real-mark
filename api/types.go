package api

type ShortReviewResp struct {
	Code int `json:"code"`
	Data struct {
		List  []ShortReview `json:"list"`
		Next  int64         `json:"next"`
		Total int           `json:"total"`
	} `json:"data"`
	Message string `json:"message"`
}

type ShortReview struct {
	Author struct {
		Avatar string `json:"avatar"`
		Mid    int    `json:"mid"`
		Uname  string `json:"uname"`
		Vip    struct {
			AvatarSubscriptUrl string `json:"avatar_subscript_url"`
			NicknameColor      string `json:"nickname_color"`
			ThemeType          int    `json:"themeType"`
			VipStatus          int    `json:"vipStatus"`
			VipType            int    `json:"vipType"`
		} `json:"vip"`
		VipLabel struct {
			BgColor     string `json:"bg_color"`
			BgStyle     int    `json:"bg_style"`
			BorderColor string `json:"border_color"`
			LabelTheme  string `json:"label_theme"`
			Path        string `json:"path"`
			Text        string `json:"text"`
			TextColor   string `json:"text_color"`
		} `json:"vip_label"`
	} `json:"author"`
	Content  string `json:"content"`
	Ctime    int    `json:"ctime"`
	MediaId  int    `json:"media_id"`
	Mid      int    `json:"mid"`
	Mtime    int    `json:"mtime"`
	Progress string `json:"progress"`
	ReviewId int    `json:"review_id"`
	Score    int    `json:"score"`
	Stat     struct {
		Disliked int `json:"disliked"`
		Liked    int `json:"liked"`
		Likes    int `json:"likes"`
	} `json:"stat"`
}

type LongReviewResp struct {
	Code int `json:"code"`
	Data struct {
		Count  int          `json:"count"`
		Folded int          `json:"folded"`
		List   []LongReview `json:"list"`
		Next   int64        `json:"next"`
		Normal int          `json:"normal"`
		Total  int          `json:"total"`
	} `json:"data"`
	Message string `json:"message"`
}

type LongReview struct {
	ArticleId int `json:"article_id"`
	Author    struct {
		Avatar string `json:"avatar"`
		Mid    int    `json:"mid"`
		Uname  string `json:"uname"`
		Vip    struct {
			AvatarSubscriptUrl string `json:"avatar_subscript_url"`
			NicknameColor      string `json:"nickname_color"`
			ThemeType          int    `json:"themeType"`
			VipStatus          int    `json:"vipStatus"`
			VipType            int    `json:"vipType"`
		} `json:"vip"`
		VipLabel struct {
			BgColor     string `json:"bg_color"`
			BgStyle     int    `json:"bg_style"`
			BorderColor string `json:"border_color"`
			LabelTheme  string `json:"label_theme"`
			Path        string `json:"path"`
			Text        string `json:"text"`
			TextColor   string `json:"text_color"`
		} `json:"vip_label"`
	} `json:"author"`
	Content   string `json:"content"`
	Ctime     int    `json:"ctime"`
	IsOrigin  int    `json:"is_origin"`
	IsSpoiler int    `json:"is_spoiler"`
	MediaId   int    `json:"media_id"`
	Mid       int    `json:"mid"`
	Mtime     int    `json:"mtime"`
	Progress  string `json:"progress,omitempty"`
	ReviewId  int    `json:"review_id"`
	Score     int    `json:"score"`
	Stat      struct {
		Likes int `json:"likes"`
		Reply int `json:"reply"`
	} `json:"stat"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type SearchVideResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Seid           string `json:"seid"`
		Page           int    `json:"page"`
		Pagesize       int    `json:"pagesize"`
		NumResults     int    `json:"numResults"`
		NumPages       int    `json:"numPages"`
		SuggestKeyword string `json:"suggest_keyword"`
		RqtType        string `json:"rqt_type"`
		CostTime       struct {
			ParamsCheck         string `json:"params_check"`
			IsRiskQuery         string `json:"is_risk_query"`
			IllegalHandler      string `json:"illegal_handler"`
			AsResponseFormat    string `json:"as_response_format"`
			AsRequest           string `json:"as_request"`
			SaveCache           string `json:"save_cache"`
			DeserializeResponse string `json:"deserialize_response"`
			AsRequestFormat     string `json:"as_request_format"`
			Total               string `json:"total"`
			MainHandler         string `json:"main_handler"`
		} `json:"cost_time"`
		ExpList struct {
			Field1 bool `json:"6609"`
			Field2 bool `json:"7704"`
			Field3 bool `json:"5508"`
		} `json:"exp_list"`
		EggHit     int        `json:"egg_hit"`
		Result     []VideInfo `json:"result"`
		ShowColumn int        `json:"show_column"`
		InBlackKey int        `json:"in_black_key"`
		InWhiteKey int        `json:"in_white_key"`
	} `json:"data"`
}

type VideInfo struct {
	Type           string        `json:"type"`
	MediaId        int64         `json:"media_id"`
	Title          string        `json:"title"`
	OrgTitle       string        `json:"org_title"`
	MediaType      int           `json:"media_type"`
	Cv             string        `json:"cv"`
	Staff          string        `json:"staff"`
	SeasonId       int           `json:"season_id"`
	IsAvid         bool          `json:"is_avid"`
	HitColumns     []interface{} `json:"hit_columns"`
	HitEpids       string        `json:"hit_epids"`
	SeasonType     int           `json:"season_type"`
	SeasonTypeName string        `json:"season_type_name"`
	SelectionStyle string        `json:"selection_style"`
	EpSize         int           `json:"ep_size"`
	Url            string        `json:"url"`
	ButtonText     string        `json:"button_text"`
	IsFollow       int           `json:"is_follow"`
	IsSelection    int           `json:"is_selection"`
	Eps            []struct {
		Id          int    `json:"id"`
		Cover       string `json:"cover"`
		Title       string `json:"title"`
		Url         string `json:"url"`
		ReleaseDate string `json:"release_date"`
		Badges      []struct {
			Text             string `json:"text"`
			TextColor        string `json:"text_color"`
			TextColorNight   string `json:"text_color_night"`
			BgColor          string `json:"bg_color"`
			BgColorNight     string `json:"bg_color_night"`
			BorderColor      string `json:"border_color"`
			BorderColorNight string `json:"border_color_night"`
			BgStyle          int    `json:"bg_style"`
		} `json:"badges"`
		IndexTitle string `json:"index_title"`
		LongTitle  string `json:"long_title"`
	} `json:"eps"`
	Badges []struct {
		Text             string `json:"text"`
		TextColor        string `json:"text_color"`
		TextColorNight   string `json:"text_color_night"`
		BgColor          string `json:"bg_color"`
		BgColorNight     string `json:"bg_color_night"`
		BorderColor      string `json:"border_color"`
		BorderColorNight string `json:"border_color_night"`
		BgStyle          int    `json:"bg_style"`
	} `json:"badges"`
	Cover         string `json:"cover"`
	Areas         string `json:"areas"`
	Styles        string `json:"styles"`
	GotoUrl       string `json:"goto_url"`
	Desc          string `json:"desc"`
	Pubtime       int    `json:"pubtime"`
	MediaMode     int    `json:"media_mode"`
	FixPubtimeStr string `json:"fix_pubtime_str"`
	MediaScore    struct {
		Score     float64 `json:"score"`
		UserCount int     `json:"user_count"`
	} `json:"media_score"`
	DisplayInfo []struct {
		Text             string `json:"text"`
		TextColor        string `json:"text_color"`
		TextColorNight   string `json:"text_color_night"`
		BgColor          string `json:"bg_color"`
		BgColorNight     string `json:"bg_color_night"`
		BorderColor      string `json:"border_color"`
		BorderColorNight string `json:"border_color_night"`
		BgStyle          int    `json:"bg_style"`
	} `json:"display_info"`
	PgcSeasonId int    `json:"pgc_season_id"`
	Corner      int    `json:"corner"`
	IndexShow   string `json:"index_show"`
}
