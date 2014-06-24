package publisher

import
    "testing"

func TestTranslate(t *testing.T){
    assertEquals(t, "Mat", "Ma2t")
}

func assertEquals(t *testing.T, value1 string, value2 string) {
    if value1 != value2 {
        t.Error("This failed")
    }
}