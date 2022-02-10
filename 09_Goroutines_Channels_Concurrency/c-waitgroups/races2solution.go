

var m = map[string]int{"a": 1}
var lock = sync.RWMutex{}

func main() {
	go Read()
	time.Sleep(1 * time.Second)
	go Write()
	time.Sleep(1 * time.Minute)
}

func Read() {
	for {
		read()
	}
}

func Write() {
	for {
		write()
	}
}

func read() {
	lock.RLock()
	defer lock.RUnlock()
	_ = m["a"]
}

func write() {
	lock.Lock()
	defer lock.Unlock()
	m["b"] = 2
}
