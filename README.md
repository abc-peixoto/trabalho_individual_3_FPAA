# trabalho_individual_3_FPAA

# Estrutura proposta do repositório

```
hamiltonian-path-go/
├─ cmd/
│  └─ hamilpath/
│     └─ main.go              # CLI para carregar grafo, rodar o algoritmo e imprimir o caminho
├─ internal/
│  ├─ graph/
│  │  └─ graph.go             # Estruturas e utilitários de grafo (dirigido/não-dirigido)
│  └─ hamil/
│     └─ hamil.go             # Backtracking para Caminho Hamiltoniano (com/sem vértice inicial)
├─ examples/
│  ├─ undirected_edges.txt    # Exemplo de grafo não-dirigido (lista de arestas)
│  ├─ directed_edges.txt      # Exemplo de grafo dirigido (lista de arestas)
│  └─ expected_path.txt       # Exemplo de caminho encontrado (um por linha)
├─ assets/
│  └─ sample.png              # (Opcional) export gerado pelo view.py
├─ view.py                    # (Opcional, Ponto extra) Visualização com NetworkX/Matplotlib
├─ README.md                  # Documentação completa + relatório técnico
└─ go.mod                     # Módulo Go
```

## Descrição do projeto

Implementamos um algoritmo de **backtracking** para determinar se existe um **Caminho Hamiltoniano** em um grafo (dirigido ou não-dirigido). Se existir, o programa imprime um caminho válido (sequência de vértices que visita **cada vértice exatamente uma vez**).

### Lógica geral do algoritmo

* O problema é intrinsecamente difícil (decisão é **NP-completo**); portanto, adotamos **backtracking** com poda simples.
* Tentamos construir o caminho incrementalmente:

  1. Escolhemos um vértice inicial (ou testamos todos, caso não informado).
  2. Mantemos um array `path` de tamanho `n` e um conjunto/array de visitados.
  3. A cada passo `pos`, tentamos adicionar um próximo vértice `v` **não visitado** que seja

     * adjacente ao último vértice do caminho (em grafo dirigido: aresta **na direção correta**),
     * e respeite eventuais restrições de início.
  4. Se `pos == n`, encontramos um caminho Hamiltoniano.
  5. Caso não haja candidato válido, retrocedemos (**backtrack**).

### Decisões de projeto

* Suporte a **grafos dirigidos e não-dirigidos** via `Graph{Directed bool}`.
* Entrada por **lista de arestas** (mais simples e clara para testes). Também inferimos o número de vértices pelo maior índice visto, permitindo vértices `0..n-1`.
* A CLI aceita:

  * `-file`: caminho do arquivo com arestas (uma aresta por linha, separadas por espaço; ex.: `0 1`).
  * `-directed`: se presente, trata o grafo como **dirigido**.
  * `-start`: opcional, fixa o vértice inicial.
* Saída:

  * `FOUND` + o caminho (ex.: `0 3 1 2`).
  * Ou `NOT FOUND` se inexistente.

---

## Como executar o projeto

### Pré-requisitos

* **Go 1.21+**
* (Opcional ponto extra) **Python 3.10+**, `networkx`, `matplotlib`

### Clonar e rodar

```bash
# Clonar
git clone https://github.com/abc-peixoto/trabalho_individual_3_FPAA.git
cd hamiltonian-path-go

go run ./cmd/hamilpath \
  -file ./examples/undirected_edges.txt

# Grafo dirigido
go run ./cmd/hamilpath \
  -file ./examples/directed_edges.txt -directed

# Com vértice inicial fixo (ex.: 0)
go run ./cmd/hamilpath \
  -file ./examples/undirected_edges.txt -start 0
```

### Formato dos arquivos de arestas

* Uma aresta por linha.
* Vértices identificados por inteiros `0..n-1`.
* Separação por espaço. Ex.: `0 2` significa aresta (0,2). Em **grafo não-dirigido**, adicionamos (2,0) automaticamente.

## Relatório técnico

### Classes de complexidade: P, NP, NP-Completo, NP-Difícil

