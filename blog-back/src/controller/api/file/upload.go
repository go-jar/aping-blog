package file

import (
	"io"

	"github.com/go-jar/goerror"

	"blog/errno"
)

func (fc *FileController) UploadAction(context *FileContext) {
	if err := fc.VerifyToken(context.ApiContext); err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return
	}

	f, e := fc.parseUploadActionParams(context)
	if e != nil {
		context.ApiData.Err = e
		return
	}

	imgUrl, err := context.tCloudSvc.PutImg(f)
	if err != nil {
		context.ApiData.Err = goerror.New(errno.ETCloudCosError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"ImgUrl": imgUrl,
		"RequestId": context.TraceId,
	}
}

func (fc *FileController) parseUploadActionParams(context *FileContext) (io.Reader, *goerror.Error) {
	f, _, err := context.Request().FormFile("file")
	if err != nil {
		return nil, goerror.New(errno.ECommonInvalidArg, "invalid param: file")
	}

	return f, nil
}
