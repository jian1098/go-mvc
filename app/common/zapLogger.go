package common

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetLogger() *zap.SugaredLogger {
	var coreArr []zapcore.Core

	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()            //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder     //显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig) //输出到控制台

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   os.Getenv("LOG_PATH") + "/info.log",      //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    cast.ToInt(os.Getenv("LOG_MAX_SIZE")),    //文件大小限制,单位MB
		MaxBackups: cast.ToInt(os.Getenv("LOG_MAX_BACKUPS")), //最大保留日志文件数量
		MaxAge:     cast.ToInt(os.Getenv("LOG_MAX_AGE")),     //日志文件保留天数
		Compress:   cast.ToBool(os.Getenv("LOG_COMPRESS")),   //是否压缩处理
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	//error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   os.Getenv("LOG_PATH") + "/error.log",     //日志文件存放目录
		MaxSize:    cast.ToInt(os.Getenv("LOG_MAX_SIZE")),    //文件大小限制,单位MB
		MaxBackups: cast.ToInt(os.Getenv("LOG_MAX_BACKUPS")), //最大保留日志文件数量
		MaxAge:     cast.ToInt(os.Getenv("LOG_MAX_AGE")),     //日志文件保留天数
		Compress:   cast.ToBool(os.Getenv("LOG_COMPRESS")),   //是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	return zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()).Sugar() //zap.AddCaller()为显示文件名和行号，可省略
}