* **Definições:**

  * **P**: problemas de decisão solucionáveis em tempo polinomial por uma máquina determinística.
  * **NP**: problemas de decisão cujas soluções podem ser **verificadas** em tempo polinomial.
  * **NP-Completo**: problemas em NP para os quais **todo** problema em NP reduz-se em tempo polinomial (são os mais “difíceis” de NP).
  * **NP-Difícil**: pelo menos tão difíceis quanto os NP-Completo; podem **não** pertencer a NP (ex.: versões de otimização).

* **Caminho Hamiltoniano (decisão)**: dado um grafo **G**, existe um caminho que visita cada vértice **uma única vez**? → Este problema é **NP-Completo** (para grafos dirigidos e não-dirigidos). A verificação é polinomial (basta checar adjacências e unicidade em O(n)). A completude decorre de reduções clássicas (e relação estreita com **TSP** decisão).

* **Relação com TSP**: o **TSP (decisão)** pergunta se existe um ciclo Hamiltoniano com custo ≤ K (em grafo completo com pesos). **Hamiltoniano** é um caso não-ponderado/estrutural; muitos livros mostram reduções entre TSP e Hamiltoniano. Já o **TSP de otimização** é **NP-Difícil**.

### Complexidade assintótica de tempo do backtracking

* No pior caso, o algoritmo tenta permutar vértices: ~ **O(n! · n)** (ou **O(n!)**) — fator `n` decorre da checagem de adjacência para cada inclusão.
* Com verificação/visitados em arrays e matriz de adjacência, cada teste é O(1), mas o **n!** domina.
* **Como determinamos**: por **contagem de estados** explorados. Cada posição do caminho escolhe entre os vértices ainda não visitados → `n · (n-1) · (n-2) · … · 1 = n!`. Poda por adjacência reduz instâncias práticas, mas não muda a ordem assintótica.

### Aplicação do **Teorema Mestre**

* **Não aplicável**: o Teorema Mestre analisa recorrências de **divisão e conquista** do tipo `T(n) = a T(n/b) + f(n)`. Nosso backtracking **não** divide o problema em subproblemas proporcionais de tamanho `n/b`. Logo, justificadamente **não usamos** o Teorema Mestre.

### Casos: melhor, médio e pior

* **Melhor caso**: encontramos um caminho rapidamente (ex.: primeiro ramo já é válido) → **O(n)** expansões (linear nas posições) + O(n) para checagens → assintoticamente **O(n)**.
* **Caso médio**: depende da estrutura do grafo e da heurística de ordem dos vizinhos; em geral **exponencial** (mas menor que `n!` com podas).
* **Pior caso**: nenhum caminho existe (ou está “profundo” na árvore), exploramos quase toda a árvore de busca → **O(n! · n)**.

---

## Explicação linha a linha (núcleo do algoritmo)

### `hamil.go`

* `HamiltonianPath(g *graph.Graph, start int)`

  * Aloca `visited` e `path`.
  * Função local `tryFrom(s)` reinicializa visitados e tenta `dfs` a partir de `s`.
  * Se `start >= 0`, testamos **apenas** esse vértice; senão, iteramos por todos os vértices.
* `dfs(g, pos, visited, path)`

  * Se `pos == g.N` → **sucesso** (construímos `n` vértices no caminho).
  * `last := path[pos-1]` é o último vértice do prefixo.
  * Iteramos `v` de `0..n-1`:

    * pulamos se `visited[v]`.
    * checamos adjacência `g.Adj[last][v]` (respeitando direção quando `Directed==true`).
    * marcamos `visited[v]=true`, gravamos em `path[pos]=v` e chamamos recursivamente.
    * se falhar, **backtrack**: `visited[v]=false`.

### `graph.go`

* `New(n, directed)`: cria matriz de adjacência `n×n`.
* `AddEdge(u,v)`: adiciona aresta; se não-dirigido, também adiciona `(v,u)`.
* `LoadEdgeList`: varre o arquivo, guarda arestas, infere `n` pelo maior índice, instancia grafo e insere arestas.

---

## Testes rápidos

```bash
# Não-dirigido — deve existir caminho 0-1-2-3 (ou permutações)
go run ./cmd/hamilpath -file ./examples/undirected_edges.txt

# Dirigido — deve existir 0->1->2->3
go run ./cmd/hamilpath -file ./examples/directed_edges.txt -directed

# Fixando início em 0
go run ./cmd/hamilpath -file ./examples/undirected_edges.txt -start 0
```