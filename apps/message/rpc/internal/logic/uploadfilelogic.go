package logic

import (
	"bufio"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"jt-chat/apps/message/model"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"
	"os"
	"path/filepath"

	"jt-chat/apps/message/rpc/internal/svc"
	"jt-chat/apps/message/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadFileLogic) UploadFile(in *pb.UploadFileIn) (*pb.UploadFileOut, error) {
	var (
		uid     string
		fileId  string
		osFile  *os.File
		newFile *model.File
		err     error
		out     *pb.UploadFileOut
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	fileId = uuid.New().String()
	out = &pb.UploadFileOut{FilePath: filepath.Join(constant.LocalPath, fileId)}
	osFile, err = os.OpenFile(out.FilePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return nil, xerr.CustomErr(xerr.FileOpenErr, l.ctx, errors.Wrapf(err, "打开文件%s", out.FilePath))
	}
	defer osFile.Close()
	writer := bufio.NewWriter(osFile)
	_, err = writer.Write(in.Data)
	err = writer.Flush()
	if err != nil {
		return nil, xerr.CustomErr(xerr.FileSaveErr, l.ctx, errors.Wrapf(err, "保存文件%s", out.FilePath))
	}
	newFile = &model.File{
		FileId: fileId,
		Hash:   fmt.Sprintf("%x", md5.Sum(in.Data)),
		Uid:    uid,
		Name:   in.Name,
		Ext:    in.Ext,
		Size:   int64(len(in.Data)),
		Path:   out.FilePath,
	}
	_, err = l.svcCtx.FileModel.Insert(l.ctx, newFile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "保存文件信息[%+v]", newFile))
	}
	return out, nil
}
