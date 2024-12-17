/*
Package firstRestfulApi предоставляет базовую структуру HTTP-сервера с функциями запуска и завершения работы.

Этот пакет предназачен для создания простого и настраиваемого HTTP-сервера с возможностью обработки запросов, установки тайм-аутов и ограничения размера заголовков. Он может быть расширен для обработки маршрутов или настройки API.

Пример:

 package main

 import (
  "context"
  "log"
  "os"
  "os/signal"
  "syscall"
  "time"
  "firstRestfulApi"
 )

 func main() {
  // Создание экземпляра сервера
  server := &firstRestfulApi.Server{}

  // Канал для обработки сигналов завершения
  quit := make(chan os.Signal, 1)
  signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

  // Запуск сервера в горутине
  go func() {
   if err := server.Run(":8080"); err != nil && err != http.ErrServerClosed {
    log.Fatalf("Failed to start server: %s", err)
   }
  }()
  log.Println("Server is running on port 8080...")

  // Ожидаем сигнала завершения
  <-quit
  log.Println("Shutting down server...")

  // Завершаем работу сервера с тайм-аутом
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  if err := server.Shutdown(ctx); err != nil {
   log.Fatalf("Server shutdown failed: %s", err)
  }

  log.Println("Server gracefully stopped")
 }
*/

package todo_app

import (
	"context"
	"net/http"
	"time"
)

// Server представляет HTTP-сервер с возможностью запуска и завершения работы.
//
// Он предоставляет базовый функционал для конфигурирования и управления сервером.
// Тайм-ауты чтения и записи, а также максимальный размер заголовков определены
// как настройки по умолчанию, которые можно изменять при необходимости.
type Server struct {
	httpServer *http.Server // Экземпляр стандартного HTTP-сервера
}

// Run запускает HTTP-сервер на указанном порту.
//
// Параметры:
//   - port (string): Адрес (в формате "localhost:8080") или порт (":8080"), на котором будет запускаться сервер.
//
// Сервер конфигурируется с минимальными настройками:
//   - MaxHeaderBytes: 1 MB (максимальный размер заголовков запросов).
//   - ReadTimeout: 10 секунд (тайм-аут чтения запросов).
//   - WriteTimeout: 10 секунд (тайм-аут записи ответов).
//
// Возвращает:
//   - error: Ошибка, возникшая при запуске сервера или во время его работы.
//
// Пример:
//
// server := &firstRestfulApi.Server{}
//
//	if err := server.Run(":8080"); err != nil {
//	 log.Fatalf("Server failed to start: %v", err)
//	}
func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown корректно завершает работу HTTP-сервера.
//
// Этот метод завершает существующие HTTP-соединения и закрывает сервер. Для этого используется `context.Context`, который помогает задавать ограничения по времени или отменять процесс завершения.
//
// Параметры:
//   - ctx (context.Context): Контекст для управления временем жизни операции завершения.
//
// Возвращает:
//   - error: Ошибка, если сервер не смог корректно завершить работу.
//
// Пример:
//
// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// defer cancel()
//
//	if err := server.Shutdown(ctx); err != nil {
//	 log.Fatalf("Server shutdown failed: %v", err)
//	}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
