-- - Enumera os dados dos autores.

SELECT 
  ROW_NUMBER() OVER (ORDER BY nombre) AS id,
  nombre AS nombre
FROM autor;


-- - Indica o nome e a idade dos alunos

SELECT nombre, edad FROM estudiante;

-- - Que alunos pertencem ao curso de informática?

SELECT nombre FROM estudiante WHERE carrera = 'informática';

-- - Quais são os autores de nacionalidade francesa ou italiana?

SELECT nombre FROM autor WHERE nacionalidad = 'francesa' OR nacionalidad = 'italiana';

-- - Quais os livros que não são da área da Internet?

SELECT titulo FROM libro WHERE area != 'Internet';

-- - Enumera os livros publicados pela Salamandra.

SELECT 
ROW_NUMBER() OVER (ORDER BY titulo) as id, FROM libro WHERE editorial = 'Salamandra';

-- - Enumera os nomes dos alunos cuja idade é superior à média.

SELECT nombre FROM estudiante WHERE edad > (SELECT AVG(edad) FROM estudiante);

-- - Enumera os nomes dos alunos cujo apelido começa com a letra G.

SELECT nombre FROM estudiante WHERE apellido LIKE 'G%';

-- - Faz uma lista dos autores do livro "O Universo: Guia de Viagem". (Apenas os nomes devem ser indicados).

SELECT nombre FROM autor WHERE idAutor IN (SELECT idAutor FROM libroAutor WHERE idLibro =
 (SELECT idLibro FROM libro WHERE titulo = 'O Universo: Guia de Viagem'));

-- - Que livros emprestaste ao leitor "Filippo Galli"?

SELECT titulo FROM libro WHERE idLibro IN (SELECT idLibro FROM prestamo WHERE idLector =
 (SELECT idLector FROM lector WHERE nombre = 'Filippo Galli'));

-- - Indica o nome do aluno mais novo.

SELECT nombre FROM estudiante WHERE edad = (SELECT MIN(edad) FROM estudiante);

-- - Enumera os nomes dos alunos a quem foram emprestados livros da Base de Dados.

SELECT nombre FROM estudiante WHERE idEstudiante IN (SELECT idEstudiante FROM prestamo WHERE idLibro IN
 (SELECT idLibro FROM libro WHERE area = 'Base de Dados'));

-- - Enumera os livros que pertencem à autora J.K. Rowling.

SELECT titulo FROM libro WHERE idLibro IN (SELECT idLibro FROM libroAutor WHERE idAutor =
 (SELECT idAutor FROM autor WHERE nombre = 'J.K. Rowling'));

-- - Enumera os títulos dos livros que deviam ser devolvidos em 16/07/2021.

SELECT titulo FROM libro WHERE idLibro IN (SELECT idLibro FROM prestamo WHERE fechaDevolucion = '2021-07-16');