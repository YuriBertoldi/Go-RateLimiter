# Rate Limiter

Este projeto consiste na criação de um limitador de taxa em Go, projetado para controlar o número máximo de requisições por segundo com base em endereços IP específicos ou tokens de acesso. O objetivo é gerenciar de maneira eficaz o tráfego de serviços web.

## Principais Recursos

- **Restrição por Endereço IP**: Restringe as requisições de um único endereço IP dentro de um intervalo de tempo definido.
- **Restrição por Token de Acesso**: Limita as requisições com base em tokens de acesso exclusivos, permitindo diferentes limites de tempo de expiração para cada token.
- **Prioridade de Configurações**: As configurações do token de acesso têm precedência sobre as configurações de endereço IP.
- **Integração com Middleware**: Funciona como um middleware integrado ao servidor web.
- **Definição de Limite de Requisições**: Permite especificar o número máximo de solicitações por intervalo de tempo.
- **Personalização da Duração do Bloqueio**: Define o período de bloqueio para IP ou Token após ultrapassar os limites de solicitação.
- **Configurações via Variáveis de Ambiente**: As configurações podem ser configuradas por meio de variáveis de ambiente ou um arquivo .env.
- **Resposta HTTP ao Exceder o Limite**: Retorna o código HTTP 429 e uma mensagem indicando a ultrapassagem do número máximo de solicitações.
- **Utilização do Banco de Dados Redis**: Utiliza Redis para armazenar e consultar informações do limitador.
- **Flexibilidade na Persistência**: Adota uma estratégia flexível para alternar entre o Redis e outros mecanismos de persistência.
- **Lógica do Limitador Independente**: A lógica do limitador é separada do middleware.
- **Configurações Personalizadas de Quantidade e Duração de Bloqueio**: Permite configurar limites e durações de bloqueio personalizados para IPs ou Tokens específicos.

## Variáveis de Configuração

As variáveis de ambiente desempenham um papel crucial na configuração e personalização do comportamento do limitador de taxa:

- **REDIS_ADDRESS**: Define o endereço do servidor Redis utilizado pelo limitador de taxa.
- **REDIS_PASSWORD**: Senha para autenticação no servidor Redis.
- **REDIS_DB**: Número do banco de dados Redis a ser utilizado pelo aplicativo.
- **DEFAULT_IP_MAX_REQ_PER_SEC**: Define o limite padrão de solicitações por segundo por endereço IP.
- **DEFAULT_TOKEN_MAX_REQ_PER_SEC**: Estabelece o limite padrão de solicitações por segundo por token de acesso.
- **DEFAULT_IP_BLOCK_DURATION**: Duração do bloqueio padrão para um endereço IP que excede seu limite de solicitações.
- **DEFAULT_TOKEN_BLOCK_DURATION**: Duração do bloqueio padrão para um token que excede seu limite de solicitações.
- **CUSTOM_MAX_REQ_PER_SEC**: Permite a definição de limites de solicitação personalizados para IPs ou tokens específicos.
- **CUSTOM_BLOCK_DURATION**: Configura durações de bloqueio personalizadas para IPs ou tokens específicos.

Essas variáveis são fundamentais para a flexibilidade e eficácia do limitador, permitindo sua adaptação às necessidades específicas do sistema.

## Funcionamento do Limitador de Requisições

O limitador de taxa é implementado como um middleware no servidor HTTP, permitindo que ele intercepte e controle as solicitações.

### Middleware de Limitação de Requisições

- O middleware `RateLimiterMiddleware` é aplicado a cada solicitação recebida pelo servidor.
- Ele identifica cada solicitação por um identificador único, que pode ser um token de acesso (se presente no cabeçalho API_KEY) ou o endereço IP do solicitante.
- Após identificar a solicitação, o middleware consulta o `RateLimit`, uma estrutura que contém as configurações de limitação e o Redis para verificar se o limite foi excedido.

### Verificação e Controle de Limite

- A função `IsLimitExceeded` da estrutura `RateLimit` é responsável por determinar se uma solicitação excede o limite configurado.
- Ela verifica o número atual de solicitações feitas pelo identificador (IP ou token) no Redis.
- Com base no identificador, a função determina o limite máximo de solicitações permitidas por segundo.

## Exemplos de Utilização

- **Limitação por IP**: Se configurado para um máximo de 5 solicitações por segundo por IP, a 6ª solicitação do IP 192.168.1.1 dentro de um segundo será bloqueada.
- **Limitação por Token**: Se um token `abc123` estiver configurado com um limite de 10 solicitações por segundo, a 11ª solicitação dentro desse segundo será bloqueada.
- **Tempo de Expiração**: Após atingir o limite, novas solicitações do mesmo IP ou token só serão possíveis após o tempo de expiração.
- **Configuração Personalizada por IP ou Token**: Permite definir limites e durações de bloqueio específicos.
