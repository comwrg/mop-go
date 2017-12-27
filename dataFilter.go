package mop

import (
	"regexp"
	"strconv"
)

type UserInfo struct {
	name, attribution, startTime, basePackage, userStatus, billingType, groupName, bundlingTime string
}

type ConsumeInfo struct {
	// not use 0th value
	callsConsume [13]string
	flowConsume  [13]string
}

type UserBaseInfo struct {
	recommendInfo, terminalType, isBoundTerminal, terminalChangeTime, isBroadBand, isRealNameUser, isSphoneUser,
	isUsimUser, is4gClient, is4gBaseProduct, isVotleTerminal, isVotleFunction, isDoubleCardSlot string
}

func NewConsumeInfo() (ci ConsumeInfo){
	for i, _ := range ci.callsConsume {
		ci.callsConsume[i] = "0"
	}
	for i, _ := range ci.flowConsume {
		ci.flowConsume[i] = "0"
	}
	return
}

func FilterUserInfo(info * JSONuserInfo) (ui UserInfo) {
	ui.name         = info.UserBaseInfo.Basicinfo[0].Context
	ui.attribution  = info.UserBaseInfo.Basicinfo[3].Context
	ui.startTime    = info.UserBaseInfo.Basicinfo[4].Context
	ui.basePackage  = info.UserBaseInfo.Basicinfo[5].Context
	ui.userStatus   = info.UserBaseInfo.Basicinfo[6].Context
	ui.billingType  = info.UserBaseInfo.Basicinfo[7].Context
	ui.groupName    = info.UserBaseInfo.Basicinfo[11].Context
	ui.bundlingTime = info.UserBaseInfo.Basicinfo[14].Context
	return
}

func FilterBusinessInfo(info * JSONbusinessInfo) (businessInfoString string, err error) {
	businessInfoString = ""
	for _, v1 := range info.BusinessList.Firstlevel {
		for _, v2 := range v1.Secondlevel {
			businessInfoString += v2.Secondvalue + "\n"
		}
	}
	return
}

func FilterConsumeInfo(info * JSONconsumeInfo) (ci ConsumeInfo, err error) {
	ci = NewConsumeInfo()
	reCalls, err := regexp.Compile(`^(\d{1,2}) 月消费/(.*)$`)
	if err != nil {
		return
	}
	reFlow, err := regexp.Compile(`^\d{4}(\d{2})/已使用优惠额度\(全时段包\+闲时包\)/(.*?) \+`)
	if err != nil {
		return
	}
	// 12 月消费/112.49
	// 201709/已使用优惠额度(全时段包+闲时包)/740.69 + 0.00MB
	for _, v1 := range info.ConsumeList.Firstlevel {
		for _, v2 := range v1.Secondlevel {
			//log.Println(v2.Secondvalue)
			b := []byte(v2.Secondvalue)
			if re := reCalls.FindSubmatch(b); len(re) > 0 {
				//log.Println(string(re[0]), string(re[1]), string(re[2]))
				month, _ := strconv.Atoi(string(re[1]))
				money := string(re[2])
				ci.callsConsume[month] = money
			} else if re := reFlow.FindSubmatch(b); len(re) > 0 {
				month, _ := strconv.Atoi(string(re[1]))
				flow := string(re[2])
				ci.flowConsume[month] = flow
			}
		}
	}
	return
}

func FilterUserBaseInfo(info * JSONuserBaseInfo) (ubi UserBaseInfo, err error) {
	ubi.recommendInfo = ""
	ubi.terminalType = info.UserBaseInfo.TerminalType
	ubi.isBoundTerminal = info.UserBaseInfo.IsBoundTerminal
	ubi.terminalChangeTime = info.UserBaseInfo.TerminalChangeTime
	ubi.isBroadBand = info.UserBaseInfo.IsBroadBand
	ubi.isRealNameUser = info.UserBaseInfo.IsRealnameUser
	ubi.isSphoneUser = info.UserBaseInfo.IsSphoneUser
	ubi.isUsimUser = info.UserBaseInfo.IsUsimUser
	ubi.is4gClient = info.UserBaseInfo.Is4GClient
	ubi.is4gBaseProduct = info.UserBaseInfo.Is4GBaseProduct
	ubi.isVotleTerminal = info.UserBaseInfo.IsVotleTerminal
	ubi.isVotleFunction = info.UserBaseInfo.IsVotleFunction
	ubi.isDoubleCardSlot = info.UserBaseInfo.IsDoubleCardSlot
	return
}