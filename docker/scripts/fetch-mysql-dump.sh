echo "Fetching data from ergast..."
curl http://ergast.com/downloads/f1db.sql.gz --output f1db.sql.gz
echo "Source data fetched"

echo "unpacking to docker entry point"
gzip -dc f1db.sql.gz > /docker-entrypoint-initdb.d/f1db.sql
