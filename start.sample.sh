
#Facebook
export FB_SECRET="{FB_SECRET}" #App Client Secret
export FB_WEBHOOK_TOKEN="{FB_WEBHOOK_TOKEN}" #Value to be set by dev from FB dashboard

#Instagram
export IG_SECRET="{IG_SECRET}" #App Client Secret

#Go-Webapp
export AUTH_TOKEN_SECRET="{AUTH_TOKEN_SECRET}" #JWT Token Secret
export MONGODB_URI="" #(without db name), defaults to `mongodb://localhost/`
export MONGODB_NAME="" #Databse Name, defaults to `go-webapp`
export PORT="" #defaults to 3000
export REDIS_URL="" #(without http/https), defaults to `localhost:6379`
export REDIS_PWD="" #Redis password


#Start Server using Gin, auto update on file changes
gin -p 8911 -i
