package adaptor

type TechMongo struct {
	db int
}

func NewTechMongo(db int) *TechMongo {
	return &TechMongo{db: db}
}

func (a *TechMongo) CreateTech() (int, error) {
	return 4, nil
}

func (a *TechMongo) UpdateTech(id int) (string, error) {
	return "4", nil
}

// GetTechByComponentName returns technique by component/resource couple
// if technique is not found returns false as second value
func (a *TechMongo) GetTechByComponentName(component string, resource string) (string, bool) {
	return "4", false
}
