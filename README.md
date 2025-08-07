# 🌐 Neo WiFi API

API geográfica desarrollada en **Go** que permite obtener los puntos WiFi más cercanos (2.4GHz y 5GHz), ciudades y aeropuertos, a partir de coordenadas GPS.

El sistema utiliza la fórmula de Haversine para calcular distancias entre puntos geográficos a gran escala y devuelve los tres puntos WiFi más cercanos disponibles.  
Además, permite buscar por nombre de red (`query`) para obtener datos específicos de una antena.

---

## 📦 Tecnologías

- [Go (Golang)](https://golang.org/)
- Haversine formula
- Datos cargados desde archivos `.json`
- Despliegue en [Render.com](https://render.com)

---

## 🚀 Cómo levantar el proyecto en local

### Requisitos

- Go instalado (`go version`)
- Git

### Pasos

```bash
git clone https://github.com/solidsnk86/neo-wifi-api.git
cd neo-wifi-api
go run main.go
```

## 📡 Endpoints

GET /api/closest
Devuelve la ciudad, antenas y aeropuerto más cercanos a unas coordenadas dadas.

Parámetros requeridos:

lat → Latitud (ej: -34.60)

lon → Longitud (ej: -58.38)

query → (opcional) nombre exacto de red WiFi para buscar una antena una ves que se hayan proporcionado los parámetros de latitud y longitud.

Ejemplo:

```bash
GET /api/closest?lat=-34.60&lon=-58.38
```

### Ejemplo con búsqueda específica:

```bash
GET /api/closest?lat=-34.60&lon=-58.38&query=WiFi3.0-CO-28
````
