# Rate limiter

Este projeto consiste no desenvolvimento de um limitador de taxa em Go, projetado para controlar o número máximo de requisições por segundo com base em endereços IP específicos ou tokens de acesso. O objetivo é gerenciar de forma eficaz o tráfego de serviços web.

## Recursos Principais

- **Limitação por Endereço IP**: Restringe as requisições de um único endereço IP dentro de um intervalo de tempo definido.
- **Limitação por Token de Acesso**: Limita as requisições com base em tokens de acesso exclusivos, permitindo diferentes limites de tempo de expiração para cada token.
- **Configurações Prioritárias**: As configurações do token de acesso têm precedência sobre as configurações de endereço IP.
- **Integração com Middleware**: Funciona como um middleware integrado ao servidor web.
- **Configuração de Limite de Requisições**: Permite definir o número máximo de requisições por intervalo de tempo.
- **Duração do Bloqueio Personalizada**: Define o tempo de bloqueio para IP ou Token após exceder os limites de requisição.
- **Configurações de Variáveis de Ambiente**: As configurações podem ser feitas via variáveis de ambiente ou através de um arquivo .env.
- **Resposta HTTP ao Exceder o Limite**: Responde com o código HTTP 429 e uma mensagem indicando a ultrapassagem do número máximo de requisições.
- **Banco de Dados Redis**: Utiliza Redis para armazenar e consultar informações do limitador.
- **Estratégia de Persistência Flexível**: Padrão de estratégia para alternar facilmente entre o Redis e outros mecanismos de persistência.
- **Lógica do Limitador Separada**: A lógica do limitador é independente do middleware.
- **Configurações Personalizadas de Quantidade de Requisições e Duração de Bloqueio**: Permite configurar limites e durações de bloqueio personalizados para IPs ou Tokens específicos.

## Exemplos de Utilização

- **Exemplo de Limitação por IP**: Se configurado para um máximo de 5 requisições por segundo por IP, a 6ª requisição do IP 192.168.1.1 dentro de um segundo deve ser bloqueada.
- **Exemplo de Limitação por Token**: Se um token `abc123` estiver configurado com um limite de 10 requisições por segundo, a 11ª requisição dentro desse segundo deve ser bloqueada.
- **Tempo de Expiração**: Após atingir o limite, novas requisições do mesmo IP ou token só são possíveis após o tempo de expiração.
- **Configuração Personalizada por IP ou Token**: Suponha que o `CUSTOM_MAX_REQ_PER_SEC` esteja configurado com "192.168.1.2=2", isso significa que o IP 192.168.1.2 terá um limite personalizado de 2 requisições por tempo, independente do limite padrão para outros IPs.

## Variáveis de Configuração

As variáveis de ambiente desempenham um papel crucial na configuração e personalização do comportamento do limitador de taxa:

- **REDIS_ADDRESS**: Define o endereço do servidor Redis utilizado pelo limitador de taxa.
- **REDIS_PASSWORD**: Senha para autenticação no servidor Redis.
- **REDIS_DB**: Número do banco de dados Redis a ser utilizado pelo aplicativo.
- **DEFAULT_IP_MAX_REQ_PER_SEC**: Define o limite padrão de requisições por segundo por endereço IP.
- **DEFAULT_TOKEN_MAX_REQ_PER_SEC**: Estabelece o limite padrão de requisições por segundo por token de acesso.
- **DEFAULT_IP_BLOCK_DURATION**: Duração do bloqueio padrão para um endereço IP que excede seu limite de requisições.
- **DEFAULT_TOKEN_BLOCK_DURATION**: Duração do bloqueio padrão para um token que excede seu limite de requisições.
- **CUSTOM_MAX_REQ_PER_SEC**: Permite a definição de limites de requisição personalizados para IPs ou tokens específicos.
- **CUSTOM_BLOCK_DURATION**: Configura durações de bloqueio personalizadas para IPs ou tokens específicos.

Essas variáveis são essenciais para a flexibilidade e eficácia do limitador, permitindo sua adaptação às necessidades específicas do sistema.

## Funcionamento do Limitador de Taxa

O limitador de taxa é implementado como um middleware no servidor HTTP, permitindo que ele intercepte e controle as requisições.

### Middleware de Limitação de Taxa

- O middleware `RateLimiterMiddleware` é aplicado a cada requisição recebida pelo servidor.
- Ele identifica cada requisição por um identificador único, que pode ser um token de acesso (se presente no cabeçalho API_KEY) ou o endereço IP do solicitante.
- Após identificar a requisição, o middleware consulta o `RateLimit`, uma estrutura que contém as configurações de limitação e o Redis para verificar se o limite foi excedido.

### Verificação e Controle de Limite

- A função `IsLimitExceeded` da estrutura `RateLimit` é responsável por determinar se uma requisição excede o limite configurado.
- Ela verifica o número atual de requisições feitas pelo identificador (IP ou
