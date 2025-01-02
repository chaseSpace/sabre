package sslice

import (
	"reflect"
	"testing"
)

// TestClone_EmptySlice_ShouldReturnEmptySlice tests cloning an empty slice.
func TestClone_EmptySlice_ShouldReturnEmptySlice(t *testing.T) {
	original := &Slice[int]{data: []int{}}
	cloned := original.Clone()

	if !reflect.DeepEqual(cloned.data, original.data) {
		t.Errorf("Expected cloned slice to be empty, got %v", cloned.data)
	}
}

// TestClone_NonEmptySlice_ShouldReturnClonedSlice tests cloning a non-empty slice.
func TestClone_NonEmptySlice_ShouldReturnClonedSlice(t *testing.T) {
	original := &Slice[int]{data: []int{1, 2, 3}}
	cloned := original.Clone()

	if !reflect.DeepEqual(cloned.data, original.data) {
		t.Errorf("Expected cloned slice to be %v, got %v", original.data, cloned.data)
	}
}

// TestClone_SliceWithDifferentTypes_ShouldReturnClonedSlice tests cloning a slice with different data types.
func TestClone_SliceWithDifferentTypes_ShouldReturnClonedSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	original := &Slice[Person]{data: []Person{{"Alice", 30}, {"Bob", 25}}}
	cloned := original.Clone()

	if !reflect.DeepEqual(cloned.data, original.data) {
		t.Errorf("Expected cloned slice to be %v, got %v", original.data, cloned.data)
	}
}

// TestFilter_EmptySlice_ShouldReturnEmptySlice tests the Filter method with an empty slice.
func TestFilter_EmptySlice_ShouldReturnEmptySlice(t *testing.T) {
	s := &Slice[int]{data: []int{}}
	s.Filter(func(x int) bool { return x > 0 })
	if len(s.data) != 0 {
		t.Errorf("Expected empty slice, got %v", s.data)
	}
}

// TestFilter_AlwaysTruePredicate_ShouldReturnSameSlice tests the Filter method with a predicate that always returns true.
func TestFilter_AlwaysTruePredicate_ShouldReturnSameSlice(t *testing.T) {
	s := &Slice[int]{data: []int{1, 2, 3}}
	s.Filter(func(x int) bool { return true })
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestFilter_AlwaysFalsePredicate_ShouldReturnEmptySlice tests the Filter method with a predicate that always returns false.
func TestFilter_AlwaysFalsePredicate_ShouldReturnEmptySlice(t *testing.T) {
	s := &Slice[int]{data: []int{1, 2, 3}}
	s.Filter(func(x int) bool { return false })
	if len(s.data) != 0 {
		t.Errorf("Expected empty slice, got %v", s.data)
	}
}

// TestFilter_MixedPredicate_ShouldReturnFilteredSlice tests the Filter method with a predicate that returns mixed results.
func TestFilter_MixedPredicate_ShouldReturnFilteredSlice(t *testing.T) {
	s := &Slice[int]{data: []int{1, 2, 3, 4, 5}}
	s.Filter(func(x int) bool { return x%2 == 0 })
	expected := []int{2, 4}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestFilter_ValueDependentPredicate_ShouldReturnFilteredSlice tests the Filter method with a predicate that depends on element values.
func TestFilter_ValueDependentPredicate_ShouldReturnFilteredSlice(t *testing.T) {
	s := &Slice[string]{data: []string{"apple", "banana", "cherry"}}
	s.Filter(func(x string) bool { return len(x) > 5 })
	expected := []string{"banana", "cherry"}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestUnique_NoDuplicates_ShouldRemainUnchanged tests the Unique method on a slice with no duplicates.
func TestUnique_NoDuplicates_ShouldRemainUnchanged(t *testing.T) {
	s := &Slice[int]{data: []int{1, 2, 3}}
	expected := []int{1, 2, 3}
	s = s.Unique()
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got: %v", expected, s.data)
	}
}

// TestUnique_AllDuplicates_ShouldRemoveAllButOne tests the Unique method on a slice with all elements the same.
func TestUnique_AllDuplicates_ShouldRemoveAllButOne(t *testing.T) {
	s := &Slice[int]{data: []int{1, 1, 1}}
	expected := []int{1}
	s = s.Unique()
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got: %v", expected, s.data)
	}
}

// TestUnique_MixedElements_ShouldRemoveDuplicates tests the Unique method on a slice with mixed elements.
func TestUnique_MixedElements_ShouldRemoveDuplicates(t *testing.T) {
	s := &Slice[int]{data: []int{1, 2, 2, 3, 4, 4, 5}}
	expected := []int{1, 2, 3, 4, 5}
	s = s.Unique()
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got: %v", expected, s.data)
	}
}

// TestUnique_StringSlice_ShouldHandleStrings tests the Unique method on a slice of strings.
func TestUnique_StringSlice_ShouldHandleStrings(t *testing.T) {
	s := &Slice[string]{data: []string{"a", "b", "a", "c"}}
	expected := []string{"a", "b", "c"}
	s = s.Unique()
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got: %v", expected, s.data)
	}
}

// TestReduce_EmptySlice_ReturnsZeroValue tests the Reduce method with an empty slice.
func TestReduce_EmptySlice_ReturnsZeroValue(t *testing.T) {
	s := Slice[int]{data: []int{}}
	result := s.Reduce(func(x, y int) int { return x + y })
	if result != 0 {
		t.Errorf("Expected zero value, got %d", result)
	}
}

// TestReduce_SingleElement_ReturnsElement tests the Reduce method with a single element.
func TestReduce_SingleElement_ReturnsElement(t *testing.T) {
	s := Slice[int]{data: []int{5}}
	result := s.Reduce(func(x, y int) int { return x + y })
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

// TestReduce_MultipleElements_SumsCorrectly tests the Reduce method with multiple elements.
func TestReduce_MultipleElements_SumsCorrectly(t *testing.T) {
	s := Slice[int]{data: []int{1, 2, 3, 4}}
	result := s.Reduce(func(x, y int) int { return x + y })
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}

// TestReduce_MultipleElements_ProductCorrectly tests the Reduce method with a product function.
func TestReduce_MultipleElements_ProductCorrectly(t *testing.T) {
	s := Slice[int]{data: []int{1, 2, 3, 4}}
	result := s.Reduce(func(x, y int) int { return x * y })
	if result != 24 {
		t.Errorf("Expected 24, got %d", result)
	}
}

// TestReduce_MultipleElements_MaxCorrectly tests the Reduce method with a max function.
func TestReduce_MultipleElements_MaxCorrectly(t *testing.T) {
	s := Slice[int]{data: []int{1, 5, 3, 4}}
	result := s.Reduce(func(x, y int) int {
		if x > y {
			return x
		}
		return y
	})
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}