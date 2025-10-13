package inputs

type CreateSupplyInput struct {
	Name         string
	Description  string
	Unit         string
	MinThreshold int
}

type UpdateSupplyInput struct {
	ID     string
	Fields map[string]interface{}
}
