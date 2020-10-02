# APIサーバプロトタイプ

## Start

```console
go run .
```

## Request Sample

### 作家一覧

```console
curl http://localhost:8080/api/artists
```

### 作家詳細

```console
curl http://localhost:8080/api/artists/artist1
```

### 作品一覧

```console
curl http://localhost:8080/api/artists/artist1/artofworks
```
