package ringbuf

type RingBuf struct {
  data []interface{}
  size int
  head int
  tail int
}

func New(size int) RingBuf {
  return RingBuf{
    data: make([]interface{}, size),
    size: size,
    head: 0,
    tail: 0,
  }
}

func (b *RingBuf) Read() interface{} {
  val := b.data[b.head]
  if b.head += 1; b.head >= b.size {
    b.head = 0
  }
  return val
}

func (b RingBuf) Peek() interface{} {
  return b.data[b.head]
}

func (b *RingBuf) Write(val interface{}) {
  b.data[b.tail] = val
  if b.tail += 1; b.tail >= b.size {
    b.tail = 0
  }
}

func (b *RingBuf) Clear() {
  b.data = make([]interface{}, b.size)
}

