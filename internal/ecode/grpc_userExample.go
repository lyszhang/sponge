package ecode

import "github.com/zhufuyi/sponge/pkg/errcode"

const (
	// todo must be modified manually
	// 每个资源名称对应唯一编号，编号范围1~1000，如果存在编号相同，启动服务会报错
	_userExampleNO = 1
	// userExample对应的中文名称
	_userExampleName = "userExample_cn_name"
)

// nolint
// 服务级别错误码
var (
	StatusCreateUserExample = errcode.NewGRPCStatus(errcode.GCode(_userExampleNO)+1, "创建"+_userExampleName+"失败") // todo 补充错误码注释，例如 400101
	StatusDeleteUserExample = errcode.NewGRPCStatus(errcode.GCode(_userExampleNO)+2, "删除"+_userExampleName+"失败")
	StatusUpdateUserExample = errcode.NewGRPCStatus(errcode.GCode(_userExampleNO)+3, "更新"+_userExampleName+"失败")
	StatusGetUserExample    = errcode.NewGRPCStatus(errcode.GCode(_userExampleNO)+4, "获取"+_userExampleName+"失败")
	StatusListUserExample   = errcode.NewGRPCStatus(errcode.GCode(_userExampleNO)+5, "获取"+_userExampleName+"列表失败")
	// 每添加一个错误码，在上一个错误码基础上+1
)
