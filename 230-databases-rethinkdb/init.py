import rethinkdb as r
import random

conn = r.connect(host="localhost", port=7704, db="players", auth_key="", timeout=20)

# all = r.db('players').table('scores').run(conn)
# print(dir(all))

# for doc in all:
#     print(doc)

while True:
    scores = r.table("scores")
    id = str(random.randint(1,1000))
    # data = scores.get(id).run(conn)
    r.table("scores").get(id).update({"Score": random.randint(1,10)}).run(conn)
