conn = new Mongo();
db = conn.getDB("Bus"); 

db.createCollection("tickets");
db.createCollection("all_buses");
db.createCollection("trips");
