## Nginx Configs

Files in this directory goes into the nginx configs at `/etc/nginx/`
1. `{domain}.conf` goes into `/sites-available/`
1. Then symlink them with `sudo ln -s /etc/nginx/sites-available/{domain}.conf
/etc/nginx/sites-enabled/`