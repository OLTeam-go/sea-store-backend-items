curl -X PATCH https://api.heroku.com/apps/sea-store-backend-items/formation \
  -d '{
  "updates": [
    {
      "type": "web",
      "docker_image": "'$1'"
    }
  ]
}' \
  -H "Content-Type: application/json" \
  -H "Accept: application/vnd.heroku+json; version=3.docker-releases" \
  -H "Authorization: Bearer $2"
