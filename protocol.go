package mop

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"net/http/cookiejar"
	"errors"
)

type JSONuserInfo struct {
	UserBaseInfo struct {
		HomeCity  string `json:"home_city"`
		Basicinfo []struct {
			Title   string `json:"title"`
			Context string `json:"context"`
		} `json:"basicinfo"`
		CntUserTime string `json:"cnt_user_time"`
		UserID      string `json:"user_id"`
	} `json:"userBaseInfo"`
	Success bool `json:"success"`
}
type JSONbusinessInfo struct {
	BusinessList struct {
		Firstlevelsize int `json:"firstlevelsize"`
		Firstlevel     []struct {
			Secondlevel []struct {
				Secondvalue    string `json:"secondvalue"`
				Thirdlevelsize int    `json:"thirdlevelsize"`
				Thirdlevel     []struct {
					Thirdcontext string `json:"thirdcontext"`
					Thirdtitle   string `json:"thirdtitle"`
				} `json:"thirdlevel"`
			} `json:"secondlevel"`
			Secondlevelsize int    `json:"secondlevelsize"`
			Firstvalue      string `json:"firstvalue"`
		} `json:"firstlevel"`
	} `json:"businessList"`
	Success bool `json:"success"`
}
type JSONconsumeInfo struct {
	ConsumeList struct {
		Firstlevelsize int `json:"firstlevelsize"`
		Firstlevel     []struct {
			Secondlevel []struct {
				Secondvalue    string `json:"secondvalue"`
				Thirdlevelsize int    `json:"thirdlevelsize"`
				Thirdlevel     []struct {
					Thirdcontext string `json:"thirdcontext"`
					Thirdtitle   string `json:"thirdtitle"`
				} `json:"thirdlevel"`
			} `json:"secondlevel"`
			Secondlevelsize int    `json:"secondlevelsize"`
			Firstvalue      string `json:"firstvalue"`
		} `json:"firstlevel"`
	} `json:"consumeList"`
	Success bool `json:"success"`
}
type JSONuserBaseInfo struct {
	BillFlag     bool `json:"bill_flag"`
	IndivRecommendInfoList []struct {
		LastRecomAction      string `json:"last_recom_action"`
		SaleName             string `json:"sale_name"`
		IntroduceInfo        string `json:"introduce_info"`
		LastRecomChannel     string `json:"last_recom_channel"`
		InureTime            string `json:"inure_time"`
		RecomTotalCnt        string `json:"recom_total_cnt"`
		ExpireTime           string `json:"expire_time"`
		RewardRegNo          string `json:"reward_reg_no"`
		RecommendInfo        string `json:"recommend_info"`
		RewardAmount         string `json:"reward_amount"`
		RecommendSubType     string `json:"recommend_sub_type"`
		LastRecomTime        string `json:"last_recom_time"`
		LastRecomChannelName string `json:"last_recom_channel_name"`
		ExistRecomInfo       string `json:"exist_recom_info"`
		AcceptType           string `json:"accept_type"`
		RecomSuccCnt         string `json:"recom_succ_cnt"`
		RecommendType        string `json:"recommend_type"`
		SaleID               string `json:"sale_id"`
	} `json:"indiv_recommend_info_list"`
	UnifyRecommendInfoList []struct {
		LastRecomAction      string        `json:"last_recom_action"`
		SaleName             string        `json:"sale_name"`
		LastRecomChannel     string        `json:"last_recom_channel"`
		RecommendBusi        string        `json:"recommend_busi"`
		GradeName            string        `json:"grade_name"`
		GradeEndDate         string        `json:"grade_end_date"`
		ProgID               string        `json:"prog_id"`
		BusiInfo []struct {
			BusinessName string `json:"business_name"`
			BusinessID   string `json:"business_id"`
			BusinessType string `json:"business_type"`
		} `json:"busi_info"`
		ProgName             string        `json:"prog_name"`
		RecommendInfo        string        `json:"recommend_info"`
		BusinessName         string        `json:"business_name"`
		GradeStartDate       string        `json:"grade_start_date"`
		BusiExt2             string        `json:"busi_ext2"`
		BusiExt1             string        `json:"busi_ext1"`
		LastRecomTime        string        `json:"last_recom_time"`
		LastRecomChannelName string        `json:"last_recom_channel_name"`
		Priority             string        `json:"priority"`
		ExistRecomInfo       string        `json:"exist_recom_info"`
		BusinessID           string        `json:"business_id"`
		RecommendType        string        `json:"recommend_type"`
		GradeID              string        `json:"grade_id"`
		BusinessType         string        `json:"business_type"`
		BusiExt3             string        `json:"busi_ext3"`
		SaleID               string        `json:"sale_id"`
	} `json:"unify_recommend_info_list"`
	UserBaseInfo struct {
		IsBoundTerminal        string `json:"is_bound_terminal"`
		Is4GBaseProduct        string `json:"is_4g_base_product"`
		RecommendOptProduct    string `json:"recommend_opt_product"`
		Is4GClient             string `json:"is_4g_client"`
		RecommendBaseProduct   string `json:"recommend_base_product"`
		IsVotleFunction        string `json:"is_votle_function"`
		IsFamilyUser           string `json:"is_family_user"`
		BaseProduct            string `json:"base_product"`
		RecommendHotBusiness   string `json:"recommend_hot_business"`
		BaseProductID          string `json:"base_product_id"`
		TerminalOs             string `json:"terminal_os"`
		Is4GUser               string `json:"is_4g_user"`
		ChangeCriticalPoint    string `json:"change_critical_point"`
		IsHvalueUser           string `json:"is_hvalue_user"`
		TerminalPrice          string `json:"terminal_price"`
		HomeCity               string `json:"home_city"`
		CurrMonthAddFlow       string `json:"curr_month_add_flow"`
		IsBroadBand            string `json:"is_broad_band"`
		UrgencyLevel           string `json:"urgency_level"`
		CreditLevel            string `json:"credit_level"`
		RecommendBaseProductID string `json:"recommend_base_product_id"`
		IsUsimUser             string `json:"is_usim_user"`
		IsTdUser               string `json:"is_td_user"`
		LastMonthUsedFlow      string `json:"last_month_used_flow"`
		AggregationMarket      string `json:"aggregation_market"`
		CurrMonthValidFlow     string `json:"curr_month_valid_flow"`
		UsimChangeTime         string `json:"usim_change_time"`
		IsGroupUser            string `json:"is_group_user"`
		IsSphoneUser           string `json:"is_sphone_user"`
		CurrMonthTotalFlow     string `json:"curr_month_total_flow"`
		ActivityName           string `json:"activity_name"`
		ExpireTime             string `json:"expire_time"`
		TerminalType           string `json:"terminal_type"`
		NetAge                 string `json:"net_age"`
		IsTargetValueUser      string `json:"is_target_value_user"`
		IsRealnameUser         string `json:"is_realname_user"`
		Is4GSims               string `json:"is_4g_sims"`
		UserName               string `json:"user_name"`
		IsDoubleCardSlot       string `json:"is_double_card_slot"`
		ChgTerminalPttUser     string `json:"chg_terminal_ptt_user"`
		RecentConsumeStatus    string `json:"recent_consume_status"`
		HomeNetPttUser         string `json:"home_net_ptt_user"`
		UserBirthday           string `json:"user_birthday"`
		IsVotleTerminal        string `json:"is_votle_terminal"`
		TerminalChangeTime     string `json:"terminal_change_time"`
	} `json:"user_base_info"`
	IsProvinceSale bool `json:"isProvinceSale"`
	Success        bool `json:"success"`
	BillInfoList   []struct {
		ConsumeType string `json:"consume_type"`
		UsedValue   string `json:"used_value"`
		TotalValue  string `json:"total_value"`
	} `json:"bill_info_list"`
}

