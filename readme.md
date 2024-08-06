# BASE STRUCTURE GOLANG
 
 Este é um projeto de backend em Go utilizando o framework Gin, bibliotecas de injeção de dependência (DI) e logging. O projeto é estruturado para seguir os princípios de arquitetura modular e Clean Architecture, facilitando a manutenção e a escalabilidade.

 ## Descrição
Este projeto é um backend para uma aplicação genérica, estruturado para demonstrar o uso de algumas bibliotecas populares em Go e boas práticas de desenvolvimento. Ele inclui funcionalidades básicas como logging, configuração, e um exemplo de API.

 ## Bibliotecas Utilizadas
- **Gin:** Um framework web para Go que oferece uma maneira rápida e fácil de construir APIs HTTP.
- **Swaggo:** Ferramenta para geração automática de documentação Swagger para APIs Go.
- **Dig:** Biblioteca para Injeção de Dependência (DI) em Go, que ajuda a gerenciar a criação e injeção de dependências.
- **Zap:** Biblioteca de logging da Uber, conhecida por seu desempenho e flexibilidade.
- **Viper:** Biblioteca para gerenciamento de configuração, que suporta várias fontes de configuração, como arquivos YAML e variáveis de ambiente.
## Estrutura do Projeto
A estrutura do projeto é organizada da seguinte forma:

```
backend
 /cmd
  /server
   -main.go                 # Ponto de entrada da aplicação. Inicializa o contêiner DI e configura o servidor. 
 /config
  -local.yml               # Arquivo de configuração local (YAML).
 /database
  -database.go             # Configuração e conexão com o banco de dados.
 /di
  -config_di.go            # Configuração do contêiner DI, registro de dependências.
 /middleware
  -cors.go                 # Configuração do middleware de CORS.
  -jwt.go                  # Middleware para autenticação JWT.
  -log.go                  # Configuração do middleware de logging.
 /modules
  /example
   /application
    -business.go           # Lógica de negócios do módulo exemplo.
    -model.go              # Modelos de dados do módulo exemplo.
   /domain
    -interface.go          # Interfaces para repositórios e serviços.
    -model.go              # Modelos de domínio.
    -repository.go         # Repositórios para acesso a dados.
    -service.go            # Serviços que implementam a lógica de negócios.
   /infrastructure
    -data.go               # Implementação de acesso a dados.
    -model.go              # Modelos específicos da infraestrutura.
   /interfaces
    -handler.go            # Handlers HTTP para o módulo exemplo.
    -handler_test.go       # Testes para os handlers.
  -example_module.go       # Arquivo principal para o módulo exemplo.
 /pkg
  /config
   -config.go              # Funções para carregar e gerenciar a configuração.
   -types.go               # Tipos e estruturas relacionados à configuração.
  /dig
   -config_dig.go          # Configurações adicionais para o contêiner DI.
  /http
   /responses
    -error.go              # Definições de respostas de erro.
    -v1.go                 # Respostas para a versão 1 da API.
   /server
    -server.go             # Configuração do servidor HTTP.
 /utils
  -convert.go              # Funções utilitárias de conversão.
  -pointer.go              # Funções utilitárias para manipulação de ponteiros.
```
## Configuração e Inicialização

### Instalação de Dependências

Execute o seguinte comando para instalar as dependências do projeto:

```bash Copiar código
  go mod tidy 
```
Configure as variáveis de ambiente e o arquivo local.yml para ajustar as configurações do projeto, como detalhes do banco de dados e opções de logging.

### Inicialização

Para iniciar o servidor, execute os seguintes comandos:

 Adicionar o caminho do arquivo com as variaveis de ambiente nas suas configurações  
 ```bash Copiar código
  APP_CONF=backend/config/local.yml
 ```
 
### Init Swag
```bash Copiar código
  swag init  -g cmd/server/main.go -o ./docs --parseDependency`
``` 

```bash Copiar código
  go run ./cmd/server/main.go
```
**Endpoints**\
A API inclui endpoints definidos no módulo example. Utilize o Swagger UI para explorar e testar os endpoints. A documentação Swagger pode ser acessada em /swagger/index.html.

**Testes**\
Os testes para os handlers e outras funcionalidades estão localizados em modules/example/interfaces/handler_test.go. Execute os testes usando:

```bash Copiar código
go test ./...
```
**Contribuição**\
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests. Por favor, siga as diretrizes de contribuição padrão.
