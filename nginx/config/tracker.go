package config

// TrackerTemplate is the default tracker nginx tmpl.
const TrackerTemplate = `
proxy_cache_path {{.cache_dir}}/metainfo levels=1:2 keys_zone=metainfo:10m max_size=256g;

upstream tracker {
  server {{.server}};
}

server {
  listen {{.port}};

  {{.client_verification}}

  access_log {{.log_dir}}/nginx-access.log;
  error_log {{.log_dir}}/nginx-error.log;

  location / {
    proxy_pass http://tracker;
  }

  location ~* ^/namespace/.*/blobs/.*/metainfo$ {
    proxy_pass http://tracker;

    proxy_cache         metainfo;
    proxy_cache_methods GET;
    proxy_cache_valid   200 5m;
    proxy_cache_valid   any 1s;
    proxy_cache_lock    on;
  }
}
`