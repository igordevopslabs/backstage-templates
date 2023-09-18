# Application: ${{ values.name }}
{% if values.typeApp == 'back-end' -%}
# Type: ${{ values.typeApp }}
{% if values.languageApp == 'iis' -%}
FROM mcr.microsoft.com/dotnet/aspnet:${{ values.versionApp }}
{%- elseif values.languageApp == 'csharp' -%}
FROM mcr.microsoft.com/dotnet/runtime:${{ values.versionApp }}
{%- elseif values.languageApp is defined -%}
FROM ${{ values.languageApp }}:${{ values.versionApp }}
{%- endif %}
LABEL TEAM=${{ values.team }}

# Criar usuário não root com o id 1234
RUN useradd -u 1234 appuser
USER appuser

# Especifica o diretório da aplicação
WORKDIR /app
RUN ls -ld /app

# Copia os arquivos para dentro do container com o usuário não root
COPY --chown=appuser:appuser src/* /app

# Comandos para o build ou execução do binário
# RUN npm install

# Executa comandos dentro do container, antes de fechar a imagem
CMD ["echo","Image created"] 
{%- elseif values.typeApp == 'front-end' -%}
# Type: ${{ values.typeApp }}
{% if values.languageApp == 'css' or values.languageApp == 'javascript' -%}
FROM nginx:${{ values.versionApp }}
LABEL TEAM=${{ values.team }}

# Criar usuário não root com o id 1234
RUN useradd -u 1234 appuser
USER appuser

# Copia os arquivos para dentro do container com o usuário root
COPY src/default.conf /etc/nginx/conf.d/default.conf
COPY src/app /usr/share/nginx/html

# Executa comandos dentro do container, antes de fechar a imagem
CMD ["echo","Image created"] 
{%- elseif values.languageApp == 'angular' or values.languageApp == 'react' or values.languageApp == 'typescript' -%}
FROM node:${{ values.versionApp }}
LABEL TEAM=${{ values.team }}

# Criar usuário não root com o id 1234
RUN useradd -u 1234 appuser
USER appuser

# Especifica o diretório da aplicação
WORKDIR /app
RUN ls -ld /app

# Copia os arquivos para dentro do container com o usuário não root
COPY --chown=appuser:appuser src/* /app

# Comandos para o build ou execução do binário
# RUN npm install

# Executa comandos dentro do container, antes de fechar a imagem
CMD ["echo","Image created"] 
{%- endif %}
{%- endif %}