import random
import string

def randomword(length):
   letters = string.ascii_lowercase
   return ''.join(random.choice(letters) for i in range(length))

def concat(m):
        return m["first"]+ m["last"]
