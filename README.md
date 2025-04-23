# Stocks API

Proyecto **backend en Go 1. con GIN y JWT** que:

1. Conecta al endpoint del reto y descarga todas las *notas de analistas*  y las *acciones*
2. Enri­quece cada acción con el **precio actual** vía Finnhub  
3. Calcula y guarda un **score** para recomendar la mejor acción  
4. Expone endpoints REST protegidos con **JWT**

Todo vive en Docker Compose: CockroachDB + migrador + API.

---

## 1 . Preparar variables

```bash
cp .env.example .env
```

Ajusta lo que necesites (los valores por defecto funcionan):

---

## 2 . Levantar la aplicación

```bash
docker-compose up --build
```

*Servicios resultantes*

| Servicio | Puerto | Descripción |
| -------- | ------ | ----------- |
| CockroachDB | 26257 | BD en modo `--insecure` |
| API | 8080 | Endpoints REST |

El primer arranque:

1. Corre migraciones.  
2. Inserta **500 stocks** de `seed/data.json` y un usuario `admin / admin`.  
3. API queda en `http://localhost:8080`.

> **Sync completo ≈ 3 000 tickers **  
> ```bash
> go run ./cmd/sync
> ```  
> Limitado a 1 req/s → ≈ 50 min. si se requieren todos los datos y se debe eliminar la data de la tabla `stocks` o el volumen de docker y luego volver a ejectuar el container migrator.

---

## 3 . Endpoints de ejemplo

| Método | Ruta | Descripción |
| ------ | ---- | ----------- |
| POST | `/users` | Crear usuario |
| POST | `/users/login` | Devuelve JWT |
| GET  | `/stocks/getBestStocks?page=0&risk=medium` | Top N por *score* |

