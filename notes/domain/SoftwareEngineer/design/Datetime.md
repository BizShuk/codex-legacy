# Datetime

## UTC

### format

`yyyy-MM-ddTHH:mm:ss<timezone>`
or
`yyyyMMddHHmmss<timezone>`

### timezone

- Z, UTC +0
- +02:00
- +08:00
- -06:00

### option meaning

1. d -> Represents the day of the month as a number from 1 through 31.
2. dd -> Represents the day of the month as a number from 01 through 31.
3. ddd-> Represents the abbreviated name of the day (Mon, Tues, Wed, etc).
4. dddd-> Represents the full name of the day (Monday, Tuesday, etc).
5. h-> 12-hour clock hour (e.g. 4).
6. hh-> 12-hour clock, with a leading 0 (e.g. 06)
7. H-> 24-hour clock hour (e.g. 15)
8. HH-> 24-hour clock hour, with a leading 0 (e.g. 22)
9. m-> Minutes
10. mm-> Minutes with a leading zero
11. M-> Month number(eg.3)
12. MM-> Month number with leading zero(eg.04)
13. MMM-> Abbreviated Month Name (e.g. Dec)
14. MMMM-> Full month name (e.g. December)
15. s-> Seconds
16. ss-> Seconds with leading zero
17. t-> Abbreviated AM / PM (e.g. A or P)
18. tt-> AM / PM (e.g. AM or PM
19. y-> Year, no leading zero (e.g. 2015 would be 15)
20. yy-> Year, leading zero (e.g. 2015 would be 015)
21. yyy-> Year, (e.g. 2015)
22. yyyy-> Year, (e.g. 2015)
23. K-> Represents the time zone information of a date and time value (e.g. +05:00)
24. z-> With DateTime values represents the signed offset of the local operating system's time zone from
25. Coordinated Universal Time (UTC), measured in hours. (e.g. +6)
26. zz-> As z, but with leading zero (e.g. +06)
27. zzz-> With DateTime values represents the signed offset of the local operating system's time zone from UTC, measured in hours and minutes. (e.g. +06:00)
28. f-> Represents the most significant digit of the seconds' fraction; that is, it represents the tenths of a second in a date and time value.
29. ff-> Represents the two most significant digits of the seconds' fraction in date and time
30. fff-> Represents the three most significant digits of the seconds' fraction; that is, it represents the milliseconds in a date and time value.
31. ffff-> Represents the four most significant digits of the seconds' fraction; that is, it represents the ten-thousandths of a second in a date and time value. While it is possible to display the ten-thousandths of a second component of a time value, that value may not be meaningful.
32. fffff-> Represents the five most significant digits of the seconds' fraction; that is, it represents the hundred-thousandths of a second in a date and time value.
33. ffffff-> Represents the six most significant digits of the seconds' fraction; that is, it represents the millionths of a second in a date and time value.
34. fffffff-> Represents the seven most significant digits of the second's fraction; that is, it represents the ten-millionths of a second in a date and time value.
