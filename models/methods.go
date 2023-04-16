package models

import "github.com/google/uuid"

func (a AppData) PrintAllLists() []ResponseBodyList {
	var lists []ResponseBodyList
	for _, list := range a.Lists {
		responseList := ResponseBodyList{
			Id:    list.Id,
			Owner: list.Owner,
			Todos: list.PrintTodos(),
		}
		lists = append(lists, responseList)
	}
	return lists
}

func (a *AppData) CreateList(requestbody *RequestBodyList, todos map[string]*ToDo, key string) {
	if a.Lists == nil {
		a.Lists = make(map[string]*ToDoList)
		a.Lists[key] = &ToDoList{
			Id:    key,
			Owner: requestbody.Owner,
			Todos: todos,
		}
	} else {
		a.Lists[key] = &ToDoList{
			Id:    key,
			Owner: requestbody.Owner,
			Todos: todos,
		}
	}
}

func (a *AppData) DeleteList(key string) {
	delete(a.Lists, key)
}

func (l ToDoList) PrintList() ResponseBodyList {
	var todos []ToDo
	for _, todo := range l.Todos {
		responseToDo := ToDo{
			Id:      todo.Id,
			Content: todo.Content,
		}
		todos = append(todos, responseToDo)
	}
	return ResponseBodyList{
		Owner: l.Owner,
		Todos: todos,
	}
}

func (l ToDoList) PrintTodos() []ToDo {
	var todos []ToDo
	for _, todo := range l.Todos {
		responseToDo := ToDo{
			Id:      todo.Id,
			Content: todo.Content,
		}
		todos = append(todos, responseToDo)
	}
	return todos
}

func (l *ToDoList) PatchList(owner string) {
	l.Owner = owner
}

func (l *ToDoList) DeleteToDo(key string) {
	delete(l.Todos, key)
}

func (l *ToDoList) CreateToDo(todo *ToDo) {
	key := uuid.New().String()
	l.Todos[key] = &ToDo{
		Id:      key,
		ListId:  l.Id,
		Content: todo.Content,
	}
}

func (t ToDo) PrintToDo() ToDo {
	return ToDo{
		ListId:  t.ListId,
		Content: t.Content,
	}
}

func (t *ToDo) PatchToDo(content string) {
	t.Content = content
}
