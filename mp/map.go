package mp

type Map struct {
	it *pkg.Atomiterator
	mp map[int64]int64
}

func NewMap() *Map {
	mp := &Map{
		mp: make(map[int64]int64),
		it: prg.AtomIterator
		} //map - это хранилище, структура ключ-значение (сначала ключ, потом значение)
	              //создает пустой объект и возвращает на него ссылку
	return mp
}





//пары_по_языкам_программирования
//ключи в мапе уникальны
func main() {
	m := map[int64]string{}
	m[1] = "one"
	m[2] = "two"

	v, ok := m[3]
	fmt.Println(v, ok)
}

чтоб пользоваться мапой, нудно инициализироваться


если написать как var m map[int64]string, то работать не будет



// со слайссом так же, если не проинициализировать 
var s []string

s = append(s, )
//если функция берет слайс по ссылке, как апенд, то она возвращает значение
 
s[1]= "one" //будет ошибка "вы вышли за пределы слайса"



//























































func (m *Map) Add(data int64) (id int64) {
	id = m.it.Get
	
	if id = nil {
		
	}
	return id
}

func (m *Map) Get(id int64) (data int64, ok bool) {
	data, ok = m.mp[id]
	return
}

func (m *Map) Del(id int64) {
	delete(m.mp, id)
}

// NewList создает новый список
func NewList() *List {
	
	return 
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (id int64) {
	
	return 
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {

	return
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	
	return
}


// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	
			return o
	
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {

			return ok
		
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	
	return
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	
		return 
}

// Clear очищает список
func (mp *Map) Clear() {
	id = it.m.Get
	mp := nil
	return	
}

// Print выводит список в консоль
func (mp *Map) Print() {
	fmt.Println(it.m.Get, "   ", mp)
}









