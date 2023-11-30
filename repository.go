package main

type Repository interface {
	GetGoatActions() []string
}

type hardcodedRepository struct{}

func NewRepository() Repository {
	return &hardcodedRepository{}
}

func (r *hardcodedRepository) GetGoatActions() []string {
	return []string{
		"eating grass",
		"riding a scooter",
		"drinking a coffee",
	}
}
