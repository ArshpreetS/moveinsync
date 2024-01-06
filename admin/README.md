1. Created a docker network




2. Start the container for mongo express
docker run \
-p 8081:8081 \
-e ME_CONFIG_MONGODB_ADMINUSERNAME=root \    
-e ME_CONFIG_MONGODB_ADMINPASSWORD=root \    
-e ME_CONFIG_MONGODB_SERVER=mongodb \
--name mongo-express \
--net moveinsync \           
mongo-express

3. Start the container for mongodb
docker run -d \
-p 27017:27017 \
-e MONGO_INITDB_ROOT_USERNAME=root \    
-e MONGO_INITDB_ROOT_PASSWORD=root \    
--name mongodb \
-v ./data/mongo:/data/db \ 
--net moveinsync \           
mongo
