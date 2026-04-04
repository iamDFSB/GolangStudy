package shape_methods_test

import (
	"shapes/shape_methods"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func (t * testing.T){
		rectangle := shape_methods.Rectangle{2.0, 2.0}
		result := rectangle.Area()
		if result != 4.0{
			t.Fatalf("O cálculo da área do Retangulo não gerou o resultado correto. (%f)", result)
		}
	})

	t.Run("Circulo", func (t *testing.T){
		rectangle := shape_methods.Circle{5.0}
		result := rectangle.Area()
		if result <= 78.539816{
			t.Errorf("O cálculo da área do Circulo não gerou o resultado correto. (%f)", result)
		}
	})
}

