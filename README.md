# GoCacheHub

Bu proje, basit bir anahtar-değer deposu sunucu uygulamasına bir arayüz sağlar. Kullanıcılar, belirli bir anahtar için değerleri ayarlayabilir (POST) ve daha sonra bu değerleri alabilir (GET).

## Nasıl Çalışır?

### Veri Ayarlama (Set)

`POST` isteği ile `/set` endpoint'ine JSON gövdesinde bir anahtar-değer çifti gönderilir. Bu, sunucuda belirtilen anahtar için değeri ayarlar.

Örnek `POST` isteği:

```
POST /set HTTP/1.1
Host: [Sunucu IP Adresi]
Content-Type: application/json

{
  "key": "anahtar_ismi",
  "value": "değer"
}
```

Başarılı bir `POST` isteğinin yanıtı:

```
{
  "message": "Key 'anahtar_ismi' set to 'değer'."
}
```

### Veri Alma (Get)

`GET` isteği ile `/get` endpoint'ine bir sorgu parametresi olarak anahtar gönderilir. Bu, sunucudan belirtilen anahtarın değerini alır.

Örnek `GET` isteği:

```
GET /get?key=anahtar_ismi HTTP/1.1
Host: [Sunucu IP Adresi]
```

Başarılı bir `GET` isteğinin yanıtı:

```
{
  "value": "anahtar_ismi için değer"
}
```

## Kurulum

Bu projeyi lokal sunucunuzda çalıştırmak için aşağıdaki adımları izleyin:

1. Projeyi klonlayın.
2. Gerekli bağımlılıkları yükleyin.
3. Sunucuyu başlatın.

## Kullanım

Projeyi kendi makinenizde veya bir sunucuda çalıştırdıktan sonra, yukarıda açıklanan `POST` ve `GET` isteklerini kullanarak veri ayarlayabilir ve alabilirsiniz.
