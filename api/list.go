package api

import "sort"

type ById []Todo

func (a ById) Len() int{
	return len(a)
}
func (a ById) Less(i, j int) bool {
	return a[i].Id < a[j].Id
}
func (a ById) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func List() []Todo {
	sort.Sort(ById(Todos))
	return Todos
}
