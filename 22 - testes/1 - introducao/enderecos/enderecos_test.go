// Test de unidade
package enderecos_test

import (
	"fmt"
	"testing"
	."tests/enderecos"
)

type testStageOfAddress struct {
	entryAddress string
	resultAdress string
}

func TestVerifyEnderecoType(t *testing.T){
	t.Parallel() // Roda todos os métodos que tem essa função em paralelo
	listOfAddress := []testStageOfAddress{
		{"Rua dos Bolos", "Rua"},
		{"Avenida Imirim", "Avenida"},
		// {"Estrada do Moinho", "Estrada"},
		{"Rodovia de Palha", "Rodovia"},
		// {"", "Validação Inválida"},
	}

	for _, AddressObject := range listOfAddress{
		testAddress, testAddressCorrectResult := AddressObject.entryAddress, AddressObject.resultAdress
		testAddressResult := VerifyEnderecoType(testAddress)
		
		if testAddressCorrectResult != testAddressResult{
			t.Errorf("Os resultados não são iguais, Esperava (%s) e Recebeu (%s)\n", testAddressCorrectResult, testAddressResult)
		}

		fmt.Printf("Objeto testado com sucesso, esperava (%s) e recebeu (%s)\n", testAddressCorrectResult, testAddressResult)
	}
}

func TestOther(t *testing.T){
	t.Parallel()
	if 1 > 2{
		t.Error("O número fornecido está errado")
	}
}
