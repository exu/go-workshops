import rethinkdb as r
import random

conn = r.connect(host="localhost", port=7704, db="players", auth_key="", timeout=20)

while True:
    scores = r.table("scores")
    id = str(random.randint(1,1000))
    data = scores.get(id).run(conn)
    newScore = random.randint(1,10) + data["Score"]
    r.table("scores").get(id).update({"Score": newScore}).run(conn)
