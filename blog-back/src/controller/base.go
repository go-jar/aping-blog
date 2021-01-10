package controller

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	ctl "github.com/go-jar/gohttp/controller"
	"github.com/go-jar/gohttp/idgen"

	"blog/resource/log"
)

const (
	RemoteRealIpHeaderKey   = "REMOTE-REAL-IP"
	RemoteRealPortHeaderKey = "REMOTE-REAL-PORT"

	DownstreamServerIp = "127.0.0.1"

	MaxAutoParseBodyLen = 1024 * 1024

	TraceIdHeaderKey = "TRACE-ID"
	TraceIdQueryKey  = "traceId"
)

type BaseContext struct {
	*ctl.BaseContext

	ReqRawBody []byte

	QueryValues    url.Values
	RemoteRealAddr struct {
		Ip   string
		Port string
	}

	TraceId   []byte
	StartTime time.Time
}

func (bc *BaseContext) AfterAction() {
	body := bc.ResponseBody()
	bc.DebugLog([]byte("Response"), body)

	log.TraceLogger.Info(
		log.FormatTraceLog(&log.TraceLogArgs{
			TraceId:   bc.TraceId,
			Point:     []byte("AllTime"),
			StartTime: bc.StartTime,
			EndTime:   time.Now(),
			Msg:       body,
		}),
	)
}

func (bc *BaseContext) DebugLog(point, msg []byte) {
	log.AccessLogger.Debug(log.FormatAccessLog(bc.TraceId, point, msg))
}

func (bc *BaseContext) InfoLog(point, msg []byte) {
	log.AccessLogger.Info(log.FormatAccessLog(bc.TraceId, point, msg))
}

func (bc *BaseContext) NoticeLog(point, msg []byte) {
	log.AccessLogger.Notice(log.FormatAccessLog(bc.TraceId, point, msg))
}

func (bc *BaseContext) WarningLog(point, msg []byte) {
	log.AccessLogger.Warn(log.FormatAccessLog(bc.TraceId, point, msg))
}

func (bc *BaseContext) ErrorLog(point, msg []byte) {
	log.AccessLogger.Error(log.FormatAccessLog(bc.TraceId, point, msg))
}

func (bc *BaseContext) CriticalLog(point, msg []byte) {
	log.AccessLogger.Critical(log.FormatAccessLog(bc.TraceId, point, msg))
}

func (bc *BaseContext) AlertLog(point, msg []byte) {
	log.AccessLogger.Alert(log.FormatAccessLog(bc.TraceId, point, msg))
}

func (bc *BaseContext) EmergencyLog(point, msg []byte) {
	log.AccessLogger.Emergency(log.FormatAccessLog(bc.TraceId, point, msg))
}

type BaseController struct {
}

func (b *BaseController) NewActionContext(req *http.Request, resp http.ResponseWriter) ctl.ActionContext {
	context := &BaseContext{
		BaseContext: ctl.NewBaseContext(req, resp),

		StartTime: time.Now(),
	}

	if req.ContentLength < MaxAutoParseBodyLen {
		context.ReqRawBody, _ = ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(context.ReqRawBody))
	}

	_ = req.ParseForm()
	context.QueryValues = req.Form
	context.RemoteRealAddr.Ip, context.RemoteRealAddr.Port = b.parseRemoteAddr(req)

	context.TraceId = b.parseTraceId(context)
	context.NoticeLog(
		[]byte("Request from "+context.RemoteRealAddr.Ip+":"+context.RemoteRealAddr.Port),
		[]byte(req.RequestURI))

	context.ResponseWriter().Header().Add("X-Powered-By", "gohttp")

	return context
}

func (b *BaseController) parseRemoteAddr(req *http.Request) (string, string) {
	rs := strings.Split(req.RemoteAddr, ":")
	if rs[0] == DownstreamServerIp {
		ip := strings.TrimSpace(req.Header.Get(RemoteRealIpHeaderKey))
		port := strings.TrimSpace(req.Header.Get(RemoteRealPortHeaderKey))
		if ip != "" && port != "" {
			return ip, port
		}
	}

	return rs[0], rs[1]
}

func (b *BaseController) parseTraceId(context *BaseContext) []byte {
	traceId := strings.TrimSpace(context.Request().Header.Get(TraceIdHeaderKey))
	if len(traceId) != 0 {
		return []byte(traceId)
	}

	traceId = strings.TrimSpace(context.QueryValues.Get(TraceIdQueryKey))
	if len(traceId) != 0 {
		return []byte(traceId)
	}

	traceIdBytes, _ := idgen.DefaultTraceIdGenerator.GenerateId(context.RemoteRealAddr.Ip, context.RemoteRealAddr.Port)
	return traceIdBytes
}
