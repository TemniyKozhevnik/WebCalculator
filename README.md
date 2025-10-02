# Web Calculator

Веб-калькулятор с простым и понятным интерфейсом, написанный на Go.

## Возможности

- ➕ Базовые арифметические операции (сложение, вычитание, умножение, деление)
- 🌐 Веб-интерфейс с адаптивным дизайном
- 🔌 REST API для программного доступа
- 🚀 Легковесный и быстрый бэкенд на Go
- 📱 Мобильная версия

## Установка

### Предварительные требования

- Go 1.16 или выше
- Git

### Шаг за шагом

1. **Клонируйте репозиторий**
   ```bash
   git clone https://github.com/TemniyKozhevnik/WebCalculator.git
   cd WebCalculator
   ```

2. **Запустите фронт**
   ```bash
   cd Calculator_Front
   npm install
   npm start
   ```

3. **Настройте переменные окружения**
Создайте файл .env с содержимым:
```env
  PORT=:8080
  HOST=localhost
  DB_PORT=5432
  DB_PASSWORD=test
```

4. **Запустите бэкэнд**
```bash
   cd ..
   cd Calculator_API
   go run cmd/main.go
```
