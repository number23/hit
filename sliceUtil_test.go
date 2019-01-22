package hit

import "testing"

func TestHasElem(t *testing.T) {
	foo := []int{23, 12, 891}
	if ok := HasElem(foo, 23); ok != true {
		t.Error("HasElem foo:", foo, "elem:", 23)
	}

	bar := []interface{}{43, "heyy", 12.1, false}
	if ok := HasElem(bar, "heyy"); ok != true {
		t.Error("HasElem bar:", bar, "elem:", "heyy")
	}
}
