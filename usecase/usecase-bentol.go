func (t *todo) Create(task string) error {
	// 登録するための構造体を作成
	todo := model.NewTodo(task)
	// DB にデータを登録
	if err := t.todoRepository.Create(todo); err != nil {
		return err
	}
	return nil
}

func (t *todo) Update(id int, task string, status model.TaskStatus) error {
	// 更新用の構造体を生成
	todo := model.NewUpdateTodo(id, task, status)
	// DBのデータの更新
	if err := t.todoRepository.Update(todo); err != nil {
		return err
	}
	return nil
}

func (t *todo) Find(id int) (*model.Todo, error) {
	todo, err := t.todoRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todo) FindAll() ([]*model.Todo, error) {
	todo, err := t.todoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todo) Delete(id int) error {
	if err := t.todoRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
