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
git clone https://github.com/tu-usuario/neo-wifi.git
cd neo-wifi
go run main.go
