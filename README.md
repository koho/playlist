# PlayList

Play your remote media folder via HTTP in your favorite media player.

## Get started

### Install ffmpeg

[ffmpeg](https://ffmpeg.org/) must be installed first.

### Create your config file

Place the config file `config.yml` to the same directory of the binary.

```yaml
listen: :6300
groups:
  - name: private
    path: /data/private
    username: john
    password: 123456

  - name: share
    path: /usr/share/family
```

### Start web server

```shell
./playlist
```

### Play

Open your favorite media player, import the playlist with the following url.

- http://server_ip:6300/private
- http://server_ip:6300/share

## Parameters

| Parameter       | Description                                                                                                        |
|-----------------|--------------------------------------------------------------------------------------------------------------------|
| listen          | Listen address of the web server.                                                                                  |
| groups          | All your media groups.                                                                                             |
| groups.name     | Name of the media group.                                                                                           |
| groups.path     | Path is where your media folder located.                                                                           |
| groups.url      | URL overwrites the default url of the media file. The final url joins the given url with the media file name.      |
| groups.username | When Username is non-empty, the HTTP Basic Auth will be enabled. Otherwise, no authentication is needed.           |
| groups.password | Password of the user.                                                                                              |
| thumb.dir       | Dir is where the thumbnail of media file stored.                                                                   |
| thumb.workers   | Workers is the number of worker processes to generate thumbnails in parallel. Default is the half of logical CPUs. |
| thumb.size      | Size sets the output thumbnail size. Default is 640:360.                                                           |


## With nginx

```
server {
  listen 8933;
  listen [::]:8933;
  location /playlist/ {
    proxy_pass        http://127.0.0.1:6300/;
    proxy_set_header  X-Real-IP         $remote_addr;
    proxy_set_header  X-Forwarded-For   $proxy_add_x_forwarded_for;
    proxy_set_header  X-Forwarded-Proto $scheme;
    proxy_set_header  Host              $http_host;
    proxy_set_header  X-Original-URI    $request_uri;
  }
}
```
