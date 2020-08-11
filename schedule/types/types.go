package types

import (
	"time"
)

// 任务的struct
type Cron struct {
	ID       int
	Name     string
	Exp      string
	Command  string
	Desc     string
	IsDel    int
	IsOnline int
	Types    int
}

// 增加执行时间的任务
type CronExpr struct {
	Cron *Cron
	Next time.Time
}

// 进行排序
type CronSort []*CronExpr

func (p CronSort) Len() int { return len(p) }

func (p CronSort) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p CronSort) Less(i, j int) bool {
	return p[i].Next.Before(p[j].Next)
}

// 输出任务的记录
type CronExcel struct {
	Cron   *Cron
	Output []byte
	Err    error
}

//// 常驻内存输出的记录
//type CronDeamonExcel struct {
//	Cron   *Deamon
//	Output []byte
//	Err    error
//}
