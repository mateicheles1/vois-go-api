package models

func (l ToDoList) PrintList() ToDoList {
	return ToDoList{
		Owner: l.Owner,
		Todos: l.Todos,
	}
}

func (t ToDo) PrintToDo() ToDo {
	return ToDo{
		ListId:  t.ListId,
		Content: t.Content,
	}
}