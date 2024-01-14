package main

import "testing"

func TestReverseSign(t *testing.T) {
	intVal := 42
	int8Val := int8(42)
	int16Val := int16(42)
	int32Val := int32(42)
	int64Val := int64(42)
	float32Val := float32(42.42)
	float64Val := 42.42

	// int
	t.Run("reverse sign for int using multiplication", func(t *testing.T) {
		res := ReverseSignMultiplication(intVal)
		if res != -intVal {
			t.Errorf("values are not equal: %v, %v", res, -intVal)
		}

	})
	// int8
	t.Run("reverse sign for int8 using multiplication", func(t *testing.T) {
		res := ReverseSignMultiplication(int8Val)
		if res != -int8Val {
			t.Errorf("values are not equal: %v, %v", res, -int8Val)
		}
	})
	// int16
	t.Run("reverse sign for int16 using multiplication", func(t *testing.T) {
		res := ReverseSignMultiplication(int16Val)
		if res != -int16Val {
			t.Errorf("values are not equal: %v, %v", res, -int16Val)
		}
	})
	// int32
	t.Run("reverse sign for int32 using multiplication", func(t *testing.T) {
		res := ReverseSignMultiplication(int32Val)
		if res != -int32Val {
			t.Errorf("values are not equal: %v, %v", res, -int32Val)
		}
	})
	// int64
	t.Run("reverse sign for int64 using multiplication", func(t *testing.T) {
		res := ReverseSignMultiplication(int64Val)
		if res != -int64Val {
			t.Errorf("values are not equal: %v, %v", res, -int64Val)
		}
	})
	// float32
	t.Run("reverse sign for float32 using multiplication", func(t *testing.T) {
		res := ReverseSignMultiplication(float32Val)
		if res != -float32Val {
			t.Errorf("values are not equal: %v, %v", res, -float32Val)
		}
	})
	// float64
	t.Run("reverse sign for float64 using multiplication", func(t *testing.T) {
		res := ReverseSignMultiplication(float64Val)
		if res != -float64Val {
			t.Errorf("values are not equal: %v, %v", res, -float64Val)
		}
	})
}

func BenchmarkReverseSignMultiplication(b *testing.B) {
	intVal := 42
	int8Val := int8(42)
	int16Val := int16(42)
	int32Val := int32(42)
	int64Val := int64(42)
	float32Val := float32(42.42)
	float64Val := float64(42.42)

	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMultiplication(intVal)
		}
	})

	b.Run("int8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMultiplication(int8Val)
		}
	})

	b.Run("int16", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMultiplication(int16Val)
		}
	})

	b.Run("int32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMultiplication(int32Val)
		}
	})

	b.Run("int64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMultiplication(int64Val)
		}
	})

	b.Run("float32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMultiplication(float32Val)
		}
	})

	b.Run("float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMultiplication(float64Val)
		}
	})
}

func BenchmarkReverseSignMinus(b *testing.B) {
	intVal := 42
	int8Val := int8(42)
	int16Val := int16(42)
	int32Val := int32(42)
	int64Val := int64(42)
	float32Val := float32(42.42)
	float64Val := float64(42.42)

	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMinus(intVal)
		}
	})

	b.Run("int8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMinus(int8Val)
		}
	})

	b.Run("int16", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMinus(int16Val)
		}
	})

	b.Run("int32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMinus(int32Val)
		}
	})

	b.Run("int64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMinus(int64Val)
		}
	})

	b.Run("float32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMinus(float32Val)
		}
	})

	b.Run("float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignMinus(float64Val)
		}
	})
}

// BenchmarkReverseSignChaos can only test for floats because math.abs() only
// works with float64 in go.
//
// It is possible to convert to float64 for other types but that would not be
// a fair and equivalent benchmark
func BenchmarkReverseSignChaos(b *testing.B) {
	var float64Val float64
	float64Val = 42.42

	b.Run("float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseSignChaos(float64Val)
		}
	})
}
