type element struct {
    value int
    elem *list.Element
}

type LRUCache struct {
    cap int
    data map[int]element
    usedOrder *list.List
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        cap: capacity,
        data: make(map[int]element, capacity),
        usedOrder: list.New(),
    }
}

func (l *LRUCache) Get(key int) int {
    if e, ok := l.data[key]; ok {
        l.usedOrder.MoveToFront(e.elem)
        return e.value
    }
    return -1
}

func (l *LRUCache) Put(key int, value int) {
    if e, ok := l.data[key]; !ok {
        elem := l.usedOrder.PushFront(key)
        l.data[key] = element{
            value: value,
            elem: elem,
        }
    } else {
        l.usedOrder.MoveToFront(e.elem)
        e.value = value
        l.data[key] = e
    }

    if len(l.data) > l.cap {
        elem := l.usedOrder.Back()
        key := elem.Value.(int)
        delete(l.data, key)
        l.usedOrder.Remove(elem)
    }
}
