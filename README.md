# GO-API

API REST básica desarrollada en Go.

## 🚀 Inicio Rápido

### Requisitos
- Go 1.25+

### Instalación

```bash
# Clonar el repositorio
git clone https://github.com/cristianjonhson/GO-API.git
cd GO-API

# Instalar dependencias
go mod download
```

### Ejecutar

```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

## 📚 Endpoints

### GET /
Ruta principal de bienvenida
```bash
curl http://localhost:8080/
```

### GET /api/health
Verificar estado de la API
```bash
curl http://localhost:8080/api/health
```

### GET /api/hello?name=Tu_Nombre
Saludo personalizado
```bash
curl http://localhost:8080/api/hello?name=Cristian
```

## 🛠️ Construir

```bash
go build -o api main.go
./api
```

## 📝 Licencia

MIT
