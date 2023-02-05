package safemap

import "testing"

func TestSafeMapInsertGet(t *testing.T) {
	m := New[int, int]()

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Insert(i, i*2)
			value, err := m.Get(i)
			if err != nil {
				t.Error(err)
			}
			if value != i {
				t.Errorf("%d should be %d", i, i*2)
			}
		}(i)
	}
}
