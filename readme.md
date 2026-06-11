# logxgo

Logger estructurado para Go, ligero, simple y seguro para concurrencia, construido sobre Logrus.

`logxgo` proporciona una API limpia para generar logs en formato JSON, agregar campos estructurados y crear loggers contextuales reutilizables.

---

# Características

* Logs estructurados en formato JSON
* Campos dinámicos
* Loggers contextuales
* Configuración opcional
* Seguro para múltiples goroutines
* API minimalista
* Basado en Logrus

---

# Instalación

```bash
go get github.com/aleonardomc/logxgo
```

---

# Uso básico

```go
package main

import "github.com/aleonardomc/logxgo"

func main() {

	log := logxgo.New()

	log.Info("Aplicación iniciada")

	log.Error("Error procesando solicitud")
}
```

Salida:

```json
{
  "level":"info",
  "msg":"Aplicación iniciada",
  "time":"2026-06-10T12:00:00-06:00"
}
```

```json
{
  "level":"error",
  "msg":"Error procesando solicitud",
  "time":"2026-06-10T12:00:01-06:00"
}
```

---

# Configuración opcional

## Nivel de log

```go
log := logxgo.New(
	logxgo.WithLevel("debug"),
)
```

Niveles soportados:

* trace
* debug
* info
* warn
* error
* panic

Por defecto se utiliza `info`.

---

## Módulo

```go
log := logxgo.New(
	logxgo.WithModule("api"),
)
```

Salida:

```json
{
  "level":"info",
  "module":"api",
  "msg":"Aplicación iniciada",
  "time":"2026-06-10T12:00:00-06:00"
}
```

---

## Configuración completa

```go
log := logxgo.New(
	logxgo.WithLevel("debug"),
	logxgo.WithModule("api"),
)
```

---

# Campos estructurados

Puedes agregar información adicional a cualquier log.

```go
log.Error(
	"Error procesando solicitud",
	logxgo.F("status_code", 500),
	logxgo.F("path", "/api/users"),
)
```

Salida:

```json
{
  "level":"error",
  "msg":"Error procesando solicitud",
  "status_code":500,
  "path":"/api/users",
  "time":"2026-06-10T12:00:00-06:00"
}
```

---

# Logger contextual

Un logger contextual permite agregar campos una sola vez y reutilizarlos automáticamente en múltiples logs.

```go
requestLogger := log.WithFields(
	logxgo.F("request_id", "abc-123"),
	logxgo.F("user_id", 1001),
)
```

Ahora todos los logs incluirán esos campos:

```go
requestLogger.Info("Solicitud recibida")
requestLogger.Info("Procesando solicitud")
requestLogger.Error("Error validando datos")
```

Salida:

```json
{
  "level":"info",
  "request_id":"abc-123",
  "user_id":1001,
  "msg":"Solicitud recibida"
}
```

```json
{
  "level":"info",
  "request_id":"abc-123",
  "user_id":1001,
  "msg":"Procesando solicitud"
}
```

```json
{
  "level":"error",
  "request_id":"abc-123",
  "user_id":1001,
  "msg":"Error validando datos"
}
```

---

# Cambio dinámico de nivel

Puedes cambiar el nivel de log en tiempo de ejecución.

```go
log.SetLevel("debug")
```

---

# API

## Crear logger

```go
logxgo.New(options ...Option)
```

o

```go
logxgo.NewLogger(options ...Option)
```

---

## Opciones

```go
logxgo.WithLevel("debug")
logxgo.WithModule("api")
```

---

## Campos

```go
logxgo.F("key", value)
```

---

## Métodos disponibles

```go
Info(msg string, fields ...Field)
Error(msg string, fields ...Field)
Warn(msg string, fields ...Field)
Debug(msg string, fields ...Field)
Trace(msg string, fields ...Field)
Panic(msg string, fields ...Field)
```

---

## Logger contextual

```go
WithFields(fields ...Field) ILogger
```

---

## Cambio de nivel

```go
SetLevel(level string)
```

---

# Filosofía

`logxgo` busca ofrecer una experiencia simple para logging estructurado en Go.

Principios:

* Configuración mínima
* Logs JSON por defecto
* Seguro para concurrencia
* Sin singletons obligatorios
* Campos estructurados
* Contexto reutilizable mediante `WithFields`
* API pequeña y fácil de aprender

---

# Concurrencia

`logxgo` utiliza Logrus internamente, por lo que puede utilizarse de forma segura desde múltiples goroutines.

No es necesario implementar mutexes ni mecanismos adicionales para registrar logs concurrentemente.

---

# ¿Por qué logxgo?

Muchos wrappers de logging agregan complejidad innecesaria mediante:

* Singletons obligatorios
* Métodos repetitivos
* Configuraciones complejas
* Acoplamiento con servicios externos

`logxgo` busca mantener una API limpia y flexible:

```go
log.Info(
	"Usuario autenticado",
	logxgo.F("user_id", 1001),
	logxgo.F("role", "admin"),
)
```

sin necesidad de crear métodos específicos para cada escenario.

---