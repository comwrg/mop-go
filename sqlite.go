package mop

import (
	"database/sql"
	"fmt"
	"strings"
	_"github.com/mattn/go-sqlite3"
)

const (
	TABLE = "info"
)

type Sqlite struct {
	db * sql.DB
}

func (s * Sqlite) Init() (err error) {
	s.db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		return
	}
	s.db.SetMaxOpenConns(1)
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS %s (
		/* user information */
		手机号 TEXT NOT NULL PRIMARY KEY,
		姓名 TEXT DEFAULT '',
		归属 TEXT DEFAULT '',
		开打时间 DATETIME DEFAULT '',
		基本套餐 TEXT DEFAULT '',
		用户状态 TEXT DEFAULT '',
		出账类型 TEXT DEFAULT '',
		集团名称 TEXT DEFAULT '',
		最晚捆绑时间 TEXT DEFAULT '',

		/* business information */
		业务信息 TEXT DEFAULT '',

		/* consume information */
		一月消费 DOUBLE DEFAULT 0,
		二月消费 DOUBLE DEFAULT 0,
		三月消费 DOUBLE DEFAULT 0,
		四月消费 DOUBLE DEFAULT 0,
		五月消费 DOUBLE DEFAULT 0,
		六月消费 DOUBLE DEFAULT 0,
		七月消费 DOUBLE DEFAULT 0,
		八月消费 DOUBLE DEFAULT 0,
		九月消费 DOUBLE DEFAULT 0,
		十月消费 DOUBLE DEFAULT 0,
		十一月消费 DOUBLE DEFAULT 0,
		十二月消费 DOUBLE DEFAULT 0,
		一月流量 DOUBLE DEFAULT 0,
		二月流量 DOUBLE DEFAULT 0,
		三月流量 DOUBLE DEFAULT 0,
		四月流量 DOUBLE DEFAULT 0,
		五月流量 DOUBLE DEFAULT 0,
		六月流量 DOUBLE DEFAULT 0,
		七月流量 DOUBLE DEFAULT 0,
		八月流量 DOUBLE DEFAULT 0,
		九月流量 DOUBLE DEFAULT 0,
		十月流量 DOUBLE DEFAULT 0,
		十一月流量 DOUBLE DEFAULT 0,
		十二月流量 DOUBLE DEFAULT 0,

		/* user base information */
		推荐信息 TEXT DEFAULT '',
		终端类型 TEXT DEFAULT '',
		绑定终端 TEXT DEFAULT '',
		换机时长 TEXT DEFAULT '',
		是否办理宽带 TEXT DEFAULT '',
		是否实名用户 TEXT DEFAULT '',
		是否智能机用户 TEXT DEFAULT '',
		是否USIM卡用户 TEXT DEFAULT '',
		是否4G终端 TEXT DEFAULT '',
		是否4G套餐 TEXT DEFAULT '',
		是否VoTLE终端用户 TEXT DEFAULT '',
		是否开通VoTLE服务 TEXT DEFAULT '',
		卡槽类型 TEXT DEFAULT ''
	)
	`
	_, err = s.db.Exec(
		fmt.Sprintf(sqlStmt, TABLE, ),
		)
	return
}

func (s * Sqlite) Insert(mobile string) (err error) {
	sqlStmt := `INSERT INTO %s (手机号) VALUES(?)`
	_, err = s.db.Exec(
		fmt.Sprintf(sqlStmt, TABLE),
		mobile,
	)
	// ignore
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			err = nil
		}
	}
	return
}

func (s * Sqlite) UpdateUserInfo(mobile string, ui UserInfo) (err error) {
	sqlStmt := `
	UPDATE %s SET
		姓名 = (?),
		归属 = (?),
		开打时间 = (?),
		基本套餐 = (?),
		用户状态 = (?),
		出账类型 = (?),
		集团名称 = (?),
		最晚捆绑时间 = (?)
		WHERE 手机号 = (?)
	`
	_, err = s.db.Exec(
		fmt.Sprintf(sqlStmt, TABLE),
		ui.name, ui.attribution, ui.startTime, ui.basePackage,
		ui.userStatus, ui.billingType, ui.groupName, ui.bundlingTime, mobile,
	)
	return
}

func (s * Sqlite) UpdateBusinessInfo(mobile, info string) (err error) {
	sqlStmt := `
	UPDATE %s SET
		业务信息 = (?)
		WHERE 手机号 = (?)
	`
	_, err = s.db.Exec(
		fmt.Sprintf(sqlStmt, TABLE),
		info,
		mobile,
	)
	return
}

func (s * Sqlite) UpdateConsumeInfo(mobile string, info ConsumeInfo) (err error) {
	sqlSmst := `
	UPDATE %s SET
		一月消费 = (?),
		二月消费 = (?),
		三月消费 = (?),
		四月消费 = (?),
		五月消费 = (?),
		六月消费 = (?),
		七月消费 = (?),
		八月消费 = (?),
		九月消费 = (?),
		十月消费 = (?),
		十一月消费 = (?),
		十二月消费 = (?),
		一月流量 = (?),
		二月流量 = (?),
		三月流量 = (?),
		四月流量 = (?),
		五月流量 = (?),
		六月流量 = (?),
		七月流量 = (?),
		八月流量 = (?),
		九月流量 = (?),
		十月流量 = (?),
		十一月流量 = (?),
		十二月流量 = (?)
		WHERE 手机号 = (?)
	`
	_, err = s.db.Exec(
		fmt.Sprintf(sqlSmst, TABLE),

		info.callsConsume[1],
		info.callsConsume[2],
		info.callsConsume[3],
		info.callsConsume[4],
		info.callsConsume[5],
		info.callsConsume[6],
		info.callsConsume[7],
		info.callsConsume[8],
		info.callsConsume[9],
		info.callsConsume[10],
		info.callsConsume[11],
		info.callsConsume[12],

		info.flowConsume[1],
		info.flowConsume[2],
		info.flowConsume[3],
		info.flowConsume[4],
		info.flowConsume[5],
		info.flowConsume[6],
		info.flowConsume[7],
		info.flowConsume[8],
		info.flowConsume[9],
		info.flowConsume[10],
		info.flowConsume[11],
		info.flowConsume[12],

		mobile,
	)
	return
}

func (s * Sqlite) UpdateUserBaseInfo(mobile string, ubi UserBaseInfo) (err error) {
	sqlSmst := `
	UPDATE %s SET
		推荐信息 = (?),
		终端类型 = (?),
		绑定终端 = (?),
		换机时长 = (?),
		是否办理宽带 = (?),
		是否实名用户 = (?),
		是否智能机用户 = (?),
		是否USIM卡用户 = (?),
		是否4G终端 = (?),
		是否4G套餐 = (?),
		是否VoTLE终端用户 = (?),
		是否开通VoTLE服务 = (?),
		卡槽类型 = (?)
		WHERE 手机号 = (?)
	`
	_, err = s.db.Exec(
		fmt.Sprintf(sqlSmst, TABLE),
		ubi.recommendInfo, ubi.terminalType, ubi.isBoundTerminal,
		ubi.terminalChangeTime, ubi.isBroadBand, ubi.isRealNameUser, ubi.isSphoneUser,
		ubi.isUsimUser, ubi.is4gClient, ubi.is4gBaseProduct, ubi.isVotleTerminal,
		ubi.isVotleFunction, ubi.isDoubleCardSlot,
		mobile,
	)
	return
}

func (s * Sqlite) Close() (err error) {
	err = s.db.Close()
	return
}

