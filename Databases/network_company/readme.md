
##  DER e SQL

Uma empresa provedora de Internet precisa de um banco de dados para armazenar cada um de seus clientes, juntamente com o plan/pack contratado.

Por meio de uma análise prévia, sabe-se que as seguintes informações devem ser armazenadas:

Para clientes: número de identificação, nome, sobrenome, data de nascimento, província, cidade.
Para planos de Internet: identificação do plano, velocidade oferecida em megabytes, preço, desconto.
 
# Exercício 1

Após a definição dos requisitos da empresa, é solicitado que eles sejam modelados por meio de um DER (Diagrama Entidade-Relacionamento).

# Exercício 2 

Depois que o banco de dados tiver sido modelado e apresentado, responda às seguintes perguntas:

a. Qual é a primary key da tabela de clientes? Justifique sua resposta

- R:  Número de identificação do cliente, por ser uma chave única.

b. Qual é a primary key da tabela de planos de Internet? Justifique a resposta.
 
- R: Identificação do plano, por ser uma chave única.

c. Como seriam os relacionamentos entre as tabelas? Em qual tabela deve haver uma foreign key? A qual campo de qual tabela essa foreign key se refere? Justifique a resposta.

- R: 1-1, o cliente só pode ter um único plano de internet e o plano de internet só pode ter um cliente vinculado.

# Exercício 3

Depois que o diagrama tiver sido criado e essas perguntas tiverem sido respondidas, use o PHPMyAdmin ou o MySQL Workbench para executar o seguinte:

É solicitado que você crie um novo banco de dados chamado "empresa_internet".
Incorpore 10 registros na tabela de clientes e 5 na tabela de planos de Internet.
Faça as associações/relacionamentos correspondentes entre esses registros.

# Exercício 4

Indique 10 consultas SQL que poderiam ser feitas no banco de dados. Expresse as instruções.


