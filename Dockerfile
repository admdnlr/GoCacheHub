# Go uygulaması için Dockerfile örneği

# Derleme aşaması için Go derleyici içeren resmi bir base image kullanın
FROM golang:1.21.6-alpine AS build

# Çalışma dizinini ayarlayın
WORKDIR /app

# Go modüllerini yönetmek için gerekli dosyaları kopyalayın
COPY go.mod ./
COPY go.sum ./

# Bağımlılıkları indirin
RUN go mod download

# Kaynak kodunu kopyalayın
COPY . .

WORKDIR /app/cmd/GoCacheHub

# Uygulamanızı derleyin
RUN go build -o /GoCacheHub

# Çalıştırma aşaması için hafif bir base image kullanın
FROM alpine:latest

# Çalışma dizinini ayarlayın
WORKDIR /

# Derleme aşamasından oluşturulan çalıştırılabilir dosyayı kopyalayın
COPY --from=build /GoCacheHub /GoCacheHub

# Portu açın (uygulamanızın kullandığı porta göre değiştirin)
EXPOSE 9090

# Uygulamanızı çalıştırın
CMD ["/GoCacheHub"]
