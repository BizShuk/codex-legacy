# Exercise: Implement a class for inspecting a log file.
Complete the implementation of the `AccessLog` class in [solution.py](solution.py).  The class is initialized with a path to a web server access log file.  You can assume the maximum log file size is 50MB.

# Log File Format
Each row in the the log file `access.log` contains the following space-separated fields:

`remote_ip` `-` `-` `[time_local]` `"request"` `status_code` `bytes` `"referer"` `"user_agent"`

Fields that contains spaces (`time_local`, `referer` and `user_agent`) are enquoted. `time_local` is enclosed in square brackets, while `referer` and `user_agent` are enclosed within double quotes.

The project will contain a function for parsing this log format.

# Please DO

  * Write good quality code and tests.
  * Implement your code in Python 2.6 or 2.7
  * Use any appropriate functionality from the Python Standard Library.

# Please do NOT

  * Change the existing method signatures.
  * Use any third-party libraries

# Example Usage
    :::python
        >>> access_log = AccessLog('./access.log')
        >>> access_log.contains_ip('203.166.249.147')
        True
        >>> access_log.contains_ip('72.14.199.38')
        True
        >>> access_log.contains_ip('1.1.1.1')
        False
        >>> access_log.contains_ip('2.2.2.2')
        False
        >>> access_log.get_status_count(302)
        5
        >>> access_log.get_status_count(403)
        0

# Please answer the following questions

## What are the pros and cons of your solution?
  * Pros
    - easy to code and collect what i want
  * Cons
    - It's not good for bigdata.

## If the log file was very large (e.g. 10TB) what effects could it have on:
  * Performance of `contains_ip`
  O(1) 
  * Performance of `get_status_count`
  O(1)
  * Memory consumption
  as large as file

## What changes would you make to your code so it could deal with very large files (e.g. 10TB). How would these changes affect performance and resource usage?
break it to many blocks , use thread things or bigdata ( spark or hadoop) 


# Please add any comments you feel are relevant.
