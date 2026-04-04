package enderecos

import (
	"strings"
)

func VerifyEnderecoType(endereco string) string{
	enderecoTypes := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoLowerCase := strings.ToLower(endereco)
	firstEnderecoWord := strings.Split(enderecoLowerCase, " ")[0]

	validationResult := false

	for _, text := range enderecoTypes{
		if text == firstEnderecoWord{
			validationResult = true
		}
	}

	if !validationResult{
		return "Validação Inválida"
	}

	return strings.Title(firstEnderecoWord)

}

// go test -v  -> "Mostra mais informações do que o normal"
// go test ./... -> "Executa todos os arquivos de test que encontrar"
// go test --cover -> "Verifica o percentual de cobertura do código de teste"
// go test --coverprofile cover.txt -> "Cria um arquivo de relatório da cobertura"
// go tool cover --func=cover.txt -> "Pega o arquivo de cobertura e mostra no terminal as informações"
// go tool cover --html=cover.txt -> "Pega as informações do cover.txt e mostra em um arquivo html temporário"
