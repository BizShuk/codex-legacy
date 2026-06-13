# Interview Questions

"""
	You have rating (0-10) of the hotels per user in this format:

	scores = [
	    {'hotel_id': 1001, 'user_id': 501, 'score': 7},
	    {'hotel_id': 1001, 'user_id': 502, 'score': 7},
	    {'hotel_id': 1001, 'user_id': 503, 'score': 7},
	    {'hotel_id': 2001, 'user_id': 504, 'score': 10},
	    {'hotel_id': 3001, 'user_id': 505, 'score': 5},
	    {'hotel_id': 2001, 'user_id': 506, 'score': 5}
	]

	Any given hotel might have more than one score.

	Implement a function, get_hotels(scores, min_avg_score) that returns a list of hotel ids that have average score equal to or higher than min_avg_score.

	get_hotels(scores, 5) -> [1001, 2001, 3001]
	get_hotels(scores, 7) -> [1001, 2001]
"""

datas = [
	{'hotel_id': 1001, 'user_id': 501, 'score': 7},
	{'hotel_id': 1001, 'user_id': 502, 'score': 7},
	{'hotel_id': 1001, 'user_id': 503, 'score': 7},
	{'hotel_id': 2001, 'user_id': 504, 'score': 10},
	{'hotel_id': 3001, 'user_id': 505, 'score': 5},
	{'hotel_id': 2001, 'user_id': 506, 'score': 5}
]


scores = datas

def get_hotels(scores,score):
    v = {}
    r = {}
    for e in scores:
        hid,uid,_score = e['hotel_id'],e['user_id'],e['score']
        v[hid] = v[hid] + _score if hid in v else _score
        r[hid] = r[hid] + 1 if hid in r else 1
    
    return [hid for hid in r if v[hid]/r[hid] >=score]


ret = get_hotels(scores,5)
print(ret)
ret = get_hotels(scores,7)
print(ret)

