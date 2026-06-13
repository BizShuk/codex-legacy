# Interview Questions
"""
	We have a digested server log with username, visited page and timestamp. Create a processing algorithm that will output the most visited page/areas in such a way that will match partial path as well.
	i.e.
	{
	    { user: "user1", page="/home" },
	    { user: "user1", page="/home/account" },
	    { user: "user1", page="/home/account/profile" },
	    { user: "user1", page="/home/account/login" },
	    { user: "user2", page="/about" },
	    { user: "user2", page="/about/contact" },
	    { user: "user2", page="/home" }
	}

	the output
	user1
	     - home/account
	     - home
	user2
	     - /about
"""


visited =   [
        { 'user': "user1", 'page':"/home" },
        { 'user': "user1", 'page':"/home/account" },
        { 'user': "user1", 'page':"/home/account/profile" },
        { 'user': "user1", 'page':"/home/account/login" },
        { 'user': "user2", 'page':"/about" },
        { 'user': "user2", 'page':"/about/contact" },
        { 'user': "user2", 'page':"/home" }
]

from collections import Counter
def most_used_path(records,uname):
	
    ret = {}
    for e in records:
       
        user = e['user']
        path = e['page']        
        if user not in ret:
            ret[user] = Counter()
        ret[user] = ret[user] + Counter(allpath(path))

        
        
    return ret[uname].most_common(4)
    




def allpath(path):
    """
        return paths
    """
    paths = path.strip("/").split("/")
    for i in range(1,len(paths)):
        paths[i] = paths[i-1] + "/" + paths[i]
    return paths

a = allpath("/home/account/profile")
#print(a)
b =most_used_path(visited,'user1')
print(b)

b =most_used_path(visited,'user2')
print(b)






