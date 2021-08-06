Uma aplicação para construir listas

* [...] + limite de 5 listas por usuario e 100 musicas por lista
* Favoritos (required)
* Recentes (limit:50, required)
* Quero aprender (required): iduser, idsong,
* Consegui tocar (required)
* Não Consegui tocar (required)
* Hard-core (required)
* Easy (required)

Extras:
* musica só pode existir em Hard-core ou Easy, nunca nas 2 ao mesmo tempo
* musica só pode existir em Consegui tocar ou Não Consegui tocar, nunca nas 2 ao mesmo tempo




## SOLID:
* **S**ingle Responsibility Principle (Princípio da responsabilidade Única)
> Uma classe deve ter um, e somente um, motivo para mudar. (Uncle Bob)

* **O**pen/Closed Principle (Princípio aberto/fechado)
> Objetos ou entidades devem estar abertos para extensão, mas fechados para modificação

* **L**iskov Substitution Principle (Princípio substituição de Liskov)
> Se S é um subtipo de T, então os objetos do tipo T, em um programa,
 podem ser substituídos pelos objetos de tipo S sem que seja necessário
 alterar as propriedades deste programa.

> “Classes derivadas devem poder ser substituídas por suas classes base” (Robert C. Martin, 1996)

* **I**nterface Segregation Principle (Princípio de segregação de interfaces)
> “Os clientes não devem ser forçados a depender de interfaces
> que eles não usam”. (Robert C. Martin,1996)

> Uma classe não deve ser forçada a implementar interfaces
> e métodos que não irão utilizar.

* **D**ependecy Inversion Principle (Princípio da inversão de dependências)
> "Módulos de alto nível não devem depender de módulos de baixo nível.
> Ambos devem depender de abstrações." (Robert C. Martin, 1996)

> "Abstrações não devem depender de detalhes;
> Detalhes devem depender de abstrações." (Robert C. Martin, 1996)

> Programe voltado para interface e não para implementação.