package template

const (
	FindOneMethod = `
func (uc *{{.upperStartCamelObject}}UseCase) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	return uc.repo.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
}
`
	FindOneInterface = `FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error)`

	FindAllMethod = `
func (uc *{{.upperStartCamelObject}}UseCase) FindAll(ctx context.Context, where []where_builder.Expr) ([]*{{.upperStartCamelObject}}, error) {
	return uc.repo.FindAll(ctx, where)
}
`
	FindAllInterface = `FindAll(ctx context.Context, where []where_builder.Expr) ([]*{{.upperStartCamelObject}}, error)`

	PaginateMethod = `
func (uc *{{.upperStartCamelObject}}UseCase) Paginate(ctx context.Context, where []where_builder.Expr, limit, offset int64, orderBy string) ([]*{{.upperStartCamelObject}}, error) {
	return uc.repo.Paginate(ctx, where, limit, offset, orderBy)
}
`
	PaginateInterface = `Paginate(ctx context.Context, where []where_builder.Expr, limit, offset int64, orderBy string) ([]*{{.upperStartCamelObject}}, error)`

	CountMethod = `
func (uc *{{.upperStartCamelObject}}UseCase) Count(ctx context.Context, where []where_builder.Expr) (int64, error) {
	return uc.repo.Count(ctx, where)
}
`
	CountInterface = `Count(ctx context.Context, where []where_builder.Expr) (int64, error)`
)
