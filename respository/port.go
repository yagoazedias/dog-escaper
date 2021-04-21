package respository

type PortRepository struct {}

type PortRepositoryInterface interface {
	GetLastStatus() (string, error)
	UpdateLastStatus(status string) error
}

func (r PortRepository) GetLastStatus() (string, error) {
	return "", nil
}

func (r PortRepository) UpdateLastStatus(status string) error {
	return nil
}