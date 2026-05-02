type mockRepo struct{}

func (m mockRepo) GetAll() ([]models.Book, error) {
	return []models.Book{{ID: 1, Title: "Mock"}}, nil
}

func (m mockRepo) GetByID(id int) (models.Book, error) {
	return models.Book{ID: int64(id), Title: "Mock"}, nil
}

func (m mockRepo) Create(b *models.Book) error {
	b.ID = 1
	return nil
}

func (m mockRepo) Update(b *models.Book) error {
	return nil
}

func (m mockRepo) Delete(id int) error {
	return nil
}