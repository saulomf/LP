SUGESTÃO DE ESPEC

A espec é a seguinte:

Pensando numa placa de petri

Teremos um mapa 10x10 para prototipação

Cada ponto do mapa terá os seguintes atributos, que podem variar cada qual de 1 a 10:
Quantidade de água
Calor
Quantidade de luz
Quantidade de ar

Sortearemos com 30% de chance quais pixels terão vida de início

O tipo especie tem os seguintes atributos que podem variar de 1 a 10:
nome
Necessidade de água
Necessidade de calor
Necessidade de gases
Necessidade de luz
1 indicando pouca necessidade e 10 indicando muita

Entao cada espécie é determinada pela combinação das necessidades desses elementos.

Os individuos terao os seguintes atributos, variando de 1 a 10 cada:
Especie
Variabilidade de agua
Variabilidade de calor
Variabilidade de gases
Variabilidade de luz
Vetor variabilidade de 6 elementos

A especie indicará a qual espécie esse individuo pertence.
Os atributos de variabilidade de água, calor, gases e luz serão modificados geracionalmente.
Se na geração anterior tinha pouca disponibilidade de água entao o atributo Variabilidade de agua do filho daquele
individuo tera um decrescimo de -1.
No momento que os atributos de variabilidade chegarem a 0 ou a 10, os atributos Necessidade serao incrementados ou decrementados
e isso configurará como uma mutação, portanto uma nova especie sera registrada no catalogo de especies.
O vetor de variabilidade pode ser substituído por um int e indicará a ID de cada indivíduo.

A disponibilidade de recursos do mapa sera regida de acordo com as necessidades de recursos dos individuos ao redor de cada pixel.
Se ha muita necessidade, na proxima geração a disponibilidade sera decrementada e vice-versa.

Assim a sobrevivencia das especies é regulada pela disponibilidade de recursos e pela sua adaptabilidade para sobreviver na presença ou
ausência deles.

Os critérios para sobrevivência e extinção ainda não foram definidos porque o programa ainda nao permite testes.
