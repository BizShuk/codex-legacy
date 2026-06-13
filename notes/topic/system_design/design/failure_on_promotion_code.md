# Failure on Promotion Code

Once the promotion code service is down, it means there is not way to know whether the code is valid or how much discount user should get. So, there are two ways to design this failure recovery.

### Solution 1

Tell user the promotion code service is down. The price keeps the same as original, but put description to let user know the situation.  Once the service back, user will get redund from discount or just cut the price before payment request.

### Solution 2

Usually, promotion code falls into few type of promotion. With proper design on the code can show type of promotion and roughly valid date. And use the theoratical promotion code to give user discount no matter the code used or not.

Increase the cost during system failure, but no user can recognize the situation.

### Solution 3 from solution 1

Before payment submitted, logger will trigger a message/date to conjunct with payment to delay the payment request.
