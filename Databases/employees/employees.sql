-- Selecionar o nome, a posição e a localização dos departamentos onde os vendedores trabalham.
select nome, sobrenome, posto from funcionario f where posto = 'Vendedor';

-- -  Mostra os departamentos com mais de cinco empregados.

SELECT * 
FROM departamento d
INNER JOIN funcionario f on d.depto_nro = f.depto_nro
WHERE (SELECT COUNT(DEPTO_NRO) FROM departamento GROUP BY depto_nro )>5

-- -  Mostra o nome, o salário e o nome do departamento dos empregados que têm a mesma posição que "Mito Barchuk".

SELECT nome, salario, posto, depto_nome  FROM funcionario f
INNER JOIN departamento d on f.depto_nro = d.depto_nro
WHERE posto = (SELECT posto FROM funcionario WHERE nome = 'Mito' AND sobrenome = 'Barchuk');

-- -  Mostra os detalhes dos empregados que trabalham no departamento de contabilidade, ordenados por nome.

SELECT * FROM funcionario f
INNER JOIN departamento d on f.depto_nro = d.depto_nro
WHERE depto_nome = 'Contabilidade' ORDER BY nome;

-- -  Mostra o nome do empregado com o salário mais baixo.

SELECT nome, salario FROM funcionario ORDER BY salario ASC LIMIT 1;

-- -  Mostra os detalhes do empregado com o salário mais alto no departamento de "Vendas".
SELECT * FROM funcionario f
INNER JOIN departamento d on f.depto_nro = d.depto_nro
WHERE depto_nome = 'Vendas' ORDER BY salario DESC LIMIT 1;
