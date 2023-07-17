Este é um programa em Go que é usado para encontrar uma senha específica em um arquivo, dada uma hash e um salt. Aqui está uma visão geral do que cada parte do código faz:

O programa primeiro solicita ao usuário que forneça o caminho completo para um arquivo. Este arquivo deve conter uma lista de senhas, uma por linha.

Em seguida, o programa solicita ao usuário que forneça uma hash e um salt. A hash é uma representação criptográfica de uma senha que foi "salgada", ou seja, teve o salt adicionado a ela antes de ser hasheada.

O programa então abre o arquivo fornecido e começa a ler as senhas uma por uma.

Para cada senha, o programa adiciona o salt fornecido, cria uma hash SHA-512 da senha salgada e compara a hash resultante com a hash fornecida pelo usuário.

Se a hash da senha salgada corresponder à hash fornecida pelo usuário, o programa imprime a senha e termina.

Se o programa chegar ao final do arquivo sem encontrar uma correspondência, ele imprime uma mensagem informando que a senha não foi encontrada.

Se ocorrer um erro ao ler o arquivo, o programa imprime uma mensagem de erro.

Este programa é um exemplo de um "ataque de força bruta" ou "ataque de dicionário" em criptografia, onde todas as possíveis senhas em um conjunto conhecido (neste caso, as senhas no arquivo fornecido) são testadas para ver se alguma delas corresponde à hash fornecida.

Forma de uso : 

go run quebrar_shadow.go 
