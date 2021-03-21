package tcloud

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-jar/mysql"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"

	"blog/conf"
	"blog/resource"
	"blog/svc"
)

type Svc struct {
	*svc.BaseSvc

	SqlOrm     *mysql.SimpleOrm
	EntityName string
}

func NewSvc(traceId []byte) *Svc {
	return &Svc{
		BaseSvc: &svc.BaseSvc{
			TraceId: traceId,
		},
		SqlOrm:     mysql.NewSimpleOrm(traceId, resource.MysqlClientPool, true),
		EntityName: "cos",
	}
}

func (s *Svc) logStatus(err error) error {
	if err == nil {
		return nil
	}

	if cos.IsNotFoundError(err) {
		// WARN
		s.WarningLog([]byte("TCloudCosSvc"), []byte(fmt.Sprintf("WARN: Resource is not existed")))
	} else if e, ok := cos.IsCOSError(err); ok {
		s.ErrorLog([]byte("TCloudCosSvc"), []byte(fmt.Sprintf("ERROR: Code: %v\n", e.Code)))
		s.ErrorLog([]byte("TCloudCosSvc"), []byte(fmt.Sprintf("ERROR: Message: %v\n", e.Message)))
		s.ErrorLog([]byte("TCloudCosSvc"), []byte(fmt.Sprintf("ERROR: Resource: %v\n", e.Resource)))
		s.ErrorLog([]byte("TCloudCosSvc"), []byte(fmt.Sprintf("ERROR: RequestId: %v\n", e.RequestID)))
		// ERROR
	} else {
		s.ErrorLog([]byte("TCloudCosSvc"), []byte(fmt.Sprintf("ERROR: %v\n", err)))
		// ERROR
	}
	return err
}

func (s *Svc) PutImg(img io.Reader) (string, error) {
	u, _ := url.Parse("https://" + conf.TcloudAccount.CosHost)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client {
		Transport: &cos.AuthorizationTransport {
			SecretID:  conf.TcloudAccount.SecretId,
			SecretKey: conf.TcloudAccount.SecretKey,
			Transport: &debug.DebugRequestTransport {
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	opt := &cos.ObjectPutOptions {
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions {
			ContentType: "image/png",
		},
	}

	id, err := s.SqlOrm.IdGenerator().GenerateId(s.EntityName)
	if err != nil {
		return "", err
	}

	imgName := strconv.FormatInt(id,10) + ".png"
	imgUrl := "https://" + conf.TcloudAccount.CosHost + "/" + imgName
	_, err = c.Object.Put(context.Background(), imgName, img, opt)

	return imgUrl, s.logStatus(err)
}
