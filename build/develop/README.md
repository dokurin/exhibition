# 開発環境Dockerイメージ

開発環境として利用するDockerコンテナイメージ

開発ツールとかはローカルにインストールするんじゃなくて、このコンテナイメージに追加してこうぜ

利用するにはDockerが必須なので **ホストOSの側にDockerをインストール** しておくこと

## Required

- [Docker Desctop](https://www.docker.com/products/docker-desktop)

## Quick Start

### イメージをビルドする

```console
./build.sh
```

イメージのビルドは初回のみでOK

### コンテナを実行する

```console
./exec.sh
```

## インストールされているツール/言語

| ツール/言語                                                        | バージョン(指定している場合は記入しよう) |
|--------------------------------------------------------------------|------------------------------------------|
| [curl](https://curl.haxx.se/)                                      | -                                        |
| [gettext](https://ayatakesi.github.io/gettext/0.18.3/gettext.html) | -                                        |
| [git](https://git-scm.com/)                                        | -                                        |
| [hub](https://hub.github.com/)                                     | -                                        |
| [jq](https://stedolan.github.io/jq/)                               | -                                        |
| [tree](http://mama.indstate.edu/users/ice/tree/)                   | -                                        |
| [vim](https://www.vim.org/)                                        | -                                        |
| [wget](https://www.gnu.org/software/wget/)                         | -                                        |
| [Docker Client](https://www.docker.com/)                           | 19.03.12, build 48a66213fe               |
| [Docker Comose](https://docs.docker.com/compose/)                  | 1.26.2, build eefe0d31                   |
| [Go](https://golang.org/)                                          | 1.15.2                                   |
| [Node.js](https://nodejs.org/ja/)                                  | 14.11.0                                  |
