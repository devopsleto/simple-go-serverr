**Exercício 4:** Digamos que eu não queira instalar o go no meu container mas ainda rodar meu servidor.
Tente pensar no que é possível fazer para não precisar instalar o go.
<details>
<summary>Abra aqui para ver um raciocínio possível e ler o restante do exercício</summary>
Podemos compilar o servidor para o sistema operacional do container (ubuntu:18.04, nesse caso) da seguinte forma:
```
GOOS=linux GOARCH=amd64 go build
```
isso vai gerar um arquivo binário simple-go-server na pasta do projeto.

Agora, eu posso usar o ubuntu como imagem base do meu container e não preciso instalar o go.
Faça essa Dockerfile e verifique que o server continua funcionando.
</details>