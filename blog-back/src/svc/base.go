package svc

import (
	"blog/resource/log"
)

type BaseSvc struct {
	TraceId []byte
}

func (bs *BaseSvc) DebugLog(point, msg []byte) {
	log.AccessLogger.Debug(log.FormatAccessLog(bs.TraceId, point, msg))
}

func (bs *BaseSvc) InfoLog(point, msg []byte) {
	log.AccessLogger.Info(log.FormatAccessLog(bs.TraceId, point, msg))
}

func (bs *BaseSvc) NoticeLog(point, msg []byte) {
	log.AccessLogger.Notice(log.FormatAccessLog(bs.TraceId, point, msg))
}

func (bs *BaseSvc) WarningLog(point, msg []byte) {
	log.AccessLogger.Warn(log.FormatAccessLog(bs.TraceId, point, msg))
}

func (bs *BaseSvc) ErrorLog(point, msg []byte) {
	log.AccessLogger.Error(log.FormatAccessLog(bs.TraceId, point, msg))
}

func (bs *BaseSvc) CriticalLog(point, msg []byte) {
	log.AccessLogger.Critical(log.FormatAccessLog(bs.TraceId, point, msg))
}

func (bs *BaseSvc) AlertLog(point, msg []byte) {
	log.AccessLogger.Alert(log.FormatAccessLog(bs.TraceId, point, msg))
}

func (bs *BaseSvc) EmergencyLog(point, msg []byte) {
	log.AccessLogger.Emergency(log.FormatAccessLog(bs.TraceId, point, msg))
}
