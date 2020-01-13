# Exportador File2BD 

Olá, esse projeto desenvolvido como avaliação consite em um programa escrito na linguagem GO que exporta um arquivo de texto para um banco de dados PostgreSQL e efetua a validação de alguns dados antes da inserção através da própria linguagem e outra higienização de dados após a inserção dos dados no banco através do uso de procedures e triggers no banco de dados.

# Instalação e execução via docker

É necessario ter o docker e o docker compose instalado na maquina hospedeira.

  - Efetuar o download ou clone do projeto
  - Acessar o diretório da raiz do projeto
  - Executar o comando "docker-compose build"
  - Executar o comando "docker-compose up"
  - Devido ao programa da aplicação ainda não esperar o banco de estar aceitando conexões é preciso executar mais uma vez o comando "docker-compose up" 
  - Acessar a maquina do banco de dados ou abrir uma conexão com o banco de dados via PgAdmin4 por exemplo e fazer uma consulta na tabela "dadosclientes"
  
É possível verificar na tabela do banco de dados após essa execução que os dados foram inseridos na tabela do banco de dados e que alguns campos foram higinenizados, os campos que contem cpf, cnpj e data não contém mais os caracteres ".", "-" ou "/".

# Observações

  - O SQL para criar a tabela, funções e trigger está no diretório ./basededados/bdgo.sql
  - O código fonte do programa na linguagem GO está no diretório ./src
  - O arquivo ".env" na raiz contém as informações para conexão do banco de dados
  - A função do banco de dados também remove, caso ocorra, a remoção de caracteres especiais e coloca todos os caracteres em minúsculo. 
  - O CNPJ e CPF são validados dentro do código GO e adicionado uma coluna no banco de dados com um valor boolean com a validação do mesmo.
  

Caso tenha ficado alguma dúvida não hesite e entrar em contato. 
