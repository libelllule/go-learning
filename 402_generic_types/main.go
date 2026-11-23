package main

import "fmt"

//оюоющенные типы (дженерики)
// any = interface{}

type List[T any] struct {
	next *List[T]
	val  T
}

func appendList[T any](e *List[T], node *List[T]) { //почему для обращения кк полям структуры обязательно передавать именно указатель
	current := e
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func main() {
	l := &List[int]{nil, 10}
	new := &List[int]{nil, 1}
	appendList(l, new)
	var current *List[int]
	current = l
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

/*
Дженерики (Generics) — это название всей концепции (обобщённого программирования) целиком.
Параметры типа (Type parameters) — это конкретный синтаксический механизм ([T comparable], [T any]), с помощью которого дженерики реализованы в Go.


List[int] — скобки после имени типа: это обобщённый тип List, в который подставлен int (по аналогии с map[string]int, где скобки тоже идут после map).


type List[T any] struct — здесь мы объявляем новый тип-шаблон. Квадратные скобки говорят: «List — это не готовый тип, а заготовка, которой нужен один параметр типа T».
next *List[T] — здесь мы этот тип используем. Поскольку List — это только заготовка, в Go нельзя просто написать next *List (компилятор спросит: «список чего именно?»). Поэтому мы обязаны уточнить в скобках, что next указывает на узел списка того же самого типа T.
То есть List[int] — это полноценное имя конкретного типа (как int или TreeNode), просто записанное с параметром в скобках.
*/
