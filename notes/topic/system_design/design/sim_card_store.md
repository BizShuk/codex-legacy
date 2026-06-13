# Sim Card store system

Requirement:

- Design a sim card store system that gives you 3 sim cards on demand and each of them has a 10-digit long phone number.
- The generated numbers should be very hard to predict.
- We need to keep track of the phone numbers that are sold out and also the total unused numbers the system has at any point in time.
- Phone numbers should be stored very efficiently and retrieval of available phone numbers should be fast.

### Solution 1

10-digits long phone number => total count of phone number: 1b

1b is not a number too big to generate.

We can generate all numbers and then shuffle it multiple times. With efficient storing those numbers, it can just keep all used numbers with some unused numbers in DB. Refill later if the system is running out of unused number.

ps: All unused numbers should be stored in another place, like file system. Only load some amount of data to DB to keep DB lean and efficient.

### Solution 2

By creating a algo to map one to another. A good math equation/algo can create 1-1 mapping without collision. The counter from 0 to 1b, but the expression will map to some number which is hard to predict. DB will store those used number sequentially. We also know the latest original number and the rest of numbers are unused. It can just give next few numbers to be available phone numbers.

ps: this algo is hard to create with randomness and 1-1 mapping.

### Solution 3

Similiar to Solution 1, we develop trunks of phone numbers. We randomly select one of trunks to serve phone numbers until running out of trunks. It doesn't pre-allocate all unused phone numbers.

ps: It's easy to predict phone number in the trunk range, but not overall.
