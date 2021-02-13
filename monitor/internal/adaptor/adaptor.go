package adaptor

type Adaptor struct {
	TechAdaptor
}

func NewAdaptor() *Adaptor {
	return &Adaptor{
		TechAdaptor: NewTechMongo(4),
	}
}

type TechAdaptor interface {
	CreateTech() (int, error)
	UpdateTech(id int) (string, error)
	GetTechByComponentName(component string, resource string) (string, bool)
}