type Protocol struct {
	client http.Client
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func setDefaultHeader(req *http.Request) {
	req.Header.Set("client_mac", "00:00:00:00:00:00")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

func (m*Protocol) Init(vc string) (err error) {
	cookieJar, _ := cookiejar.New(nil)
	m.client = http.Client{
		Jar:       cookieJar,
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxIdleConnsPerHost: 1<<32 - 1,
			IdleConnTimeout:     0,
			DisableCompression:  false,
			DisableKeepAlives:   false,
			Proxy:               nil,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},

	}
	url := "http://112.5.185.82:8881/MBossWeb/mbop/index_hidden.jsp?vc={vc}&ptid=770489400020&opType=0"
	url = strings.Replace(url, "{vc}", vc, 1)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	setDefaultHeader(req)
	resp, err := m.client.Do(req)
	cookie := resp.Header.Get("Set-Cookie")
	if strings.Index(cookie, "JSESSIONID") < 0 {
		return errors.New("init failed")
	}
	return
}

func (m*Protocol) Query(mobile string, inter interface{}) (err error) {
	var method string
	if _, ok := (inter).(*JSONuserInfo); ok {
		method = "queryUserInfo"
	} else if _, ok := (inter).(*JSONconsumeInfo); ok {
		method = "queryConsumeInfo"
	} else if _, ok := (inter).(*JSONbusinessInfo); ok {
		method = "queryBusinessInfo"
	} else if _, ok := (inter).(*JSONuserBaseInfo); ok {
		method = "QueryUserBaseInfo"
	} else {
		log.Fatal("no this method.")
		return
	}

	url := "http://112.5.185.82:8881/MBossWeb/bmaccept/4assambleQueryMgr.do?method=" + method
	body := "msisdn=" + mobile
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	setDefaultHeader(req)
	resp, err := m.client.Do(req)
	if err != nil {
		return
	}
	str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	str, err = GbkToUtf8(str)
	if err != nil {
		return
	}
	err = json.Unmarshal(str, inter)
	return
}

