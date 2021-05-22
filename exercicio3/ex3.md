**Exercício 3:** Digamos que a gente quisesse criar um simple-server-2 a partir da nossa Dockerfile original. 

Tenha em mente que uma imagem simple-server:1.4 já existe construída na minha maquina local e eu quero só mudar o comando que ela roda. Digamos que agora quero que o comando seja 

```shell
echo "rodando:" && go run .
```

Mas temos um problema: A imagem golang foi apagada e eu estou sem internet. 

É possível eu criar uma nova imagem mudando somente o comando sem a imagem golang presente em meu computador mas com a imagem simple-server:1.4 presente? Como?