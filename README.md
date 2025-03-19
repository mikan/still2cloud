still2cloud
-----------

[![Go Report Card](https://goreportcard.com/badge/github.com/mikan/still2cloud)](https://goreportcard.com/report/github.com/mikan/still2cloud)

Take a still picture from IP camera and upload to Amazon S3 or S3-compatible object storages.

## Download

See [Releases](https://github.com/mikan/still2cloud/releases) page.

## Usage

Run executable binary:

```sh
./still2cloud -c config.json
```

Run with Golang toolchain:

```sh
go run . -c config.json
```

Build executable binary:

```sh
make build
```

## Sources

### HTTP

#### No authentication

Example configuration:

```json
{
  "src": {
    "type": "http",
    "url": "http://192.168.1.100:81/snapshot.cgi?user=username&pwd=password"
  },
  "dest": {
  }
}
```

#### Basic authentication

Example configuration:

```json
{
  "src": {
    "type": "http",
    "url": "http://192.168.1.100:81/snapshot.cgi",
    "auth": "basic",
    "user": "username",
    "password": "password"
  },
  "dest": {
  }
}
```

#### Digest authentication

Example configuration for HikVision IP cameras:

```json
{
  "src": {
    "type": "http",
    "url": "http://192.168.1.100/ISAPI/Streaming/channels/101/picture",
    "auth": "digest",
    "user": "username",
    "password": "password"
  },
  "dest": {
  }
}
```

### File

Example configuration:

```json
{
  "src": {
    "type": "file",
    "path": "snapshot.jpg"
  },
  "dest": {
  }
}
```

### RTSP

Example configuration:

```json
{
  "src": {
    "type": "rtsp",
    "url": "rtsp://user:password@192.168.1.100:554/channel1",
    "path": "tmp.jpg"
  },
  "dest": {
  }
}
```

NOTE: RTSP option requires `ffmpeg` command.

## Destinations

### Amazon S3

Example configuration for Cloudflare R2:

```json
{
  "src": {
  },
  "dest": {
    "type": "s3",
    "endpoint": "https://xxx.r2.cloudflarestorage.com",
    "bucket": "bucket",
    "path_layout": "still/location-a/camera-1/20060102-150405.jpg",
    "layout_mode": 2,
    "region": "auto",
    "access_key_id": "xxx",
    "secret_access_key": "xxx",
    "create_latest_file": true,
    "latest_file_path": "still/location-a/camera-1/latest.txt"
  }
}
```

`path_layout` allows Golang's time layout (`20060102-150405` means `yyyyMMdd-HHmmss`).
You can disable with `layout_mode` to `1` (`0`: apply to full path, `2`: apply to file name).

### File

Example configuration:

```json
{
  "src": {
  },
  "dest": {
    "type": "file",
    "path_layout": "snapshot-20060102-150405.jpg"
  }
}
```

`path_layout` allows Golang's time layout (`20060102-150405` means `yyyyMMdd-HHmmss`).
You can disable with `layout_mode` to `1` (`0`: apply to full path, `2`: apply to file name).

## Converting

The converting (formatting and resizing) feature is currently under development.

## License

[BSD 3-clause](LICENSE)

## Author

[mikan](https://github.com/mikan)
