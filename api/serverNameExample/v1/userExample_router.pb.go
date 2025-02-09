// Code generated by https://github.com/zhufuyi/sponge, DO NOT EDIT.

package v1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	errcode "github.com/zhufuyi/sponge/pkg/errcode"
	middleware "github.com/zhufuyi/sponge/pkg/gin/middleware"
	zap "go.uber.org/zap"
	strings "strings"
)

// import packages: strings. context. errcode. middleware. zap. gin.

type UserExampleServiceLogicer interface {
	Create(ctx context.Context, req *CreateUserExampleRequest) (*CreateUserExampleReply, error)
	DeleteByID(ctx context.Context, req *DeleteUserExampleByIDRequest) (*DeleteUserExampleByIDReply, error)
	DeleteByIDs(ctx context.Context, req *DeleteUserExampleByIDsRequest) (*DeleteUserExampleByIDsReply, error)
	GetByID(ctx context.Context, req *GetUserExampleByIDRequest) (*GetUserExampleByIDReply, error)
	List(ctx context.Context, req *ListUserExampleRequest) (*ListUserExampleReply, error)
	ListByIDs(ctx context.Context, req *ListUserExampleByIDsRequest) (*ListUserExampleByIDsReply, error)
	UpdateByID(ctx context.Context, req *UpdateUserExampleByIDRequest) (*UpdateUserExampleByIDReply, error)
}

type UserExampleServiceOption func(*userExampleServiceOptions)

type userExampleServiceOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *userExampleServiceOptions) apply(opts ...UserExampleServiceOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithUserExampleServiceHTTPResponse() UserExampleServiceOption {
	return func(o *userExampleServiceOptions) {
		o.isFromRPC = false
	}
}

func WithUserExampleServiceRPCResponse() UserExampleServiceOption {
	return func(o *userExampleServiceOptions) {
		o.isFromRPC = true
	}
}

func WithUserExampleServiceResponser(responser errcode.Responser) UserExampleServiceOption {
	return func(o *userExampleServiceOptions) {
		o.responser = responser
	}
}

func WithUserExampleServiceLogger(zapLog *zap.Logger) UserExampleServiceOption {
	return func(o *userExampleServiceOptions) {
		o.zapLog = zapLog
	}
}

func WithUserExampleServiceErrorToHTTPCode(e ...*errcode.Error) UserExampleServiceOption {
	return func(o *userExampleServiceOptions) {
		o.httpErrors = e
	}
}

func WithUserExampleServiceRPCStatusToHTTPCode(s ...*errcode.RPCStatus) UserExampleServiceOption {
	return func(o *userExampleServiceOptions) {
		o.rpcStatus = s
	}
}

func WithUserExampleServiceWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) UserExampleServiceOption {
	return func(o *userExampleServiceOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterUserExampleServiceRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic UserExampleServiceLogicer,
	opts ...UserExampleServiceOption) {

	o := &userExampleServiceOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &userExampleServiceRouter{
		iRouter:               iRouter,
		groupPathMiddlewares:  groupPathMiddlewares,
		singlePathMiddlewares: singlePathMiddlewares,
		iLogic:                iLogic,
		iResponse:             o.responser,
		zapLog:                o.zapLog,
		wrapCtxFn:             o.wrapCtxFn,
	}
	r.register()
}

type userExampleServiceRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                UserExampleServiceLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *userExampleServiceRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/userExample", r.withMiddleware("POST", "/api/v1/userExample", r.Create_0)...)
	r.iRouter.Handle("DELETE", "/api/v1/userExample/:id", r.withMiddleware("DELETE", "/api/v1/userExample/:id", r.DeleteByID_0)...)
	r.iRouter.Handle("POST", "/api/v1/userExample/delete/ids", r.withMiddleware("POST", "/api/v1/userExample/delete/ids", r.DeleteByIDs_0)...)
	r.iRouter.Handle("PUT", "/api/v1/userExample/:id", r.withMiddleware("PUT", "/api/v1/userExample/:id", r.UpdateByID_0)...)
	r.iRouter.Handle("GET", "/api/v1/userExample/:id", r.withMiddleware("GET", "/api/v1/userExample/:id", r.GetByID_0)...)
	r.iRouter.Handle("POST", "/api/v1/userExample/list/ids", r.withMiddleware("POST", "/api/v1/userExample/list/ids", r.ListByIDs_0)...)
	r.iRouter.Handle("POST", "/api/v1/userExample/list", r.withMiddleware("POST", "/api/v1/userExample/list", r.List_0)...)

}

func (r *userExampleServiceRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
	handlerFns := []gin.HandlerFunc{}

	// determine if a route group is hit or miss, left prefix rule
	for groupPath, fns := range r.groupPathMiddlewares {
		if groupPath == "" || groupPath == "/" {
			handlerFns = append(handlerFns, fns...)
			continue
		}
		size := len(groupPath)
		if len(path) < size {
			continue
		}
		if groupPath == path[:size] {
			handlerFns = append(handlerFns, fns...)
		}
	}

	// determine if a single route has been hit
	key := strings.ToUpper(method) + "->" + path
	if fns, ok := r.singlePathMiddlewares[key]; ok {
		handlerFns = append(handlerFns, fns...)
	}

	return append(handlerFns, fn)
}

func (r *userExampleServiceRouter) Create_0(c *gin.Context) {
	req := &CreateUserExampleRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = c
	}

	out, err := r.iLogic.Create(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userExampleServiceRouter) DeleteByID_0(c *gin.Context) {
	req := &DeleteUserExampleByIDRequest{}
	var err error

	if err = c.ShouldBindUri(req); err != nil {
		r.zapLog.Warn("ShouldBindUri error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	if err = c.ShouldBindQuery(req); err != nil {
		r.zapLog.Warn("ShouldBindQuery error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = c
	}

	out, err := r.iLogic.DeleteByID(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userExampleServiceRouter) DeleteByIDs_0(c *gin.Context) {
	req := &DeleteUserExampleByIDsRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = c
	}

	out, err := r.iLogic.DeleteByIDs(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userExampleServiceRouter) UpdateByID_0(c *gin.Context) {
	req := &UpdateUserExampleByIDRequest{}
	var err error

	if err = c.ShouldBindUri(req); err != nil {
		r.zapLog.Warn("ShouldBindUri error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = c
	}

	out, err := r.iLogic.UpdateByID(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userExampleServiceRouter) GetByID_0(c *gin.Context) {
	req := &GetUserExampleByIDRequest{}
	var err error

	if err = c.ShouldBindUri(req); err != nil {
		r.zapLog.Warn("ShouldBindUri error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	if err = c.ShouldBindQuery(req); err != nil {
		r.zapLog.Warn("ShouldBindQuery error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = c
	}

	out, err := r.iLogic.GetByID(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userExampleServiceRouter) ListByIDs_0(c *gin.Context) {
	req := &ListUserExampleByIDsRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = c
	}

	out, err := r.iLogic.ListByIDs(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userExampleServiceRouter) List_0(c *gin.Context) {
	req := &ListUserExampleRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = c
	}

	out, err := r.iLogic.List(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}
