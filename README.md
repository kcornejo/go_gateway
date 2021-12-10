
# Consumo de WS con Golang

Aplicacion realizada para consumos tanto soap como rest desde Go 1.7




## Instalacion

Instalar docker

```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
```
```bash
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```
```bash
sudo apt-get update
 $ sudo apt-get install docker-ce docker-ce-cli containerd.io

```
Instalar Git
```bash
sudo apt install git-all
```
Descargar el Proyecto
```bash
git clone https://github.com/kcornejo/go_gateway.git .
```
Crear imagen con docker
```bash
docker build -t gateway .
```
Crear contenedor
```bash
docker run -dp 8000:8000 gateway
```
Luego ingresar a la url http://localhost:8000/swagger/index.html
## Authors

- [@kcornejo](https://www.github.com/kcornejo)

