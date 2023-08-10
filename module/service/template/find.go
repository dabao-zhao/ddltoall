package template

const (
	FindOneMethod = `
func (s *{{.upperStartCamelObject}}Service) FindOne(ctx context.Context, req *types.FindOne{{.upperStartCamelObject}}Request) (*types.FindOne{{.upperStartCamelObject}}Response, error) {
	data, err := s.useCase.FindOne(ctx, req.{{.upperStartCamelPrimaryKey}})
	if err != nil {
		s.logger.Error(err)
		return nil, errors.New("find one err")
	}
	var resp *types.FindOne{{.upperStartCamelObject}}Response
	_ = copier.Copy(resp, data)
	return resp, nil
}
`
	FindAllMethod = `
func (s *{{.upperStartCamelObject}}Service) FindAll(ctx context.Context, req *types.FindAll{{.upperStartCamelObject}}Request) (*types.FindAll{{.upperStartCamelObject}}Response, error) {
	var where []where_builder.Expr

	/** where **/
	
	data, err := s.useCase.FindAll(ctx, where)
	if err != nil {
		s.logger.Error(err)
		return nil, errors.New("find all err")
	}

	resp := &types.FindAll{{.upperStartCamelObject}}Response{Items: nil}
	_ = copier.Copy(&resp.Items, data)
	return resp, nil
}
`
	PaginateMethod = `
func (s *{{.upperStartCamelObject}}Service) Paginate(ctx context.Context, req *types.Paginate{{.upperStartCamelObject}}Request) (*types.Paginate{{.upperStartCamelObject}}Response, error) {
	var where []where_builder.Expr

	if req.PageSize == 0 {
		req.PageSize = 20
	}
	if req.Page == 0 {
		req.Page = 1
	}

	/** where **/
	
	data, err := s.useCase.Paginate(ctx, where, req.PageSize, (req.Page-1)*req.PageSize, req.OrderBy)
	if err != nil {
		s.logger.Error(err)
		return nil, errors.New("find err")
	}
	count, err := s.useCase.Count(ctx, where)
	if err != nil {
		s.logger.Error(err)
		return nil, errors.New("find err")
	}
	resp := &types.Paginate{{.upperStartCamelObject}}Response{Items: nil, PageCount: 0}
	_ = copier.Copy(&resp.Items, data)
	resp.PageCount = int64(math.Ceil(float64(count) / float64(req.PageSize)))
	return resp, nil
}
`
)
