db.createUser(
        {
            user: "root",
            pwd: "root",
            roles: [
                {
                    role: "readWrite",
                    db: "Bus"
                }
            ]
        }
);


db = new Mongo("mongodb://root:root@localhost:27017").getDB("Bus");


db.createCollection("trips", {capped: false});
db.createCollection("all_buses", {capped: false});
db.createCollection("tickets", {capped: false});
