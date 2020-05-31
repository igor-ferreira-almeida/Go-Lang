# Go-Lang

## Testes 

- Arquivos de testes devem estar no mesmo pacote;
- Devem possuir o mesmo nome do arquivo que será testado com o sufixo "_test.go" (Exemplo: triangle_test.go);
- Métodos por convensão são iniciados com o prefixo "Test" (Exemplo: func TestCalculateArea());.

```
go test
go test -cover
go tool cover -func=resultado.out
go tool cover -html=resultado.out
```
