package service

type Repo interface{

}


type CatalogService struct{
	repo Repo
}



func NewService(repo Repo) *CatalogService{
	return &CatalogService{
		repo: repo,
	}
}